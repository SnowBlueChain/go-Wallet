// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// XRPDetailTX is an object representing the database table.
type XRPDetailTX struct {
	ID                    int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	TXID                  int64     `boil:"tx_id" json:"tx_id" toml:"tx_id" yaml:"tx_id"`
	UUID                  string    `boil:"uuid" json:"uuid" toml:"uuid" yaml:"uuid"`
	CurrentTXType         int8      `boil:"current_tx_type" json:"current_tx_type" toml:"current_tx_type" yaml:"current_tx_type"`
	SenderAccount         string    `boil:"sender_account" json:"sender_account" toml:"sender_account" yaml:"sender_account"`
	SenderAddress         string    `boil:"sender_address" json:"sender_address" toml:"sender_address" yaml:"sender_address"`
	ReceiverAccount       string    `boil:"receiver_account" json:"receiver_account" toml:"receiver_account" yaml:"receiver_account"`
	ReceiverAddress       string    `boil:"receiver_address" json:"receiver_address" toml:"receiver_address" yaml:"receiver_address"`
	Amount                string    `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	XRPTXType             string    `boil:"xrp_tx_type" json:"xrp_tx_type" toml:"xrp_tx_type" yaml:"xrp_tx_type"`
	Fee                   string    `boil:"fee" json:"fee" toml:"fee" yaml:"fee"`
	Flags                 uint64    `boil:"flags" json:"flags" toml:"flags" yaml:"flags"`
	LastLedgerSequence    uint64    `boil:"last_ledger_sequence" json:"last_ledger_sequence" toml:"last_ledger_sequence" yaml:"last_ledger_sequence"`
	Sequence              uint64    `boil:"sequence" json:"sequence" toml:"sequence" yaml:"sequence"`
	SigningPubkey         string    `boil:"signing_pubkey" json:"signing_pubkey" toml:"signing_pubkey" yaml:"signing_pubkey"`
	TXNSignature          string    `boil:"txn_signature" json:"txn_signature" toml:"txn_signature" yaml:"txn_signature"`
	Hash                  string    `boil:"hash" json:"hash" toml:"hash" yaml:"hash"`
	EarliestLedgerVersion uint64    `boil:"earliest_ledger_version" json:"earliest_ledger_version" toml:"earliest_ledger_version" yaml:"earliest_ledger_version"`
	SignedTXID            string    `boil:"signed_tx_id" json:"signed_tx_id" toml:"signed_tx_id" yaml:"signed_tx_id"`
	TXBlob                string    `boil:"tx_blob" json:"tx_blob" toml:"tx_blob" yaml:"tx_blob"`
	SentUpdatedAt         null.Time `boil:"sent_updated_at" json:"sent_updated_at,omitempty" toml:"sent_updated_at" yaml:"sent_updated_at,omitempty"`

	R *xrpDetailTXR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L xrpDetailTXL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var XRPDetailTXColumns = struct {
	ID                    string
	TXID                  string
	UUID                  string
	CurrentTXType         string
	SenderAccount         string
	SenderAddress         string
	ReceiverAccount       string
	ReceiverAddress       string
	Amount                string
	XRPTXType             string
	Fee                   string
	Flags                 string
	LastLedgerSequence    string
	Sequence              string
	SigningPubkey         string
	TXNSignature          string
	Hash                  string
	EarliestLedgerVersion string
	SignedTXID            string
	TXBlob                string
	SentUpdatedAt         string
}{
	ID:                    "id",
	TXID:                  "tx_id",
	UUID:                  "uuid",
	CurrentTXType:         "current_tx_type",
	SenderAccount:         "sender_account",
	SenderAddress:         "sender_address",
	ReceiverAccount:       "receiver_account",
	ReceiverAddress:       "receiver_address",
	Amount:                "amount",
	XRPTXType:             "xrp_tx_type",
	Fee:                   "fee",
	Flags:                 "flags",
	LastLedgerSequence:    "last_ledger_sequence",
	Sequence:              "sequence",
	SigningPubkey:         "signing_pubkey",
	TXNSignature:          "txn_signature",
	Hash:                  "hash",
	EarliestLedgerVersion: "earliest_ledger_version",
	SignedTXID:            "signed_tx_id",
	TXBlob:                "tx_blob",
	SentUpdatedAt:         "sent_updated_at",
}

// Generated where

var XRPDetailTXWhere = struct {
	ID                    whereHelperint64
	TXID                  whereHelperint64
	UUID                  whereHelperstring
	CurrentTXType         whereHelperint8
	SenderAccount         whereHelperstring
	SenderAddress         whereHelperstring
	ReceiverAccount       whereHelperstring
	ReceiverAddress       whereHelperstring
	Amount                whereHelperstring
	XRPTXType             whereHelperstring
	Fee                   whereHelperstring
	Flags                 whereHelperuint64
	LastLedgerSequence    whereHelperuint64
	Sequence              whereHelperuint64
	SigningPubkey         whereHelperstring
	TXNSignature          whereHelperstring
	Hash                  whereHelperstring
	EarliestLedgerVersion whereHelperuint64
	SignedTXID            whereHelperstring
	TXBlob                whereHelperstring
	SentUpdatedAt         whereHelpernull_Time
}{
	ID:                    whereHelperint64{field: "`xrp_detail_tx`.`id`"},
	TXID:                  whereHelperint64{field: "`xrp_detail_tx`.`tx_id`"},
	UUID:                  whereHelperstring{field: "`xrp_detail_tx`.`uuid`"},
	CurrentTXType:         whereHelperint8{field: "`xrp_detail_tx`.`current_tx_type`"},
	SenderAccount:         whereHelperstring{field: "`xrp_detail_tx`.`sender_account`"},
	SenderAddress:         whereHelperstring{field: "`xrp_detail_tx`.`sender_address`"},
	ReceiverAccount:       whereHelperstring{field: "`xrp_detail_tx`.`receiver_account`"},
	ReceiverAddress:       whereHelperstring{field: "`xrp_detail_tx`.`receiver_address`"},
	Amount:                whereHelperstring{field: "`xrp_detail_tx`.`amount`"},
	XRPTXType:             whereHelperstring{field: "`xrp_detail_tx`.`xrp_tx_type`"},
	Fee:                   whereHelperstring{field: "`xrp_detail_tx`.`fee`"},
	Flags:                 whereHelperuint64{field: "`xrp_detail_tx`.`flags`"},
	LastLedgerSequence:    whereHelperuint64{field: "`xrp_detail_tx`.`last_ledger_sequence`"},
	Sequence:              whereHelperuint64{field: "`xrp_detail_tx`.`sequence`"},
	SigningPubkey:         whereHelperstring{field: "`xrp_detail_tx`.`signing_pubkey`"},
	TXNSignature:          whereHelperstring{field: "`xrp_detail_tx`.`txn_signature`"},
	Hash:                  whereHelperstring{field: "`xrp_detail_tx`.`hash`"},
	EarliestLedgerVersion: whereHelperuint64{field: "`xrp_detail_tx`.`earliest_ledger_version`"},
	SignedTXID:            whereHelperstring{field: "`xrp_detail_tx`.`signed_tx_id`"},
	TXBlob:                whereHelperstring{field: "`xrp_detail_tx`.`tx_blob`"},
	SentUpdatedAt:         whereHelpernull_Time{field: "`xrp_detail_tx`.`sent_updated_at`"},
}

// XRPDetailTXRels is where relationship names are stored.
var XRPDetailTXRels = struct {
}{}

// xrpDetailTXR is where relationships are stored.
type xrpDetailTXR struct {
}

// NewStruct creates a new relationship struct
func (*xrpDetailTXR) NewStruct() *xrpDetailTXR {
	return &xrpDetailTXR{}
}

// xrpDetailTXL is where Load methods for each relationship are stored.
type xrpDetailTXL struct{}

var (
	xrpDetailTXAllColumns            = []string{"id", "tx_id", "uuid", "current_tx_type", "sender_account", "sender_address", "receiver_account", "receiver_address", "amount", "xrp_tx_type", "fee", "flags", "last_ledger_sequence", "sequence", "signing_pubkey", "txn_signature", "hash", "earliest_ledger_version", "signed_tx_id", "tx_blob", "sent_updated_at"}
	xrpDetailTXColumnsWithoutDefault = []string{"tx_id", "uuid", "sender_account", "sender_address", "receiver_account", "receiver_address", "amount", "xrp_tx_type", "fee", "flags", "last_ledger_sequence", "sequence", "signing_pubkey", "txn_signature", "hash", "earliest_ledger_version", "signed_tx_id", "tx_blob", "sent_updated_at"}
	xrpDetailTXColumnsWithDefault    = []string{"id", "current_tx_type"}
	xrpDetailTXPrimaryKeyColumns     = []string{"id"}
)

type (
	// XRPDetailTXSlice is an alias for a slice of pointers to XRPDetailTX.
	// This should generally be used opposed to []XRPDetailTX.
	XRPDetailTXSlice []*XRPDetailTX

	xrpDetailTXQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	xrpDetailTXType                 = reflect.TypeOf(&XRPDetailTX{})
	xrpDetailTXMapping              = queries.MakeStructMapping(xrpDetailTXType)
	xrpDetailTXPrimaryKeyMapping, _ = queries.BindMapping(xrpDetailTXType, xrpDetailTXMapping, xrpDetailTXPrimaryKeyColumns)
	xrpDetailTXInsertCacheMut       sync.RWMutex
	xrpDetailTXInsertCache          = make(map[string]insertCache)
	xrpDetailTXUpdateCacheMut       sync.RWMutex
	xrpDetailTXUpdateCache          = make(map[string]updateCache)
	xrpDetailTXUpsertCacheMut       sync.RWMutex
	xrpDetailTXUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single xrpDetailTX record from the query.
func (q xrpDetailTXQuery) One(ctx context.Context, exec boil.ContextExecutor) (*XRPDetailTX, error) {
	o := &XRPDetailTX{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for xrp_detail_tx")
	}

	return o, nil
}

// All returns all XRPDetailTX records from the query.
func (q xrpDetailTXQuery) All(ctx context.Context, exec boil.ContextExecutor) (XRPDetailTXSlice, error) {
	var o []*XRPDetailTX

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to XRPDetailTX slice")
	}

	return o, nil
}

// Count returns the count of all XRPDetailTX records in the query.
func (q xrpDetailTXQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count xrp_detail_tx rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q xrpDetailTXQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if xrp_detail_tx exists")
	}

	return count > 0, nil
}

// XRPDetailTxes retrieves all the records using an executor.
func XRPDetailTxes(mods ...qm.QueryMod) xrpDetailTXQuery {
	mods = append(mods, qm.From("`xrp_detail_tx`"))
	return xrpDetailTXQuery{NewQuery(mods...)}
}

// FindXRPDetailTX retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindXRPDetailTX(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*XRPDetailTX, error) {
	xrpDetailTXObj := &XRPDetailTX{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `xrp_detail_tx` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, xrpDetailTXObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from xrp_detail_tx")
	}

	return xrpDetailTXObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *XRPDetailTX) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no xrp_detail_tx provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(xrpDetailTXColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	xrpDetailTXInsertCacheMut.RLock()
	cache, cached := xrpDetailTXInsertCache[key]
	xrpDetailTXInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			xrpDetailTXAllColumns,
			xrpDetailTXColumnsWithDefault,
			xrpDetailTXColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(xrpDetailTXType, xrpDetailTXMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(xrpDetailTXType, xrpDetailTXMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `xrp_detail_tx` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `xrp_detail_tx` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `xrp_detail_tx` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, xrpDetailTXPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into xrp_detail_tx")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == xrpDetailTXMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for xrp_detail_tx")
	}

CacheNoHooks:
	if !cached {
		xrpDetailTXInsertCacheMut.Lock()
		xrpDetailTXInsertCache[key] = cache
		xrpDetailTXInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the XRPDetailTX.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *XRPDetailTX) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	xrpDetailTXUpdateCacheMut.RLock()
	cache, cached := xrpDetailTXUpdateCache[key]
	xrpDetailTXUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			xrpDetailTXAllColumns,
			xrpDetailTXPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update xrp_detail_tx, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `xrp_detail_tx` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, xrpDetailTXPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(xrpDetailTXType, xrpDetailTXMapping, append(wl, xrpDetailTXPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update xrp_detail_tx row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for xrp_detail_tx")
	}

	if !cached {
		xrpDetailTXUpdateCacheMut.Lock()
		xrpDetailTXUpdateCache[key] = cache
		xrpDetailTXUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q xrpDetailTXQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for xrp_detail_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for xrp_detail_tx")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o XRPDetailTXSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), xrpDetailTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `xrp_detail_tx` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, xrpDetailTXPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in xrpDetailTX slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all xrpDetailTX")
	}
	return rowsAff, nil
}

var mySQLXRPDetailTXUniqueColumns = []string{
	"id",
	"uuid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *XRPDetailTX) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no xrp_detail_tx provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(xrpDetailTXColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLXRPDetailTXUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	xrpDetailTXUpsertCacheMut.RLock()
	cache, cached := xrpDetailTXUpsertCache[key]
	xrpDetailTXUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			xrpDetailTXAllColumns,
			xrpDetailTXColumnsWithDefault,
			xrpDetailTXColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			xrpDetailTXAllColumns,
			xrpDetailTXPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert xrp_detail_tx, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`xrp_detail_tx`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `xrp_detail_tx` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(xrpDetailTXType, xrpDetailTXMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(xrpDetailTXType, xrpDetailTXMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for xrp_detail_tx")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == xrpDetailTXMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(xrpDetailTXType, xrpDetailTXMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for xrp_detail_tx")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for xrp_detail_tx")
	}

CacheNoHooks:
	if !cached {
		xrpDetailTXUpsertCacheMut.Lock()
		xrpDetailTXUpsertCache[key] = cache
		xrpDetailTXUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single XRPDetailTX record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *XRPDetailTX) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no XRPDetailTX provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), xrpDetailTXPrimaryKeyMapping)
	sql := "DELETE FROM `xrp_detail_tx` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from xrp_detail_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for xrp_detail_tx")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q xrpDetailTXQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no xrpDetailTXQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from xrp_detail_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for xrp_detail_tx")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o XRPDetailTXSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), xrpDetailTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `xrp_detail_tx` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, xrpDetailTXPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from xrpDetailTX slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for xrp_detail_tx")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *XRPDetailTX) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindXRPDetailTX(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *XRPDetailTXSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := XRPDetailTXSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), xrpDetailTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `xrp_detail_tx`.* FROM `xrp_detail_tx` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, xrpDetailTXPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in XRPDetailTXSlice")
	}

	*o = slice

	return nil
}

// XRPDetailTXExists checks if the XRPDetailTX row exists.
func XRPDetailTXExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `xrp_detail_tx` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if xrp_detail_tx exists")
	}

	return exists, nil
}

// InsertAll inserts all rows with the specified column values, using an executor.
func (o XRPDetailTXSlice) InsertAll(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	var sql string
	vals := []interface{}{}
	for i, row := range o {

		nzDefaults := queries.NonZeroDefaultSet(xrpDetailTXColumnsWithDefault, row)
		wl, _ := columns.InsertColumnSet(
			xrpDetailTXAllColumns,
			xrpDetailTXColumnsWithDefault,
			xrpDetailTXColumnsWithoutDefault,
			nzDefaults,
		)
		if i == 0 {
			sql = "INSERT INTO `xrp_detail_tx` " + "(`" + strings.Join(wl, "`,`") + "`)" + " VALUES "
		}
		sql += strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), len(vals)+1, len(wl))
		if i != len(o)-1 {
			sql += ","
		}
		valMapping, err := queries.BindMapping(xrpDetailTXType, xrpDetailTXMapping, wl)
		if err != nil {
			return err
		}
		value := reflect.Indirect(reflect.ValueOf(row))
		vals = append(vals, queries.ValuesFromMapping(value, valMapping)...)
	}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, vals...)
	}

	_, err := exec.ExecContext(ctx, sql, vals...)
	if err != nil {
		return errors.Wrap(err, "models: unable to insert into xrp_detail_tx")
	}

	return nil
}
