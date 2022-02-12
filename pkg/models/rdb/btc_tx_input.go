// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// BTCTXInput is an object representing the database table.
type BTCTXInput struct { // ID
	ID int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	// tx table ID
	TXID int64 `boil:"tx_id" json:"tx_id" toml:"tx_id" yaml:"tx_id"`
	// txid for input
	InputTxid string `boil:"input_txid" json:"input_txid" toml:"input_txid" yaml:"input_txid"`
	// vout for input
	InputVout uint32 `boil:"input_vout" json:"input_vout" toml:"input_vout" yaml:"input_vout"`
	// sender address for input
	InputAddress string `boil:"input_address" json:"input_address" toml:"input_address" yaml:"input_address"`
	// sender account for input
	InputAccount string `boil:"input_account" json:"input_account" toml:"input_account" yaml:"input_account"`
	// amount of coin to send for input
	InputAmount types.Decimal `boil:"input_amount" json:"input_amount" toml:"input_amount" yaml:"input_amount"`
	// block confirmations when unspent rpc returned
	InputConfirmations uint64 `boil:"input_confirmations" json:"input_confirmations" toml:"input_confirmations" yaml:"input_confirmations"`
	// updated date
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *btcTXInputR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L btcTXInputL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BTCTXInputColumns = struct {
	ID                 string
	TXID               string
	InputTxid          string
	InputVout          string
	InputAddress       string
	InputAccount       string
	InputAmount        string
	InputConfirmations string
	UpdatedAt          string
}{
	ID:                 "id",
	TXID:               "tx_id",
	InputTxid:          "input_txid",
	InputVout:          "input_vout",
	InputAddress:       "input_address",
	InputAccount:       "input_account",
	InputAmount:        "input_amount",
	InputConfirmations: "input_confirmations",
	UpdatedAt:          "updated_at",
}

var BTCTXInputTableColumns = struct {
	ID                 string
	TXID               string
	InputTxid          string
	InputVout          string
	InputAddress       string
	InputAccount       string
	InputAmount        string
	InputConfirmations string
	UpdatedAt          string
}{
	ID:                 "btc_tx_input.id",
	TXID:               "btc_tx_input.tx_id",
	InputTxid:          "btc_tx_input.input_txid",
	InputVout:          "btc_tx_input.input_vout",
	InputAddress:       "btc_tx_input.input_address",
	InputAccount:       "btc_tx_input.input_account",
	InputAmount:        "btc_tx_input.input_amount",
	InputConfirmations: "btc_tx_input.input_confirmations",
	UpdatedAt:          "btc_tx_input.updated_at",
}

// Generated where

type whereHelperuint32 struct{ field string }

func (w whereHelperuint32) EQ(x uint32) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint32) NEQ(x uint32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint32) LT(x uint32) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint32) LTE(x uint32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint32) GT(x uint32) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint32) GTE(x uint32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint32) IN(slice []uint32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint32) NIN(slice []uint32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperuint64 struct{ field string }

func (w whereHelperuint64) EQ(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint64) NEQ(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint64) LT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint64) LTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint64) GT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint64) GTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint64) IN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint64) NIN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var BTCTXInputWhere = struct {
	ID                 whereHelperint64
	TXID               whereHelperint64
	InputTxid          whereHelperstring
	InputVout          whereHelperuint32
	InputAddress       whereHelperstring
	InputAccount       whereHelperstring
	InputAmount        whereHelpertypes_Decimal
	InputConfirmations whereHelperuint64
	UpdatedAt          whereHelpernull_Time
}{
	ID:                 whereHelperint64{field: "`btc_tx_input`.`id`"},
	TXID:               whereHelperint64{field: "`btc_tx_input`.`tx_id`"},
	InputTxid:          whereHelperstring{field: "`btc_tx_input`.`input_txid`"},
	InputVout:          whereHelperuint32{field: "`btc_tx_input`.`input_vout`"},
	InputAddress:       whereHelperstring{field: "`btc_tx_input`.`input_address`"},
	InputAccount:       whereHelperstring{field: "`btc_tx_input`.`input_account`"},
	InputAmount:        whereHelpertypes_Decimal{field: "`btc_tx_input`.`input_amount`"},
	InputConfirmations: whereHelperuint64{field: "`btc_tx_input`.`input_confirmations`"},
	UpdatedAt:          whereHelpernull_Time{field: "`btc_tx_input`.`updated_at`"},
}

// BTCTXInputRels is where relationship names are stored.
var BTCTXInputRels = struct {
}{}

// btcTXInputR is where relationships are stored.
type btcTXInputR struct {
}

// NewStruct creates a new relationship struct
func (*btcTXInputR) NewStruct() *btcTXInputR {
	return &btcTXInputR{}
}

// btcTXInputL is where Load methods for each relationship are stored.
type btcTXInputL struct{}

var (
	btcTXInputAllColumns            = []string{"id", "tx_id", "input_txid", "input_vout", "input_address", "input_account", "input_amount", "input_confirmations", "updated_at"}
	btcTXInputColumnsWithoutDefault = []string{"tx_id", "input_txid", "input_vout", "input_address", "input_account", "input_amount", "input_confirmations"}
	btcTXInputColumnsWithDefault    = []string{"id", "updated_at"}
	btcTXInputPrimaryKeyColumns     = []string{"id"}
	btcTXInputGeneratedColumns      = []string{}
)

type (
	// BTCTXInputSlice is an alias for a slice of pointers to BTCTXInput.
	// This should almost always be used instead of []BTCTXInput.
	BTCTXInputSlice []*BTCTXInput

	btcTXInputQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	btcTXInputType                 = reflect.TypeOf(&BTCTXInput{})
	btcTXInputMapping              = queries.MakeStructMapping(btcTXInputType)
	btcTXInputPrimaryKeyMapping, _ = queries.BindMapping(btcTXInputType, btcTXInputMapping, btcTXInputPrimaryKeyColumns)
	btcTXInputInsertCacheMut       sync.RWMutex
	btcTXInputInsertCache          = make(map[string]insertCache)
	btcTXInputUpdateCacheMut       sync.RWMutex
	btcTXInputUpdateCache          = make(map[string]updateCache)
	btcTXInputUpsertCacheMut       sync.RWMutex
	btcTXInputUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single btcTXInput record from the query.
func (q btcTXInputQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BTCTXInput, error) {
	o := &BTCTXInput{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for btc_tx_input")
	}

	return o, nil
}

// All returns all BTCTXInput records from the query.
func (q btcTXInputQuery) All(ctx context.Context, exec boil.ContextExecutor) (BTCTXInputSlice, error) {
	var o []*BTCTXInput

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BTCTXInput slice")
	}

	return o, nil
}

// Count returns the count of all BTCTXInput records in the query.
func (q btcTXInputQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count btc_tx_input rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q btcTXInputQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if btc_tx_input exists")
	}

	return count > 0, nil
}

// BTCTXInputs retrieves all the records using an executor.
func BTCTXInputs(mods ...qm.QueryMod) btcTXInputQuery {
	mods = append(mods, qm.From("`btc_tx_input`"))
	return btcTXInputQuery{NewQuery(mods...)}
}

// FindBTCTXInput retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBTCTXInput(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*BTCTXInput, error) {
	btcTXInputObj := &BTCTXInput{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `btc_tx_input` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, btcTXInputObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from btc_tx_input")
	}

	return btcTXInputObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BTCTXInput) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no btc_tx_input provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(btcTXInputColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	btcTXInputInsertCacheMut.RLock()
	cache, cached := btcTXInputInsertCache[key]
	btcTXInputInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			btcTXInputAllColumns,
			btcTXInputColumnsWithDefault,
			btcTXInputColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(btcTXInputType, btcTXInputMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(btcTXInputType, btcTXInputMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `btc_tx_input` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `btc_tx_input` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `btc_tx_input` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, btcTXInputPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into btc_tx_input")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == btcTXInputMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for btc_tx_input")
	}

CacheNoHooks:
	if !cached {
		btcTXInputInsertCacheMut.Lock()
		btcTXInputInsertCache[key] = cache
		btcTXInputInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the BTCTXInput.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BTCTXInput) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	btcTXInputUpdateCacheMut.RLock()
	cache, cached := btcTXInputUpdateCache[key]
	btcTXInputUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			btcTXInputAllColumns,
			btcTXInputPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update btc_tx_input, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `btc_tx_input` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, btcTXInputPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(btcTXInputType, btcTXInputMapping, append(wl, btcTXInputPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update btc_tx_input row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for btc_tx_input")
	}

	if !cached {
		btcTXInputUpdateCacheMut.Lock()
		btcTXInputUpdateCache[key] = cache
		btcTXInputUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q btcTXInputQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for btc_tx_input")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for btc_tx_input")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BTCTXInputSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), btcTXInputPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `btc_tx_input` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, btcTXInputPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in btcTXInput slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all btcTXInput")
	}
	return rowsAff, nil
}

var mySQLBTCTXInputUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BTCTXInput) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no btc_tx_input provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(btcTXInputColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBTCTXInputUniqueColumns, o)

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

	btcTXInputUpsertCacheMut.RLock()
	cache, cached := btcTXInputUpsertCache[key]
	btcTXInputUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			btcTXInputAllColumns,
			btcTXInputColumnsWithDefault,
			btcTXInputColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			btcTXInputAllColumns,
			btcTXInputPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert btc_tx_input, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`btc_tx_input`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `btc_tx_input` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(btcTXInputType, btcTXInputMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(btcTXInputType, btcTXInputMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for btc_tx_input")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == btcTXInputMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(btcTXInputType, btcTXInputMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for btc_tx_input")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for btc_tx_input")
	}

CacheNoHooks:
	if !cached {
		btcTXInputUpsertCacheMut.Lock()
		btcTXInputUpsertCache[key] = cache
		btcTXInputUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single BTCTXInput record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BTCTXInput) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BTCTXInput provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), btcTXInputPrimaryKeyMapping)
	sql := "DELETE FROM `btc_tx_input` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from btc_tx_input")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for btc_tx_input")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q btcTXInputQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no btcTXInputQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from btc_tx_input")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for btc_tx_input")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BTCTXInputSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), btcTXInputPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `btc_tx_input` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, btcTXInputPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from btcTXInput slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for btc_tx_input")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BTCTXInput) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBTCTXInput(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BTCTXInputSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BTCTXInputSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), btcTXInputPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `btc_tx_input`.* FROM `btc_tx_input` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, btcTXInputPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BTCTXInputSlice")
	}

	*o = slice

	return nil
}

// BTCTXInputExists checks if the BTCTXInput row exists.
func BTCTXInputExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `btc_tx_input` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if btc_tx_input exists")
	}

	return exists, nil
}
