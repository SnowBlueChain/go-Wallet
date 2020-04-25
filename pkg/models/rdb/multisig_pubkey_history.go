// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// MultisigPubkeyHistory is an object representing the database table.
type MultisigPubkeyHistory struct {
	ID                    uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Coin                  null.String `boil:"coin" json:"coin,omitempty" toml:"coin" yaml:"coin,omitempty"`
	Account               null.String `boil:"account" json:"account,omitempty" toml:"account" yaml:"account,omitempty"`
	FullPublicKey         string      `boil:"full_public_key" json:"full_public_key" toml:"full_public_key" yaml:"full_public_key"`
	AuthAddress1          string      `boil:"auth_address1" json:"auth_address1" toml:"auth_address1" yaml:"auth_address1"`
	AuthAddress2          string      `boil:"auth_address2" json:"auth_address2" toml:"auth_address2" yaml:"auth_address2"`
	WalletMultisigAddress string      `boil:"wallet_multisig_address" json:"wallet_multisig_address" toml:"wallet_multisig_address" yaml:"wallet_multisig_address"`
	RedeemScript          string      `boil:"redeem_script" json:"redeem_script" toml:"redeem_script" yaml:"redeem_script"`
	IsExported            null.Int8   `boil:"is_exported" json:"is_exported,omitempty" toml:"is_exported" yaml:"is_exported,omitempty"`
	UpdatedAt             null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *multisigPubkeyHistoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L multisigPubkeyHistoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MultisigPubkeyHistoryColumns = struct {
	ID                    string
	Coin                  string
	Account               string
	FullPublicKey         string
	AuthAddress1          string
	AuthAddress2          string
	WalletMultisigAddress string
	RedeemScript          string
	IsExported            string
	UpdatedAt             string
}{
	ID:                    "id",
	Coin:                  "coin",
	Account:               "account",
	FullPublicKey:         "full_public_key",
	AuthAddress1:          "auth_address1",
	AuthAddress2:          "auth_address2",
	WalletMultisigAddress: "wallet_multisig_address",
	RedeemScript:          "redeem_script",
	IsExported:            "is_exported",
	UpdatedAt:             "updated_at",
}

// Generated where

type whereHelpernull_Int8 struct{ field string }

func (w whereHelpernull_Int8) EQ(x null.Int8) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int8) NEQ(x null.Int8) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int8) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int8) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int8) LT(x null.Int8) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int8) LTE(x null.Int8) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int8) GT(x null.Int8) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int8) GTE(x null.Int8) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var MultisigPubkeyHistoryWhere = struct {
	ID                    whereHelperuint64
	Coin                  whereHelpernull_String
	Account               whereHelpernull_String
	FullPublicKey         whereHelperstring
	AuthAddress1          whereHelperstring
	AuthAddress2          whereHelperstring
	WalletMultisigAddress whereHelperstring
	RedeemScript          whereHelperstring
	IsExported            whereHelpernull_Int8
	UpdatedAt             whereHelpernull_Time
}{
	ID:                    whereHelperuint64{field: "`multisig_pubkey_history`.`id`"},
	Coin:                  whereHelpernull_String{field: "`multisig_pubkey_history`.`coin`"},
	Account:               whereHelpernull_String{field: "`multisig_pubkey_history`.`account`"},
	FullPublicKey:         whereHelperstring{field: "`multisig_pubkey_history`.`full_public_key`"},
	AuthAddress1:          whereHelperstring{field: "`multisig_pubkey_history`.`auth_address1`"},
	AuthAddress2:          whereHelperstring{field: "`multisig_pubkey_history`.`auth_address2`"},
	WalletMultisigAddress: whereHelperstring{field: "`multisig_pubkey_history`.`wallet_multisig_address`"},
	RedeemScript:          whereHelperstring{field: "`multisig_pubkey_history`.`redeem_script`"},
	IsExported:            whereHelpernull_Int8{field: "`multisig_pubkey_history`.`is_exported`"},
	UpdatedAt:             whereHelpernull_Time{field: "`multisig_pubkey_history`.`updated_at`"},
}

// MultisigPubkeyHistoryRels is where relationship names are stored.
var MultisigPubkeyHistoryRels = struct {
}{}

// multisigPubkeyHistoryR is where relationships are stored.
type multisigPubkeyHistoryR struct {
}

// NewStruct creates a new relationship struct
func (*multisigPubkeyHistoryR) NewStruct() *multisigPubkeyHistoryR {
	return &multisigPubkeyHistoryR{}
}

// multisigPubkeyHistoryL is where Load methods for each relationship are stored.
type multisigPubkeyHistoryL struct{}

var (
	multisigPubkeyHistoryAllColumns            = []string{"id", "coin", "account", "full_public_key", "auth_address1", "auth_address2", "wallet_multisig_address", "redeem_script", "is_exported", "updated_at"}
	multisigPubkeyHistoryColumnsWithoutDefault = []string{"coin", "account", "full_public_key", "auth_address1", "auth_address2", "wallet_multisig_address", "redeem_script"}
	multisigPubkeyHistoryColumnsWithDefault    = []string{"id", "is_exported", "updated_at"}
	multisigPubkeyHistoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// MultisigPubkeyHistorySlice is an alias for a slice of pointers to MultisigPubkeyHistory.
	// This should generally be used opposed to []MultisigPubkeyHistory.
	MultisigPubkeyHistorySlice []*MultisigPubkeyHistory

	multisigPubkeyHistoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	multisigPubkeyHistoryType                 = reflect.TypeOf(&MultisigPubkeyHistory{})
	multisigPubkeyHistoryMapping              = queries.MakeStructMapping(multisigPubkeyHistoryType)
	multisigPubkeyHistoryPrimaryKeyMapping, _ = queries.BindMapping(multisigPubkeyHistoryType, multisigPubkeyHistoryMapping, multisigPubkeyHistoryPrimaryKeyColumns)
	multisigPubkeyHistoryInsertCacheMut       sync.RWMutex
	multisigPubkeyHistoryInsertCache          = make(map[string]insertCache)
	multisigPubkeyHistoryUpdateCacheMut       sync.RWMutex
	multisigPubkeyHistoryUpdateCache          = make(map[string]updateCache)
	multisigPubkeyHistoryUpsertCacheMut       sync.RWMutex
	multisigPubkeyHistoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single multisigPubkeyHistory record from the query.
func (q multisigPubkeyHistoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MultisigPubkeyHistory, error) {
	o := &MultisigPubkeyHistory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for multisig_pubkey_history")
	}

	return o, nil
}

// All returns all MultisigPubkeyHistory records from the query.
func (q multisigPubkeyHistoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (MultisigPubkeyHistorySlice, error) {
	var o []*MultisigPubkeyHistory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MultisigPubkeyHistory slice")
	}

	return o, nil
}

// Count returns the count of all MultisigPubkeyHistory records in the query.
func (q multisigPubkeyHistoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count multisig_pubkey_history rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q multisigPubkeyHistoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if multisig_pubkey_history exists")
	}

	return count > 0, nil
}

// MultisigPubkeyHistories retrieves all the records using an executor.
func MultisigPubkeyHistories(mods ...qm.QueryMod) multisigPubkeyHistoryQuery {
	mods = append(mods, qm.From("`multisig_pubkey_history`"))
	return multisigPubkeyHistoryQuery{NewQuery(mods...)}
}

// FindMultisigPubkeyHistory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMultisigPubkeyHistory(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*MultisigPubkeyHistory, error) {
	multisigPubkeyHistoryObj := &MultisigPubkeyHistory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `multisig_pubkey_history` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, multisigPubkeyHistoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from multisig_pubkey_history")
	}

	return multisigPubkeyHistoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MultisigPubkeyHistory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no multisig_pubkey_history provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(multisigPubkeyHistoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	multisigPubkeyHistoryInsertCacheMut.RLock()
	cache, cached := multisigPubkeyHistoryInsertCache[key]
	multisigPubkeyHistoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			multisigPubkeyHistoryAllColumns,
			multisigPubkeyHistoryColumnsWithDefault,
			multisigPubkeyHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(multisigPubkeyHistoryType, multisigPubkeyHistoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(multisigPubkeyHistoryType, multisigPubkeyHistoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `multisig_pubkey_history` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `multisig_pubkey_history` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `multisig_pubkey_history` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, multisigPubkeyHistoryPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into multisig_pubkey_history")
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

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == multisigPubkeyHistoryMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for multisig_pubkey_history")
	}

CacheNoHooks:
	if !cached {
		multisigPubkeyHistoryInsertCacheMut.Lock()
		multisigPubkeyHistoryInsertCache[key] = cache
		multisigPubkeyHistoryInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the MultisigPubkeyHistory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MultisigPubkeyHistory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	multisigPubkeyHistoryUpdateCacheMut.RLock()
	cache, cached := multisigPubkeyHistoryUpdateCache[key]
	multisigPubkeyHistoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			multisigPubkeyHistoryAllColumns,
			multisigPubkeyHistoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update multisig_pubkey_history, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `multisig_pubkey_history` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, multisigPubkeyHistoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(multisigPubkeyHistoryType, multisigPubkeyHistoryMapping, append(wl, multisigPubkeyHistoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update multisig_pubkey_history row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for multisig_pubkey_history")
	}

	if !cached {
		multisigPubkeyHistoryUpdateCacheMut.Lock()
		multisigPubkeyHistoryUpdateCache[key] = cache
		multisigPubkeyHistoryUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q multisigPubkeyHistoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for multisig_pubkey_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for multisig_pubkey_history")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MultisigPubkeyHistorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), multisigPubkeyHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `multisig_pubkey_history` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, multisigPubkeyHistoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in multisigPubkeyHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all multisigPubkeyHistory")
	}
	return rowsAff, nil
}

var mySQLMultisigPubkeyHistoryUniqueColumns = []string{
	"id",
	"full_public_key",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MultisigPubkeyHistory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no multisig_pubkey_history provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(multisigPubkeyHistoryColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMultisigPubkeyHistoryUniqueColumns, o)

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

	multisigPubkeyHistoryUpsertCacheMut.RLock()
	cache, cached := multisigPubkeyHistoryUpsertCache[key]
	multisigPubkeyHistoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			multisigPubkeyHistoryAllColumns,
			multisigPubkeyHistoryColumnsWithDefault,
			multisigPubkeyHistoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			multisigPubkeyHistoryAllColumns,
			multisigPubkeyHistoryPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert multisig_pubkey_history, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "multisig_pubkey_history", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `multisig_pubkey_history` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(multisigPubkeyHistoryType, multisigPubkeyHistoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(multisigPubkeyHistoryType, multisigPubkeyHistoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for multisig_pubkey_history")
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

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == multisigPubkeyHistoryMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(multisigPubkeyHistoryType, multisigPubkeyHistoryMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for multisig_pubkey_history")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for multisig_pubkey_history")
	}

CacheNoHooks:
	if !cached {
		multisigPubkeyHistoryUpsertCacheMut.Lock()
		multisigPubkeyHistoryUpsertCache[key] = cache
		multisigPubkeyHistoryUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single MultisigPubkeyHistory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MultisigPubkeyHistory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MultisigPubkeyHistory provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), multisigPubkeyHistoryPrimaryKeyMapping)
	sql := "DELETE FROM `multisig_pubkey_history` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from multisig_pubkey_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for multisig_pubkey_history")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q multisigPubkeyHistoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no multisigPubkeyHistoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from multisig_pubkey_history")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for multisig_pubkey_history")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MultisigPubkeyHistorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), multisigPubkeyHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `multisig_pubkey_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, multisigPubkeyHistoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from multisigPubkeyHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for multisig_pubkey_history")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *MultisigPubkeyHistory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMultisigPubkeyHistory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MultisigPubkeyHistorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MultisigPubkeyHistorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), multisigPubkeyHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `multisig_pubkey_history`.* FROM `multisig_pubkey_history` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, multisigPubkeyHistoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MultisigPubkeyHistorySlice")
	}

	*o = slice

	return nil
}

// MultisigPubkeyHistoryExists checks if the MultisigPubkeyHistory row exists.
func MultisigPubkeyHistoryExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `multisig_pubkey_history` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if multisig_pubkey_history exists")
	}

	return exists, nil
}

// InsertAll inserts all rows with the specified column values, using an executor.
func (o MultisigPubkeyHistorySlice) InsertAll(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}
	var sql string
	vals := []interface{}{}
	for i, row := range o {
		if !boil.TimestampsAreSkipped(ctx) {
			currTime := time.Now().In(boil.GetLocation())

			if queries.MustTime(row.UpdatedAt).IsZero() {
				queries.SetScanner(&row.UpdatedAt, currTime)
			}
		}

		nzDefaults := queries.NonZeroDefaultSet(multisigPubkeyHistoryColumnsWithDefault, row)
		wl, _ := columns.InsertColumnSet(
			multisigPubkeyHistoryAllColumns,
			multisigPubkeyHistoryColumnsWithDefault,
			multisigPubkeyHistoryColumnsWithoutDefault,
			nzDefaults,
		)
		if i == 0 {
			sql = "INSERT INTO `multisig_pubkey_history` " + "(`" + strings.Join(wl, "`,`") + "`)" + " VALUES "
		}
		sql += strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), len(vals)+1, len(wl))
		if i != len(o)-1 {
			sql += ","
		}
		valMapping, err := queries.BindMapping(multisigPubkeyHistoryType, multisigPubkeyHistoryMapping, wl)
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
		return errors.Wrap(err, "models: unable to insert into multisig_pubkey_history")
	}

	return nil
}
