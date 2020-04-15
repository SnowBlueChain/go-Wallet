package coldwallet

import (
	"fmt"
	"strings"

	"github.com/bookerzzz/grok"
	"github.com/btcsuite/btcd/wire"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/hiromaily/go-bitcoin/pkg/account"
	"github.com/hiromaily/go-bitcoin/pkg/action"
	"github.com/hiromaily/go-bitcoin/pkg/model/rdb/coldrepo"
	"github.com/hiromaily/go-bitcoin/pkg/serial"
	"github.com/hiromaily/go-bitcoin/pkg/tx"
	"github.com/hiromaily/go-bitcoin/pkg/wallets/api/btc"
	"github.com/hiromaily/go-bitcoin/pkg/wallets/types"
)

// sing on unsigned transaction

// SignTx 渡されたファイルからtransactionを読み取り、署名を行う
// TODO:いずれにせよ、入金と出金で署名もMultisigかどうかで変わってくる
func (w *ColdWallet) SignTx(filePath string) (string, bool, string, error) {

	//ファイル名から、tx_receipt_idを取得する
	//payment_5_unsigned_1534466246366489473
	actionType, _, txReceiptID, err := w.txFileRepo.ValidateFilePath(filePath, []tx.TxType{tx.TxTypeUnsigned, tx.TxTypeUnsigned2nd})
	if err != nil {
		return "", false, "", err
	}

	//ファイルからhexを読み取る
	data, err := w.txFileRepo.ReadFile(filePath)
	if err != nil {
		return "", false, "", err
	}

	var hex, encodedAddrsPrevs string

	//encodedPrevTxs
	tmp := strings.Split(data, ",")
	hex = tmp[0]
	if len(tmp) > 1 {
		encodedAddrsPrevs = tmp[1]
	}

	//署名
	hexTx, isSigned, newEncodedAddrsPrevs, err := w.signatureByHex(hex, encodedAddrsPrevs, actionType)
	if err != nil {
		return "", isSigned, "", err
	}

	//ファイルに書き込むデータ
	savedata := hexTx

	//署名が完了していないとき、TxTypeUnsigned2nd
	txType := tx.TxTypeSigned
	if isSigned == false {
		txType = tx.TxTypeUnsigned2nd
		if newEncodedAddrsPrevs != "" {
			savedata = fmt.Sprintf("%s,%s", savedata, newEncodedAddrsPrevs)
		}
	}

	//ファイルに書き込む
	path := w.txFileRepo.CreateFilePath(actionType, txType, txReceiptID)
	generatedFileName, err := w.txFileRepo.WriteFile(path, savedata)
	if err != nil {
		return "", isSigned, "", err
	}

	return hexTx, isSigned, generatedFileName, nil
}

// signatureByHex 署名する
// オフラインで使うことを想定
func (w *ColdWallet) signatureByHex(hex, encodedAddrsPrevs string, actionType action.ActionType) (string, bool, string, error) {
	// Hexからトランザクションを取得
	msgTx, err := w.btc.ToMsgTx(hex)
	if err != nil {
		return "", false, "", err
	}

	// 署名
	var (
		signedTx             *wire.MsgTx
		isSigned             bool
		addrsPrevs           btc.AddrsPrevTxs
		accountKeys          []coldrepo.AccountKeyTable
		wips                 []string
		newEncodedAddrsPrevs string
	)

	if encodedAddrsPrevs == "" {
		//Bitcoin coreのバージョン17から、常に必要
		return "", false, "", errors.New("encodedAddrsPrevs must be set")
	}

	//decodeする
	serial.DecodeFromString(encodedAddrsPrevs, &addrsPrevs)

	//WIPs, RedeedScriptを取得
	//TODO:coldwallet1とcoldwallet2で挙動が違う
	//TODO:receiptの場合、wipsは不要
	//coldwallet2の場合、AccountTypeAuthorizationが必要
	if w.wtype == types.WalletTypeSignature {
		//account_key_authorizationテーブルから情報を取得
		accountKey, err := w.storager.GetOneByMaxIDOnAccountKeyTable(account.AccountTypeAuthorization)
		if err != nil {
			return "", false, "", errors.Errorf("DB.GetOneByMaxIDOnAccountKeyTable() error: %s", err)
		}
		accountKeys = append(accountKeys, *accountKey)
	} else {
		//TODO:ActionTypeが`transfer`の場合、AccountのFromから判別しないといけない。。。
		//=> addrsPrevs.SenderAccount を使うように変更
		//if val, ok := action.ActionToAccountMap[actionType]; ok {
		//	//account_key_payment/account_key_clientテーブルから取得
		//	accountKeys, err = w.DB.GetAllAccountKeyByMultiAddrs(val, addrsPrevs.Addrs)
		//	if err != nil {
		//		return "", false, "", errors.Errorf("DB.GetWIPByMultiAddrs() error: %s", err)
		//	}
		//} else {
		//	return "", false, "", errors.New("[Fatal] actionType can not be retrieved. it should be fixed programmatically")
		//}
		//account_key_payment/account_key_clientテーブルから取得
		accountKeys, err = w.storager.GetAllAccountKeyByMultiAddrs(addrsPrevs.SenderAccount, addrsPrevs.Addrs)
		if err != nil {
			return "", false, "", errors.Errorf("DB.GetWIPByMultiAddrs() error: %s", err)
		}
	}

	//wip
	for _, val := range accountKeys {
		wips = append(wips, val.WalletImportFormat)
	}

	//multisigの場合のみの処理
	//accountType, ok := action.ActionToAccountMap[actionType]
	if account.AccountTypeMultisig[addrsPrevs.SenderAccount] {
		if w.wtype == types.WalletTypeKeyGen {
			//取得したredeemScriptをPrevTxsにマッピング
			for idx, val := range addrsPrevs.Addrs {
				rs := coldrepo.GetRedeedScriptByAddress(accountKeys, val)
				if rs == "" {
					w.logger.Error("redeemScript can not be found")
					continue
				}
				addrsPrevs.PrevTxs[idx].RedeemScript = rs
			}
			grok.Value(addrsPrevs)

			//redeemScriptセット後、シリアライズして戻す
			newEncodedAddrsPrevs, err = serial.EncodeToString(addrsPrevs)
			if err != nil {
				return "", false, "", errors.Errorf("serial.EncodeToString(): error: %s", err)
			}
		} else {
			newEncodedAddrsPrevs = encodedAddrsPrevs
		}
	}

	//署名
	//multisigかどうかで判別
	if account.AccountTypeMultisig[addrsPrevs.SenderAccount] {
		signedTx, isSigned, err = w.btc.SignRawTransactionWithKey(msgTx, wips, addrsPrevs.PrevTxs)
	} else {
		signedTx, isSigned, err = w.btc.SignRawTransaction(msgTx, addrsPrevs.PrevTxs)
	}

	if err != nil {
		return "", false, "", err
	}
	w.logger.Debug(
		"call btc.SignRawTransaction()",
		zap.Bool("isSigned", isSigned))

	hexTx, err := w.btc.ToHex(signedTx)
	if err != nil {
		return "", false, "", errors.Errorf("w.BTC.ToHex(msgTx): error: %s", err)
	}

	return hexTx, isSigned, newEncodedAddrsPrevs, nil

}
