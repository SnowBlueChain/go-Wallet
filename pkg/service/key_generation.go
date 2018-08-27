package service

import (
	"github.com/hiromaily/go-bitcoin/pkg/key"
	"github.com/hiromaily/go-bitcoin/pkg/logger"
	"github.com/pkg/errors"
)

//Seed生成は完全に分離したほうがいい
//1.Seedの生成+DBに登録
//2.Multisig Keyの生成+DBに登録(承認用は端末を分けて管理しないと意味がないかも)

//CreateMultiSig(addmultisigaddress)にwalletにmultisig用のprivate keyを登録する
//これのパラメータには、multisigしないと送金許可しないアドレス(receipt, payment)+承認用のアドレスをセット
//これによって、生成されたアドレスから、送金する場合、パラメータにセットしたアドレスに紐づく秘密鍵が必要
//payment,receiptのアドレスは、実際には、addmultisigaddressによって生成されたアドレスに置き換えられる。

//含まれるもの
//coldwallet1: client, receipt, payment
//coldwallet2: multisig address
//TODO: =>どこのマシンで、addmultisigaddressを行う？？

//https://www.slideshare.net/ssusere174e3/ss-33733512

//3.Client Keyの生成+DBに登録
//4.Receipt Keyの生成 + Multisig対応 + DBに登録 (1日1Key消費するイメージ)
//5.Payment Keyの生成+ Multisig + DBに登録 (1日1Key消費するイメージ)

// InitialKeyGeneration 初回生成用のfunc
func (w *Wallet) InitialKeyGeneration() error {

	// Seed
	seed, err := w.retrieveSeed()
	//seed, err := w.GenerateSeed()
	if err != nil {
		return errors.Errorf("retrieveSeed() error: %v", err)
	}

	//accounts := []key.AccountType{
	//	key.AccountTypeClient,
	//	key.AccountTypeReceipt,
	//	key.AccountTypePayment,
	//	key.AccountTypeMultisig,
	//}

	// TODO:初回の生成数をどのように調整し、決定するか？とりあえず固定
	// アカウント(Client)を生成
	_, err = w.GenerateClientAccount(seed, 0, 10000)
	if err != nil {
		return errors.Errorf("GenerateClientAccount() error: %v", err)
	}

	// アカウント(Receipt)を生成
	_, err = w.GenerateReceiptAccount(seed, 0, 100)
	if err != nil {
		return errors.Errorf("GenerateReceiptAccount() error: %v", err)
	}

	// アカウント(Payment)を生成
	_, err = w.GeneratePaymentAccount(seed, 0, 100)
	if err != nil {
		return errors.Errorf("GeneratePaymentAccount() error: %v", err)
	}

	// アカウント(Authorization)を生成
	_, err = w.GenerateAuthorizationAccount(seed, 0, 3)
	if err != nil {
		return errors.Errorf("GenerateMultisigAccount() error: %v", err)
	}

	return nil
}

// GenerateSeed seedを生成する
func (w *Wallet) GenerateSeed() ([]byte, error) {

	//seed, err := w.DB.GetSeedOne()
	//if err == nil && seed.Seed != "" {
	//	logger.Info("seed have already been generated")
	//	return key.SeedToByte(seed.Seed)
	//}
	bSeed, err := w.retrieveSeed()
	if err == nil {
		return bSeed, nil
	}

	// seed生成
	bSeed, err = key.GenerateSeed()
	if err != nil {
		return nil, errors.Errorf("key.GenerateSeed() error: %s", err)
	}
	strSeed := key.SeedToString(bSeed)

	// DBにseed情報を登録
	_, err = w.DB.InsertSeed(strSeed, nil, true)
	if err != nil {
		return nil, errors.Errorf("key.InsertSeed() error: %s", err)
	}

	return bSeed, nil
}

func (w *Wallet) retrieveSeed() ([]byte, error) {
	// DBからseed情報を登録
	seed, err := w.DB.GetSeedOne()
	if err == nil && seed.Seed != "" {
		logger.Info("seed have already been generated")
		return key.SeedToByte(seed.Seed)
	}

	return nil, errors.Errorf("DB.GetSeedOne() error: %v", err)
}

// generateAccount アカウントを生成する
func (w *Wallet) generateAccount(seed []byte, idxFrom, count uint32, account key.AccountType) ([]key.WalletKey, error) {
	// アカウント(Client)を生成
	priv, _, err := key.CreateAccount(w.BTC.GetChainConf(), seed, account)
	if err != nil {
		return nil, errors.Errorf("key.CreateAccount() error: %s", err)
	}

	walletKeys, err := key.CreateKeysWithIndex(w.BTC.GetChainConf(), priv, idxFrom, count)
	if err != nil {
		return nil, errors.Errorf("key.CreateKeysWithIndex() error: %s", err)
	}

	return walletKeys, nil
}

// GenerateClientAccount Clientアカウントを生成する
func (w *Wallet) GenerateClientAccount(seed []byte, idxFrom, count uint32) ([]key.WalletKey, error) {
	walletKeys, err := w.generateAccount(seed, idxFrom, count, key.AccountTypeClient)
	if err != nil {
		return nil, errors.Errorf("key.generateAccount(AccountTypeClient) error: %s", err)
	}

	// TODO:DBにClientAccountのKey情報を登録

	return walletKeys, err
}

// GenerateReceiptAccount Receiptアカウントを生成する
func (w *Wallet) GenerateReceiptAccount(seed []byte, idxFrom, count uint32) ([]key.WalletKey, error) {
	walletKeys, err := w.generateAccount(seed, idxFrom, count, key.AccountTypeReceipt)
	if err != nil {
		return nil, errors.Errorf("key.generateAccount(AccountTypeReceipt) error: %s", err)
	}

	// TODO:DBにReceiptAccountのKey情報を登録

	return walletKeys, err
}

// GeneratePaymentAccount Paymentアカウントを生成する
func (w *Wallet) GeneratePaymentAccount(seed []byte, idxFrom, count uint32) ([]key.WalletKey, error) {
	walletKeys, err := w.generateAccount(seed, idxFrom, count, key.AccountTypePayment)
	if err != nil {
		return nil, errors.Errorf("key.generateAccount(AccountTypePayment) error: %s", err)
	}

	// TODO:DBにPaymentAccountのKey情報を登録

	return walletKeys, err
}

// GenerateAuthorizationAccount Paymentアカウントを生成する
func (w *Wallet) GenerateAuthorizationAccount(seed []byte, idxFrom, count uint32) ([]key.WalletKey, error) {
	walletKeys, err := w.generateAccount(seed, idxFrom, count, key.AccountTypeAuthorization)
	if err != nil {
		return nil, errors.Errorf("key.generateAccount(AccountTypeAuthorization) error: %s", err)
	}

	// TODO:DBにAuthorizationAccountのKey情報を登録

	return walletKeys, err
}