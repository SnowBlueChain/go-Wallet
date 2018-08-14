package api

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/hiromaily/go-bitcoin/pkg/toml"
)

// Bitcoin includes Client to call Json-RPC
type Bitcoin struct {
	client            *rpcclient.Client
	chainConf         *chaincfg.Params
	storedAddr        string
	confirmationBlock int64
}

// Connection is to local bitcoin core RPC server using HTTP POST mode
//func Connection(host, user, pass string, postMode, tls, isMain bool) (*Bitcoin, error) {
func Connection(conf *toml.BitcoinConf) (*Bitcoin, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         conf.Host,
		User:         conf.User,
		Pass:         conf.Pass,
		HTTPPostMode: conf.PostMode,   // Bitcoin core only supports HTTP POST mode
		DisableTLS:   conf.DisableTLS, // Bitcoin core does not provide TLS by default
	}

	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}

	bit := Bitcoin{client: client}
	if conf.IsMain {
		bit.chainConf = &chaincfg.MainNetParams
	} else {
		bit.chainConf = &chaincfg.TestNet3Params
	}

	bit.storedAddr = conf.Addr.Stored
	bit.confirmationBlock = conf.Block.ConfirmationNum

	return &bit, err
}

// Close コネクションを切断する
func (b *Bitcoin) Close() {
	b.client.Shutdown()
}

// GetChainConf 接続先であるMainNet/TestNetに応じて必要なconfを返す
func (b *Bitcoin) GetChainConf() *chaincfg.Params {
	return b.chainConf
}

// Client clientオブジェクトを返す
func (b *Bitcoin) Client() *rpcclient.Client {
	return b.client
}

// StoreAddr 保管用アドレスを返す
func (b *Bitcoin) StoreAddr() string {
	return b.storedAddr
}

// ConfirmationBlock Confirmationに必要なブロック数を返す
func (b *Bitcoin) ConfirmationBlock() int64 {
	return b.confirmationBlock
}
