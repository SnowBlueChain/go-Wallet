// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// BTCTX is an object representing the database table.
type BTCTX struct {
	ID                int64         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Coin              string        `boil:"coin" json:"coin" toml:"coin" yaml:"coin"`
	Action            string        `boil:"action" json:"action" toml:"action" yaml:"action"`
	UnsignedHexTX     string        `boil:"unsigned_hex_tx" json:"unsigned_hex_tx" toml:"unsigned_hex_tx" yaml:"unsigned_hex_tx"`
	SignedHexTX       string        `boil:"signed_hex_tx" json:"signed_hex_tx" toml:"signed_hex_tx" yaml:"signed_hex_tx"`
	SentHashTX        string        `boil:"sent_hash_tx" json:"sent_hash_tx" toml:"sent_hash_tx" yaml:"sent_hash_tx"`
	TotalInputAmount  types.Decimal `boil:"total_input_amount" json:"total_input_amount" toml:"total_input_amount" yaml:"total_input_amount"`
	TotalOutputAmount types.Decimal `boil:"total_output_amount" json:"total_output_amount" toml:"total_output_amount" yaml:"total_output_amount"`
	Fee               types.Decimal `boil:"fee" json:"fee" toml:"fee" yaml:"fee"`
	CurrentTXType     int8          `boil:"current_tx_type" json:"current_tx_type" toml:"current_tx_type" yaml:"current_tx_type"`
	UnsignedUpdatedAt null.Time     `boil:"unsigned_updated_at" json:"unsigned_updated_at,omitempty" toml:"unsigned_updated_at" yaml:"unsigned_updated_at,omitempty"`
	SentUpdatedAt     null.Time     `boil:"sent_updated_at" json:"sent_updated_at,omitempty" toml:"sent_updated_at" yaml:"sent_updated_at,omitempty"`

	R *btcTXR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L btcTXL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BTCTXColumns = struct {
	ID                string
	Coin              string
	Action            string
	UnsignedHexTX     string
	SignedHexTX       string
	SentHashTX        string
	TotalInputAmount  string
	TotalOutputAmount string
	Fee               string
	CurrentTXType     string
	UnsignedUpdatedAt string
	SentUpdatedAt     string
}{
	ID:                "id",
	Coin:              "coin",
	Action:            "action",
	UnsignedHexTX:     "unsigned_hex_tx",
	SignedHexTX:       "signed_hex_tx",
	SentHashTX:        "sent_hash_tx",
	TotalInputAmount:  "total_input_amount",
	TotalOutputAmount: "total_output_amount",
	Fee:               "fee",
	CurrentTXType:     "current_tx_type",
	UnsignedUpdatedAt: "unsigned_updated_at",
	SentUpdatedAt:     "sent_updated_at",
}

// Generated where

type whereHelpertypes_Decimal struct{ field string }

func (w whereHelpertypes_Decimal) EQ(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_Decimal) NEQ(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_Decimal) LT(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Decimal) LTE(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Decimal) GT(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Decimal) GTE(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var BTCTXWhere = struct {
	ID                whereHelperint64
	Coin              whereHelperstring
	Action            whereHelperstring
	UnsignedHexTX     whereHelperstring
	SignedHexTX       whereHelperstring
	SentHashTX        whereHelperstring
	TotalInputAmount  whereHelpertypes_Decimal
	TotalOutputAmount whereHelpertypes_Decimal
	Fee               whereHelpertypes_Decimal
	CurrentTXType     whereHelperint8
	UnsignedUpdatedAt whereHelpernull_Time
	SentUpdatedAt     whereHelpernull_Time
}{
	ID:                whereHelperint64{field: "`btc_tx`.`id`"},
	Coin:              whereHelperstring{field: "`btc_tx`.`coin`"},
	Action:            whereHelperstring{field: "`btc_tx`.`action`"},
	UnsignedHexTX:     whereHelperstring{field: "`btc_tx`.`unsigned_hex_tx`"},
	SignedHexTX:       whereHelperstring{field: "`btc_tx`.`signed_hex_tx`"},
	SentHashTX:        whereHelperstring{field: "`btc_tx`.`sent_hash_tx`"},
	TotalInputAmount:  whereHelpertypes_Decimal{field: "`btc_tx`.`total_input_amount`"},
	TotalOutputAmount: whereHelpertypes_Decimal{field: "`btc_tx`.`total_output_amount`"},
	Fee:               whereHelpertypes_Decimal{field: "`btc_tx`.`fee`"},
	CurrentTXType:     whereHelperint8{field: "`btc_tx`.`current_tx_type`"},
	UnsignedUpdatedAt: whereHelpernull_Time{field: "`btc_tx`.`unsigned_updated_at`"},
	SentUpdatedAt:     whereHelpernull_Time{field: "`btc_tx`.`sent_updated_at`"},
}

// BTCTXRels is where relationship names are stored.
var BTCTXRels = struct {
}{}

// btcTXR is where relationships are stored.
type btcTXR struct {
}

// NewStruct creates a new relationship struct
func (*btcTXR) NewStruct() *btcTXR {
	return &btcTXR{}
}

// btcTXL is where Load methods for each relationship are stored.
type btcTXL struct{}

var (
	btcTXAllColumns            = []string{"id", "coin", "action", "unsigned_hex_tx", "signed_hex_tx", "sent_hash_tx", "total_input_amount", "total_output_amount", "fee", "current_tx_type", "unsigned_updated_at", "sent_updated_at"}
	btcTXColumnsWithoutDefault = []string{"coin", "action", "unsigned_hex_tx", "signed_hex_tx", "sent_hash_tx", "total_input_amount", "total_output_amount", "fee", "sent_updated_at"}
	btcTXColumnsWithDefault    = []string{"id", "current_tx_type", "unsigned_updated_at"}
	btcTXPrimaryKeyColumns     = []string{"id"}
)

type (
	// BTCTXSlice is an alias for a slice of pointers to BTCTX.
	// This should generally be used opposed to []BTCTX.
	BTCTXSlice []*BTCTX

	btcTXQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	btcTXType                 = reflect.TypeOf(&BTCTX{})
	btcTXMapping              = queries.MakeStructMapping(btcTXType)
	btcTXPrimaryKeyMapping, _ = queries.BindMapping(btcTXType, btcTXMapping, btcTXPrimaryKeyColumns)
	btcTXInsertCacheMut       sync.RWMutex
	btcTXInsertCache          = make(map[string]insertCache)
	btcTXUpdateCacheMut       sync.RWMutex
	btcTXUpdateCache          = make(map[string]updateCache)
	btcTXUpsertCacheMut       sync.RWMutex
	btcTXUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single btcTX record from the query.
func (q btcTXQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BTCTX, error) {
	o := &BTCTX{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for btc_tx")
	}

	return o, nil
}

// All returns all BTCTX records from the query.
func (q btcTXQuery) All(ctx context.Context, exec boil.ContextExecutor) (BTCTXSlice, error) {
	var o []*BTCTX

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BTCTX slice")
	}

	return o, nil
}

// Count returns the count of all BTCTX records in the query.
func (q btcTXQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count btc_tx rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q btcTXQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if btc_tx exists")
	}

	return count > 0, nil
}

// BTCTxes retrieves all the records using an executor.
func BTCTxes(mods ...qm.QueryMod) btcTXQuery {
	mods = append(mods, qm.From("`btc_tx`"))
	return btcTXQuery{NewQuery(mods...)}
}

// FindBTCTX retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBTCTX(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*BTCTX, error) {
	btcTXObj := &BTCTX{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `btc_tx` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, btcTXObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from btc_tx")
	}

	return btcTXObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BTCTX) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no btc_tx provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(btcTXColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	btcTXInsertCacheMut.RLock()
	cache, cached := btcTXInsertCache[key]
	btcTXInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			btcTXAllColumns,
			btcTXColumnsWithDefault,
			btcTXColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(btcTXType, btcTXMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(btcTXType, btcTXMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `btc_tx` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `btc_tx` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `btc_tx` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, btcTXPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into btc_tx")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == btcTXMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for btc_tx")
	}

CacheNoHooks:
	if !cached {
		btcTXInsertCacheMut.Lock()
		btcTXInsertCache[key] = cache
		btcTXInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the BTCTX.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BTCTX) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	btcTXUpdateCacheMut.RLock()
	cache, cached := btcTXUpdateCache[key]
	btcTXUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			btcTXAllColumns,
			btcTXPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update btc_tx, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `btc_tx` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, btcTXPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(btcTXType, btcTXMapping, append(wl, btcTXPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update btc_tx row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for btc_tx")
	}

	if !cached {
		btcTXUpdateCacheMut.Lock()
		btcTXUpdateCache[key] = cache
		btcTXUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q btcTXQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for btc_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for btc_tx")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BTCTXSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), btcTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `btc_tx` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, btcTXPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in btcTX slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all btcTX")
	}
	return rowsAff, nil
}

var mySQLBTCTXUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BTCTX) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no btc_tx provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(btcTXColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBTCTXUniqueColumns, o)

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

	btcTXUpsertCacheMut.RLock()
	cache, cached := btcTXUpsertCache[key]
	btcTXUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			btcTXAllColumns,
			btcTXColumnsWithDefault,
			btcTXColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			btcTXAllColumns,
			btcTXPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert btc_tx, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "btc_tx", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `btc_tx` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(btcTXType, btcTXMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(btcTXType, btcTXMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for btc_tx")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == btcTXMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(btcTXType, btcTXMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for btc_tx")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for btc_tx")
	}

CacheNoHooks:
	if !cached {
		btcTXUpsertCacheMut.Lock()
		btcTXUpsertCache[key] = cache
		btcTXUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single BTCTX record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BTCTX) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BTCTX provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), btcTXPrimaryKeyMapping)
	sql := "DELETE FROM `btc_tx` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from btc_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for btc_tx")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q btcTXQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no btcTXQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from btc_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for btc_tx")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BTCTXSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), btcTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `btc_tx` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, btcTXPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from btcTX slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for btc_tx")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BTCTX) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBTCTX(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BTCTXSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BTCTXSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), btcTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `btc_tx`.* FROM `btc_tx` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, btcTXPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BTCTXSlice")
	}

	*o = slice

	return nil
}

// BTCTXExists checks if the BTCTX row exists.
func BTCTXExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `btc_tx` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if btc_tx exists")
	}

	return exists, nil
}

// InsertAll inserts all rows with the specified column values, using an executor.
func (o BTCTXSlice) InsertAll(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	var sql string
	vals := []interface{}{}
	for i, row := range o {

		nzDefaults := queries.NonZeroDefaultSet(btcTXColumnsWithDefault, row)
		wl, _ := columns.InsertColumnSet(
			btcTXAllColumns,
			btcTXColumnsWithDefault,
			btcTXColumnsWithoutDefault,
			nzDefaults,
		)
		if i == 0 {
			sql = "INSERT INTO `btc_tx` " + "(`" + strings.Join(wl, "`,`") + "`)" + " VALUES "
		}
		sql += strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), len(vals)+1, len(wl))
		if i != len(o)-1 {
			sql += ","
		}
		valMapping, err := queries.BindMapping(btcTXType, btcTXMapping, wl)
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
		return errors.Wrap(err, "models: unable to insert into btc_tx")
	}

	return nil
}
