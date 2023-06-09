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
	"github.com/volatiletech/strmangle"
)

// XRPAccountKey is an object representing the database table.
type XRPAccountKey struct { // ID
	ID int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	// coin type code
	Coin string `boil:"coin" json:"coin" toml:"coin" yaml:"coin"`
	// account type
	Account string `boil:"account" json:"account" toml:"account" yaml:"account"`
	// account_id
	AccountID string `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	// key_type
	KeyType int8 `boil:"key_type" json:"key_type" toml:"key_type" yaml:"key_type"`
	// master_key, DEPRECATED
	MasterKey string `boil:"master_key" json:"master_key" toml:"master_key" yaml:"master_key"`
	// master_seed
	MasterSeed string `boil:"master_seed" json:"master_seed" toml:"master_seed" yaml:"master_seed"`
	// master_seed_hex
	MasterSeedHex string `boil:"master_seed_hex" json:"master_seed_hex" toml:"master_seed_hex" yaml:"master_seed_hex"`
	// public_key
	PublicKey string `boil:"public_key" json:"public_key" toml:"public_key" yaml:"public_key"`
	// public_key_hex
	PublicKeyHex string `boil:"public_key_hex" json:"public_key_hex" toml:"public_key_hex" yaml:"public_key_hex"`
	// true: this key is for regular key pair
	IsRegularKeyPair bool `boil:"is_regular_key_pair" json:"is_regular_key_pair" toml:"is_regular_key_pair" yaml:"is_regular_key_pair"`
	// index for hd wallet
	AllocatedID int64 `boil:"allocated_id" json:"allocated_id" toml:"allocated_id" yaml:"allocated_id"`
	// progress status for address generating
	AddrStatus int8 `boil:"addr_status" json:"addr_status" toml:"addr_status" yaml:"addr_status"`
	// updated date
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *xrpAccountKeyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L xrpAccountKeyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var XRPAccountKeyColumns = struct {
	ID               string
	Coin             string
	Account          string
	AccountID        string
	KeyType          string
	MasterKey        string
	MasterSeed       string
	MasterSeedHex    string
	PublicKey        string
	PublicKeyHex     string
	IsRegularKeyPair string
	AllocatedID      string
	AddrStatus       string
	UpdatedAt        string
}{
	ID:               "id",
	Coin:             "coin",
	Account:          "account",
	AccountID:        "account_id",
	KeyType:          "key_type",
	MasterKey:        "master_key",
	MasterSeed:       "master_seed",
	MasterSeedHex:    "master_seed_hex",
	PublicKey:        "public_key",
	PublicKeyHex:     "public_key_hex",
	IsRegularKeyPair: "is_regular_key_pair",
	AllocatedID:      "allocated_id",
	AddrStatus:       "addr_status",
	UpdatedAt:        "updated_at",
}

var XRPAccountKeyTableColumns = struct {
	ID               string
	Coin             string
	Account          string
	AccountID        string
	KeyType          string
	MasterKey        string
	MasterSeed       string
	MasterSeedHex    string
	PublicKey        string
	PublicKeyHex     string
	IsRegularKeyPair string
	AllocatedID      string
	AddrStatus       string
	UpdatedAt        string
}{
	ID:               "xrp_account_key.id",
	Coin:             "xrp_account_key.coin",
	Account:          "xrp_account_key.account",
	AccountID:        "xrp_account_key.account_id",
	KeyType:          "xrp_account_key.key_type",
	MasterKey:        "xrp_account_key.master_key",
	MasterSeed:       "xrp_account_key.master_seed",
	MasterSeedHex:    "xrp_account_key.master_seed_hex",
	PublicKey:        "xrp_account_key.public_key",
	PublicKeyHex:     "xrp_account_key.public_key_hex",
	IsRegularKeyPair: "xrp_account_key.is_regular_key_pair",
	AllocatedID:      "xrp_account_key.allocated_id",
	AddrStatus:       "xrp_account_key.addr_status",
	UpdatedAt:        "xrp_account_key.updated_at",
}

// Generated where

var XRPAccountKeyWhere = struct {
	ID               whereHelperint64
	Coin             whereHelperstring
	Account          whereHelperstring
	AccountID        whereHelperstring
	KeyType          whereHelperint8
	MasterKey        whereHelperstring
	MasterSeed       whereHelperstring
	MasterSeedHex    whereHelperstring
	PublicKey        whereHelperstring
	PublicKeyHex     whereHelperstring
	IsRegularKeyPair whereHelperbool
	AllocatedID      whereHelperint64
	AddrStatus       whereHelperint8
	UpdatedAt        whereHelpernull_Time
}{
	ID:               whereHelperint64{field: "`xrp_account_key`.`id`"},
	Coin:             whereHelperstring{field: "`xrp_account_key`.`coin`"},
	Account:          whereHelperstring{field: "`xrp_account_key`.`account`"},
	AccountID:        whereHelperstring{field: "`xrp_account_key`.`account_id`"},
	KeyType:          whereHelperint8{field: "`xrp_account_key`.`key_type`"},
	MasterKey:        whereHelperstring{field: "`xrp_account_key`.`master_key`"},
	MasterSeed:       whereHelperstring{field: "`xrp_account_key`.`master_seed`"},
	MasterSeedHex:    whereHelperstring{field: "`xrp_account_key`.`master_seed_hex`"},
	PublicKey:        whereHelperstring{field: "`xrp_account_key`.`public_key`"},
	PublicKeyHex:     whereHelperstring{field: "`xrp_account_key`.`public_key_hex`"},
	IsRegularKeyPair: whereHelperbool{field: "`xrp_account_key`.`is_regular_key_pair`"},
	AllocatedID:      whereHelperint64{field: "`xrp_account_key`.`allocated_id`"},
	AddrStatus:       whereHelperint8{field: "`xrp_account_key`.`addr_status`"},
	UpdatedAt:        whereHelpernull_Time{field: "`xrp_account_key`.`updated_at`"},
}

// XRPAccountKeyRels is where relationship names are stored.
var XRPAccountKeyRels = struct {
}{}

// xrpAccountKeyR is where relationships are stored.
type xrpAccountKeyR struct {
}

// NewStruct creates a new relationship struct
func (*xrpAccountKeyR) NewStruct() *xrpAccountKeyR {
	return &xrpAccountKeyR{}
}

// xrpAccountKeyL is where Load methods for each relationship are stored.
type xrpAccountKeyL struct{}

var (
	xrpAccountKeyAllColumns            = []string{"id", "coin", "account", "account_id", "key_type", "master_key", "master_seed", "master_seed_hex", "public_key", "public_key_hex", "is_regular_key_pair", "allocated_id", "addr_status", "updated_at"}
	xrpAccountKeyColumnsWithoutDefault = []string{"coin", "account", "account_id", "master_key", "master_seed", "master_seed_hex", "public_key", "public_key_hex"}
	xrpAccountKeyColumnsWithDefault    = []string{"id", "key_type", "is_regular_key_pair", "allocated_id", "addr_status", "updated_at"}
	xrpAccountKeyPrimaryKeyColumns     = []string{"id"}
	xrpAccountKeyGeneratedColumns      = []string{}
)

type (
	// XRPAccountKeySlice is an alias for a slice of pointers to XRPAccountKey.
	// This should almost always be used instead of []XRPAccountKey.
	XRPAccountKeySlice []*XRPAccountKey

	xrpAccountKeyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	xrpAccountKeyType                 = reflect.TypeOf(&XRPAccountKey{})
	xrpAccountKeyMapping              = queries.MakeStructMapping(xrpAccountKeyType)
	xrpAccountKeyPrimaryKeyMapping, _ = queries.BindMapping(xrpAccountKeyType, xrpAccountKeyMapping, xrpAccountKeyPrimaryKeyColumns)
	xrpAccountKeyInsertCacheMut       sync.RWMutex
	xrpAccountKeyInsertCache          = make(map[string]insertCache)
	xrpAccountKeyUpdateCacheMut       sync.RWMutex
	xrpAccountKeyUpdateCache          = make(map[string]updateCache)
	xrpAccountKeyUpsertCacheMut       sync.RWMutex
	xrpAccountKeyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single xrpAccountKey record from the query.
func (q xrpAccountKeyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*XRPAccountKey, error) {
	o := &XRPAccountKey{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for xrp_account_key")
	}

	return o, nil
}

// All returns all XRPAccountKey records from the query.
func (q xrpAccountKeyQuery) All(ctx context.Context, exec boil.ContextExecutor) (XRPAccountKeySlice, error) {
	var o []*XRPAccountKey

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to XRPAccountKey slice")
	}

	return o, nil
}

// Count returns the count of all XRPAccountKey records in the query.
func (q xrpAccountKeyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count xrp_account_key rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q xrpAccountKeyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if xrp_account_key exists")
	}

	return count > 0, nil
}

// XRPAccountKeys retrieves all the records using an executor.
func XRPAccountKeys(mods ...qm.QueryMod) xrpAccountKeyQuery {
	mods = append(mods, qm.From("`xrp_account_key`"))
	return xrpAccountKeyQuery{NewQuery(mods...)}
}

// FindXRPAccountKey retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindXRPAccountKey(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*XRPAccountKey, error) {
	xrpAccountKeyObj := &XRPAccountKey{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `xrp_account_key` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, xrpAccountKeyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from xrp_account_key")
	}

	return xrpAccountKeyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *XRPAccountKey) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no xrp_account_key provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(xrpAccountKeyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	xrpAccountKeyInsertCacheMut.RLock()
	cache, cached := xrpAccountKeyInsertCache[key]
	xrpAccountKeyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			xrpAccountKeyAllColumns,
			xrpAccountKeyColumnsWithDefault,
			xrpAccountKeyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(xrpAccountKeyType, xrpAccountKeyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(xrpAccountKeyType, xrpAccountKeyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `xrp_account_key` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `xrp_account_key` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `xrp_account_key` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, xrpAccountKeyPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into xrp_account_key")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == xrpAccountKeyMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for xrp_account_key")
	}

CacheNoHooks:
	if !cached {
		xrpAccountKeyInsertCacheMut.Lock()
		xrpAccountKeyInsertCache[key] = cache
		xrpAccountKeyInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the XRPAccountKey.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *XRPAccountKey) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	xrpAccountKeyUpdateCacheMut.RLock()
	cache, cached := xrpAccountKeyUpdateCache[key]
	xrpAccountKeyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			xrpAccountKeyAllColumns,
			xrpAccountKeyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update xrp_account_key, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `xrp_account_key` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, xrpAccountKeyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(xrpAccountKeyType, xrpAccountKeyMapping, append(wl, xrpAccountKeyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update xrp_account_key row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for xrp_account_key")
	}

	if !cached {
		xrpAccountKeyUpdateCacheMut.Lock()
		xrpAccountKeyUpdateCache[key] = cache
		xrpAccountKeyUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q xrpAccountKeyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for xrp_account_key")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for xrp_account_key")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o XRPAccountKeySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), xrpAccountKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `xrp_account_key` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, xrpAccountKeyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in xrpAccountKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all xrpAccountKey")
	}
	return rowsAff, nil
}

var mySQLXRPAccountKeyUniqueColumns = []string{
	"id",
	"account_id",
	"master_seed",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *XRPAccountKey) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no xrp_account_key provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(xrpAccountKeyColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLXRPAccountKeyUniqueColumns, o)

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

	xrpAccountKeyUpsertCacheMut.RLock()
	cache, cached := xrpAccountKeyUpsertCache[key]
	xrpAccountKeyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			xrpAccountKeyAllColumns,
			xrpAccountKeyColumnsWithDefault,
			xrpAccountKeyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			xrpAccountKeyAllColumns,
			xrpAccountKeyPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert xrp_account_key, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`xrp_account_key`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `xrp_account_key` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(xrpAccountKeyType, xrpAccountKeyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(xrpAccountKeyType, xrpAccountKeyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for xrp_account_key")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == xrpAccountKeyMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(xrpAccountKeyType, xrpAccountKeyMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for xrp_account_key")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for xrp_account_key")
	}

CacheNoHooks:
	if !cached {
		xrpAccountKeyUpsertCacheMut.Lock()
		xrpAccountKeyUpsertCache[key] = cache
		xrpAccountKeyUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single XRPAccountKey record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *XRPAccountKey) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no XRPAccountKey provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), xrpAccountKeyPrimaryKeyMapping)
	sql := "DELETE FROM `xrp_account_key` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from xrp_account_key")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for xrp_account_key")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q xrpAccountKeyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no xrpAccountKeyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from xrp_account_key")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for xrp_account_key")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o XRPAccountKeySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), xrpAccountKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `xrp_account_key` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, xrpAccountKeyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from xrpAccountKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for xrp_account_key")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *XRPAccountKey) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindXRPAccountKey(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *XRPAccountKeySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := XRPAccountKeySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), xrpAccountKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `xrp_account_key`.* FROM `xrp_account_key` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, xrpAccountKeyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in XRPAccountKeySlice")
	}

	*o = slice

	return nil
}

// XRPAccountKeyExists checks if the XRPAccountKey row exists.
func XRPAccountKeyExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `xrp_account_key` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if xrp_account_key exists")
	}

	return exists, nil
}
