package main

import (
	"database/sql"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/opentracing/opentracing-go"
	"github.com/volatiletech/sqlboiler/boil"
	"go.uber.org/zap"

	"github.com/hiromaily/go-bitcoin/pkg/address"
	"github.com/hiromaily/go-bitcoin/pkg/config"
	mysql "github.com/hiromaily/go-bitcoin/pkg/db/rdb"
	"github.com/hiromaily/go-bitcoin/pkg/logger"
	"github.com/hiromaily/go-bitcoin/pkg/repository/watchrepo"
	"github.com/hiromaily/go-bitcoin/pkg/tracer"
	"github.com/hiromaily/go-bitcoin/pkg/tx"
	wtype "github.com/hiromaily/go-bitcoin/pkg/wallet"
	"github.com/hiromaily/go-bitcoin/pkg/wallet/api"
	"github.com/hiromaily/go-bitcoin/pkg/wallet/service/watchsrv"
	"github.com/hiromaily/go-bitcoin/pkg/wallet/wallets"
)

// Registry is for registry interface
type Registry interface {
	NewWalleter() wallets.Watcher
}

type registry struct {
	conf        *config.Config
	walletType  wtype.WalletType
	logger      *zap.Logger
	rpcClient   *rpcclient.Client
	btc         api.Bitcoiner
	mysqlClient *sql.DB
}

// NewRegistry is to register registry interface
func NewRegistry(conf *config.Config, walletType wtype.WalletType) Registry {
	return &registry{
		conf:       conf,
		walletType: walletType,
	}
}

// NewWalleter is to register for walleter interface
func (r *registry) NewWalleter() wallets.Watcher {
	return wallets.NewWatch(
		r.newBTC(),
		r.newMySQLClient(),
		r.newLogger(),
		r.newTracer(),
		r.newAddressImporter(),
		r.newTxCreator(),
		r.newTxSender(),
		r.newTxMonitorer(),
		r.newPaymentRequestCreator(),
		r.walletType,
	)
}

func (r *registry) newAddressImporter() watchsrv.AddressImporter {
	return watchsrv.NewAddressImport(
		r.newBTC(),
		r.newLogger(),
		r.newMySQLClient(),
		r.newAddressRepo(),
		r.newAddressFileRepo(),
		r.conf.CoinTypeCode,
		r.walletType,
	)
}

func (r *registry) newTxCreator() watchsrv.TxCreator {
	return watchsrv.NewTxCreate(
		r.newBTC(),
		r.newLogger(),
		r.newMySQLClient(),
		r.newAddressRepo(),
		r.newTxRepo(),
		r.newTxInputRepo(),
		r.newTxOutputRepo(),
		r.newPaymentRequestRepo(),
		r.newTxFileRepo(),
		r.walletType,
	)
}

func (r *registry) newTxSender() watchsrv.TxSender {
	return watchsrv.NewTxSend(
		r.newBTC(),
		r.newLogger(),
		r.newMySQLClient(),
		r.newAddressRepo(),
		r.newTxRepo(),
		r.newTxOutputRepo(),
		r.newTxFileRepo(),
		r.walletType,
	)
}

func (r *registry) newTxMonitorer() watchsrv.TxMonitorer {
	return watchsrv.NewTxMonitor(
		r.newBTC(),
		r.newLogger(),
		r.newMySQLClient(),
		r.newTxRepo(),
		r.newTxInputRepo(),
		r.newPaymentRequestRepo(),
		r.walletType,
	)
}

func (r *registry) newPaymentRequestCreator() watchsrv.PaymentRequestCreator {
	return watchsrv.NewPaymentRequestCreate(
		r.newBTC(),
		r.newLogger(),
		r.newMySQLClient(),
		r.newAddressRepo(),
		r.newPaymentRequestRepo(),
		r.walletType,
	)
}

func (r *registry) newRPCClient() *rpcclient.Client {
	if r.rpcClient == nil {
		var err error
		r.rpcClient, err = api.NewRPCClient(&r.conf.Bitcoin)
		if err != nil {
			panic(err)
		}
	}
	return r.rpcClient
}

func (r *registry) newBTC() api.Bitcoiner {
	if r.btc == nil {
		var err error
		r.btc, err = api.NewBitcoin(
			r.newRPCClient(),
			&r.conf.Bitcoin,
			r.newLogger(),
			r.conf.CoinTypeCode,
		)
		if err != nil {
			panic(err)
		}
	}
	return r.btc
}

func (r *registry) newLogger() *zap.Logger {
	if r.logger == nil {
		r.logger = logger.NewZapLogger(&r.conf.Logger)
	}
	return r.logger
}

func (r *registry) newTracer() opentracing.Tracer {
	return tracer.NewTracer(r.conf.Tracer)
}

func (r *registry) newTxRepo() watchrepo.TxRepositorier {
	return watchrepo.NewTxRepository(
		r.newMySQLClient(),
		r.conf.CoinTypeCode,
		r.newLogger(),
	)
}

func (r *registry) newTxInputRepo() watchrepo.TxInputRepositorier {
	return watchrepo.NewTxInputRepository(
		r.newMySQLClient(),
		r.conf.CoinTypeCode,
		r.newLogger(),
	)
}

func (r *registry) newTxOutputRepo() watchrepo.TxOutputRepositorier {
	return watchrepo.NewTxOutputRepository(
		r.newMySQLClient(),
		r.conf.CoinTypeCode,
		r.newLogger(),
	)
}

func (r *registry) newPaymentRequestRepo() watchrepo.PaymentRequestRepositorier {
	return watchrepo.NewPaymentRequestRepository(
		r.newMySQLClient(),
		r.conf.CoinTypeCode,
		r.newLogger(),
	)
}

func (r *registry) newAddressRepo() watchrepo.AddressRepositorier {
	return watchrepo.NewAddressRepository(
		r.newMySQLClient(),
		r.conf.CoinTypeCode,
		r.newLogger(),
	)
}

func (r *registry) newMySQLClient() *sql.DB {
	if r.mysqlClient == nil {
		dbConn, err := mysql.NewMySQL(&r.conf.MySQL)
		if err != nil {
			panic(err)
		}
		r.mysqlClient = dbConn
	}
	if r.conf.MySQL.Debug {
		boil.DebugMode = true
	}
	return r.mysqlClient
}

func (r *registry) newAddressFileRepo() address.FileRepositorier {
	return address.NewFileRepository(
		r.conf.PubkeyFile.BasePath,
		r.newLogger(),
	)
}

func (r *registry) newTxFileRepo() tx.FileRepositorier {
	return tx.NewFileRepository(
		r.conf.TxFile.BasePath,
		r.newLogger(),
	)
}
