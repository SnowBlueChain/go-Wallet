package wallets

import (
	"database/sql"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/hiromaily/go-bitcoin/pkg/account"
	wtype "github.com/hiromaily/go-bitcoin/pkg/wallet"
	"github.com/hiromaily/go-bitcoin/pkg/wallet/api"
	"github.com/hiromaily/go-bitcoin/pkg/wallet/service/watchsrv"
)

// Watcher is for watch only wallet service interface
type Watcher interface {
	ImportAddress(fileName string, accountType account.AccountType, isRescan bool) error
	CreateDepositTx(adjustmentFee float64) (string, string, error)
	CreatePaymentTx(adjustmentFee float64) (string, string, error)
	CreateTransferTx(sender, receiver account.AccountType, floatAmount, adjustmentFee float64) (string, string, error)
	SendTx(filePath string) (string, error)
	UpdateTxStatus() error

	Done()
	//GetDB() watchrepo.WalletRepositorier // for debug use
	GetBTC() api.Bitcoiner
	GetType() wtype.WalletType
}

// Watch watch only wallet object
type Watch struct {
	btc    api.Bitcoiner
	dbConn *sql.DB
	logger *zap.Logger
	tracer opentracing.Tracer
	watchsrv.AddressImporter
	watchsrv.TxCreator
	watchsrv.TxSender
	watchsrv.TxMonitorer
	wtype wtype.WalletType
}

// NewWatch returns Watch object
func NewWatch(
	btc api.Bitcoiner,
	dbConn *sql.DB,
	logger *zap.Logger,
	tracer opentracing.Tracer,
	addrImporter watchsrv.AddressImporter,
	txCreator watchsrv.TxCreator,
	txSender watchsrv.TxSender,
	txMonitorer watchsrv.TxMonitorer,
	wtype wtype.WalletType) *Watch {

	return &Watch{
		btc:             btc,
		logger:          logger,
		dbConn:          dbConn,
		tracer:          tracer,
		AddressImporter: addrImporter,
		TxCreator:       txCreator,
		TxSender:        txSender,
		TxMonitorer:     txMonitorer,
		wtype:           wtype,
	}
}

// ImportAddress imports address
func (w *Watch) ImportAddress(fileName string, accountType account.AccountType, isRescan bool) error {
	return w.AddressImporter.ImportAddress(fileName, accountType, isRescan)
}

// CreateDepositTx creates deposit unsigned transaction
func (w *Watch) CreateDepositTx(adjustmentFee float64) (string, string, error) {
	return w.TxCreator.CreateDepositTx(adjustmentFee)
}

// CreatePaymentTx creates payment unsigned transaction
func (w *Watch) CreatePaymentTx(adjustmentFee float64) (string, string, error) {
	return w.TxCreator.CreatePaymentTx(adjustmentFee)
}

// CreateTransferTx creates transfer unsigned transaction
func (w *Watch) CreateTransferTx(sender, receiver account.AccountType, floatAmount, adjustmentFee float64) (string, string, error) {
	return w.TxCreator.CreateTransferTx(sender, receiver, floatAmount, adjustmentFee)
}

// UpdateTxStatus updates transaction status
func (w *Watch) UpdateTxStatus() error {
	return w.TxMonitorer.UpdateTxStatus()
}

// SendTx sends signed transaction
func (w *Watch) SendTx(filePath string) (string, error) {
	return w.TxSender.SendTx(filePath)
}

// Done should be called before exit
func (w *Watch) Done() {
	w.dbConn.Close()
	w.btc.Close()
}

// GetDB gets repository
//func (w *Watch) GetDB() watchrepo.WalletRepositorier {
//	return w.
//}

// GetBTC gets btc
func (w *Watch) GetBTC() api.Bitcoiner {
	return w.btc
}

// GetType gets wallet type
func (w *Watch) GetType() wtype.WalletType {
	return w.wtype
}