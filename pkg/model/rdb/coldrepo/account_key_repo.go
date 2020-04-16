package coldrepo

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/hiromaily/go-bitcoin/pkg/account"
	"github.com/hiromaily/go-bitcoin/pkg/address"
)

// AccountKeyTable account_key_clientテーブル
type AccountKeyTable struct {
	ID                    int64      `db:"id"`
	WalletAddress         string     `db:"wallet_address"`
	P2shSegwitAddress     string     `db:"p2sh_segwit_address"`
	FullPublicKey         string     `db:"full_public_key"`
	WalletMultisigAddress string     `db:"wallet_multisig_address"`
	RedeemScript          string     `db:"redeem_script"`
	WalletImportFormat    string     `db:"wallet_import_format"`
	Account               string     `db:"account"`
	Idx                   uint32     `db:"idx"`
	AddrStatus            uint8      `db:"addr_status"`
	UpdatedAt             *time.Time `db:"updated_at"`
}

var accountKeyTableName = map[account.AccountType]string{
	account.AccountTypeClient:        "account_key_client",
	account.AccountTypeReceipt:       "account_key_receipt",
	account.AccountTypePayment:       "account_key_payment",
	account.AccountTypeFee:           "account_key_fee",
	account.AccountTypeStored:        "account_key_stored",
	account.AccountTypeAuthorization: "account_key_authorization",
}

//getMaxIndexOnAccountKeyTable indexの最大値を返す
func (r *ColdRepository) getMaxIndexOnAccountKeyTable(tbl string) (int64, error) {
	sql := "SELECT MAX(idx) from %s;"
	sql = fmt.Sprintf(sql, tbl)
	//logger.Debugf("sql: %s", sql)

	var idx int64
	err := r.db.Get(&idx, sql)

	return idx, err
}

//GetMaxIndexOnAccountKeyTable indexの最大値を返す
func (r *ColdRepository) GetMaxIndexOnAccountKeyTable(accountType account.AccountType) (int64, error) {
	return r.getMaxIndexOnAccountKeyTable(accountKeyTableName[accountType])
}

//getOneByMaxIDOnAccountKeyTable idが最大の1レコードを返す
func (r *ColdRepository) getOneByMaxIDOnAccountKeyTable(tbl string, accountType account.AccountType) (*AccountKeyTable, error) {
	sql := "SELECT * FROM %s ORDER BY ID DESC LIMIT 1;"
	sql = fmt.Sprintf(sql, tbl)
	//logger.Debugf("sql: %s", sql)

	var accountKeyTable AccountKeyTable
	err := r.db.Get(&accountKeyTable, sql)
	if err != nil {
		return nil, err
	}

	return &accountKeyTable, nil
}

//GetOneByMaxIDOnAccountKeyTable idが最大の1レコードを返す
func (r *ColdRepository) GetOneByMaxIDOnAccountKeyTable(accountType account.AccountType) (*AccountKeyTable, error) {
	return r.getOneByMaxIDOnAccountKeyTable(accountKeyTableName[accountType], accountType)
}

// getAllAccountKeyByAddrStatus 指定したaddrStatusのレコードをすべて返す
func (r *ColdRepository) getAllAccountKeyByAddrStatus(tbl string, addrStatus address.AddrStatus) ([]AccountKeyTable, error) {
	//sql := "SELECT * FROM %s WHERE is_imported_priv_key=false;"
	sql := "SELECT * FROM %s WHERE addr_status=?;"
	sql = fmt.Sprintf(sql, tbl)
	//logger.Debugf("sql: %s", sql)

	var accountKeyTable []AccountKeyTable
	err := r.db.Select(&accountKeyTable, sql, address.AddrStatusValue[addrStatus])
	if err != nil {
		return nil, err
	}

	return accountKeyTable, nil
}

// GetAllAccountKeyByAddrStatus 指定したaddrStatusのレコードをすべて返す
func (r *ColdRepository) GetAllAccountKeyByAddrStatus(accountType account.AccountType, addrStatus address.AddrStatus) ([]AccountKeyTable, error) {
	return r.getAllAccountKeyByAddrStatus(accountKeyTableName[accountType], addrStatus)
}

func (r *ColdRepository) getAllAccountKeyByMultiAddrs(tbl string, addrs []string) ([]AccountKeyTable, error) {
	sql := "SELECT * FROM %s WHERE wallet_multisig_address IN (?);"
	sql = fmt.Sprintf(sql, tbl)
	//logger.Debugf("sql: %s", sql)

	//In対応
	query, args, err := sqlx.In(sql, addrs)
	if err != nil {
		return nil, errors.Errorf("sqlx.In() error: %v", err)
	}
	query = r.db.Rebind(query)
	//logger.Debugf("sql: %s", query)

	var accountKeyTable []AccountKeyTable
	err = r.db.Select(&accountKeyTable, query, args...)
	if err != nil {
		return nil, err
	}

	return accountKeyTable, nil
}

// GetAllAccountKeyByMultiAddrs WIPをmultiAddressから取得する
func (r *ColdRepository) GetAllAccountKeyByMultiAddrs(accountType account.AccountType, addrs []string) ([]AccountKeyTable, error) {
	return r.getAllAccountKeyByMultiAddrs(accountKeyTableName[accountType], addrs)
}

// insertAccountKeyClient account_key_table(client, payment, receipt...)テーブルにレコードを作成する
//TODO:BulkInsertがやりたい
func (r *ColdRepository) insertAccountKeyTable(tbl string, accountKeyTables []AccountKeyTable, tx *sqlx.Tx, isCommit bool) error {

	sql := `
INSERT INTO %s (wallet_address, p2sh_segwit_address, full_public_key, wallet_multisig_address, redeem_script, wallet_import_format, account, idx) 
VALUES (:wallet_address, :p2sh_segwit_address,:full_public_key, :wallet_multisig_address, :redeem_script, :wallet_import_format, :account, :idx)
`
	sql = fmt.Sprintf(sql, tbl)
	//logger.Debugf("sql: %s", sql)

	if tx == nil {
		tx = r.db.MustBegin()
	}

	for _, accountKeyClient := range accountKeyTables {
		_, err := tx.NamedExec(sql, accountKeyClient)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if isCommit {
		tx.Commit()
	}

	return nil
}

// InsertAccountKeyTable account_key_table(client, payment, receipt...)テーブルにレコードを作成する
func (r *ColdRepository) InsertAccountKeyTable(accountType account.AccountType, accountKeyTables []AccountKeyTable, tx *sqlx.Tx, isCommit bool) error {
	return r.insertAccountKeyTable(accountKeyTableName[accountType], accountKeyTables, tx, isCommit)
}

// updateAddrStatus addr_statusを更新する
func (r *ColdRepository) updateAddrStatusByWIF(tbl string, addrStatus address.AddrStatus, strWIF string, tx *sqlx.Tx, isCommit bool) (int64, error) {
	sql := `
UPDATE %s SET addr_status=? WHERE wallet_import_format=?
`
	sql = fmt.Sprintf(sql, tbl)
	//logger.Debugf("sql: %s", sql)

	if tx == nil {
		tx = r.db.MustBegin()
	}

	res, err := tx.Exec(sql, address.AddrStatusValue[addrStatus], strWIF)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if isCommit {
		tx.Commit()
	}
	affectedNum, _ := res.RowsAffected()

	return affectedNum, nil
}

// UpdateAddrStatusByWIF addr_statusを更新する
func (r *ColdRepository) UpdateAddrStatusByWIF(accountType account.AccountType, addrStatus address.AddrStatus, strWIF string, tx *sqlx.Tx, isCommit bool) (int64, error) {
	return r.updateAddrStatusByWIF(accountKeyTableName[accountType], addrStatus, strWIF, tx, isCommit)
}

// updateAddrStatusByWIFs addr_statusを更新する
func (r *ColdRepository) updateAddrStatusByWIFs(tbl string, addrStatus address.AddrStatus, wifs []string, tx *sqlx.Tx, isCommit bool) (int64, error) {
	var sql string
	sql = "UPDATE %s SET addr_status=%d WHERE wallet_import_format IN (?);"
	sql = fmt.Sprintf(sql, tbl, address.AddrStatusValue[addrStatus])

	//In対応
	query, args, err := sqlx.In(sql, wifs)
	if err != nil {
		return 0, errors.Errorf("sqlx.In() error: %v", err)
	}
	query = r.db.Rebind(query)
	//logger.Debugf("sql: %s", query)

	if tx == nil {
		tx = r.db.MustBegin()
	}

	res, err := tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if isCommit {
		tx.Commit()
	}
	affectedNum, _ := res.RowsAffected()

	return affectedNum, nil
}

// UpdateAddrStatusByWIFs addr_statusを更新する
func (r *ColdRepository) UpdateAddrStatusByWIFs(accountType account.AccountType, addrStatus address.AddrStatus, wifs []string, tx *sqlx.Tx, isCommit bool) (int64, error) {
	return r.updateAddrStatusByWIFs(accountKeyTableName[accountType], addrStatus, wifs, tx, isCommit)
}

// updateMultisigAddrOnAccountKeyTableByFullPubKey wallet_multisig_addressを更新する
func (r *ColdRepository) updateMultisigAddrOnAccountKeyTableByFullPubKey(tbl string, accountKeyTable []AccountKeyTable, tx *sqlx.Tx, isCommit bool) error {
	sql := `
UPDATE %s SET wallet_multisig_address=:wallet_multisig_address, redeem_script=:redeem_script, addr_status=:addr_status, updated_at=:updated_at 
WHERE full_public_key=:full_public_key
`
	sql = fmt.Sprintf(sql, tbl)
	//logger.Debugf("sql: %s", sql)

	if tx == nil {
		tx = r.db.MustBegin()
	}

	for _, accountKey := range accountKeyTable {
		_, err := tx.NamedExec(sql, accountKey)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if isCommit {
		tx.Commit()
	}

	return nil
}

// UpdateMultisigAddrOnAccountKeyTableByFullPubKey wallet_multisig_addressを更新する
func (r *ColdRepository) UpdateMultisigAddrOnAccountKeyTableByFullPubKey(accountType account.AccountType, accountKeyTable []AccountKeyTable, tx *sqlx.Tx, isCommit bool) error {
	return r.updateMultisigAddrOnAccountKeyTableByFullPubKey(accountKeyTableName[accountType], accountKeyTable, tx, isCommit)
}

// GetRedeedScriptByAddress 与えられたmultiSigアドレスから、RedeemScriptを取得する
func GetRedeedScriptByAddress(accountKeys []AccountKeyTable, addr string) string {
	for _, val := range accountKeys {
		if val.WalletMultisigAddress == addr {
			return val.RedeemScript
		}
	}
	return ""
}
