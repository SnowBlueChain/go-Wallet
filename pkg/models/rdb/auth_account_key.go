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
)

// AuthAccountKey is an object representing the database table.
type AuthAccountKey struct {
	ID                 int16     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Coin               string    `boil:"coin" json:"coin" toml:"coin" yaml:"coin"`
	AuthAccount        string    `boil:"auth_account" json:"auth_account" toml:"auth_account" yaml:"auth_account"`
	P2PKHAddress       string    `boil:"p2pkh_address" json:"p2pkh_address" toml:"p2pkh_address" yaml:"p2pkh_address"`
	P2SHSegwitAddress  string    `boil:"p2sh_segwit_address" json:"p2sh_segwit_address" toml:"p2sh_segwit_address" yaml:"p2sh_segwit_address"`
	Bech32Address      string    `boil:"bech32_address" json:"bech32_address" toml:"bech32_address" yaml:"bech32_address"`
	FullPublicKey      string    `boil:"full_public_key" json:"full_public_key" toml:"full_public_key" yaml:"full_public_key"`
	MultisigAddress    string    `boil:"multisig_address" json:"multisig_address" toml:"multisig_address" yaml:"multisig_address"`
	RedeemScript       string    `boil:"redeem_script" json:"redeem_script" toml:"redeem_script" yaml:"redeem_script"`
	WalletImportFormat string    `boil:"wallet_import_format" json:"wallet_import_format" toml:"wallet_import_format" yaml:"wallet_import_format"`
	Idx                int64     `boil:"idx" json:"idx" toml:"idx" yaml:"idx"`
	AddrStatus         int8      `boil:"addr_status" json:"addr_status" toml:"addr_status" yaml:"addr_status"`
	UpdatedAt          null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *authAccountKeyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authAccountKeyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AuthAccountKeyColumns = struct {
	ID                 string
	Coin               string
	AuthAccount        string
	P2PKHAddress       string
	P2SHSegwitAddress  string
	Bech32Address      string
	FullPublicKey      string
	MultisigAddress    string
	RedeemScript       string
	WalletImportFormat string
	Idx                string
	AddrStatus         string
	UpdatedAt          string
}{
	ID:                 "id",
	Coin:               "coin",
	AuthAccount:        "auth_account",
	P2PKHAddress:       "p2pkh_address",
	P2SHSegwitAddress:  "p2sh_segwit_address",
	Bech32Address:      "bech32_address",
	FullPublicKey:      "full_public_key",
	MultisigAddress:    "multisig_address",
	RedeemScript:       "redeem_script",
	WalletImportFormat: "wallet_import_format",
	Idx:                "idx",
	AddrStatus:         "addr_status",
	UpdatedAt:          "updated_at",
}

// Generated where

type whereHelperint16 struct{ field string }

func (w whereHelperint16) EQ(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint16) NEQ(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint16) LT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint16) LTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint16) GT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint16) GTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var AuthAccountKeyWhere = struct {
	ID                 whereHelperint16
	Coin               whereHelperstring
	AuthAccount        whereHelperstring
	P2PKHAddress       whereHelperstring
	P2SHSegwitAddress  whereHelperstring
	Bech32Address      whereHelperstring
	FullPublicKey      whereHelperstring
	MultisigAddress    whereHelperstring
	RedeemScript       whereHelperstring
	WalletImportFormat whereHelperstring
	Idx                whereHelperint64
	AddrStatus         whereHelperint8
	UpdatedAt          whereHelpernull_Time
}{
	ID:                 whereHelperint16{field: "`auth_account_key`.`id`"},
	Coin:               whereHelperstring{field: "`auth_account_key`.`coin`"},
	AuthAccount:        whereHelperstring{field: "`auth_account_key`.`auth_account`"},
	P2PKHAddress:       whereHelperstring{field: "`auth_account_key`.`p2pkh_address`"},
	P2SHSegwitAddress:  whereHelperstring{field: "`auth_account_key`.`p2sh_segwit_address`"},
	Bech32Address:      whereHelperstring{field: "`auth_account_key`.`bech32_address`"},
	FullPublicKey:      whereHelperstring{field: "`auth_account_key`.`full_public_key`"},
	MultisigAddress:    whereHelperstring{field: "`auth_account_key`.`multisig_address`"},
	RedeemScript:       whereHelperstring{field: "`auth_account_key`.`redeem_script`"},
	WalletImportFormat: whereHelperstring{field: "`auth_account_key`.`wallet_import_format`"},
	Idx:                whereHelperint64{field: "`auth_account_key`.`idx`"},
	AddrStatus:         whereHelperint8{field: "`auth_account_key`.`addr_status`"},
	UpdatedAt:          whereHelpernull_Time{field: "`auth_account_key`.`updated_at`"},
}

// AuthAccountKeyRels is where relationship names are stored.
var AuthAccountKeyRels = struct {
}{}

// authAccountKeyR is where relationships are stored.
type authAccountKeyR struct {
}

// NewStruct creates a new relationship struct
func (*authAccountKeyR) NewStruct() *authAccountKeyR {
	return &authAccountKeyR{}
}

// authAccountKeyL is where Load methods for each relationship are stored.
type authAccountKeyL struct{}

var (
	authAccountKeyAllColumns            = []string{"id", "coin", "auth_account", "p2pkh_address", "p2sh_segwit_address", "bech32_address", "full_public_key", "multisig_address", "redeem_script", "wallet_import_format", "idx", "addr_status", "updated_at"}
	authAccountKeyColumnsWithoutDefault = []string{"coin", "auth_account", "p2pkh_address", "p2sh_segwit_address", "bech32_address", "full_public_key", "multisig_address", "redeem_script", "wallet_import_format", "idx"}
	authAccountKeyColumnsWithDefault    = []string{"id", "addr_status", "updated_at"}
	authAccountKeyPrimaryKeyColumns     = []string{"id"}
)

type (
	// AuthAccountKeySlice is an alias for a slice of pointers to AuthAccountKey.
	// This should generally be used opposed to []AuthAccountKey.
	AuthAccountKeySlice []*AuthAccountKey

	authAccountKeyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authAccountKeyType                 = reflect.TypeOf(&AuthAccountKey{})
	authAccountKeyMapping              = queries.MakeStructMapping(authAccountKeyType)
	authAccountKeyPrimaryKeyMapping, _ = queries.BindMapping(authAccountKeyType, authAccountKeyMapping, authAccountKeyPrimaryKeyColumns)
	authAccountKeyInsertCacheMut       sync.RWMutex
	authAccountKeyInsertCache          = make(map[string]insertCache)
	authAccountKeyUpdateCacheMut       sync.RWMutex
	authAccountKeyUpdateCache          = make(map[string]updateCache)
	authAccountKeyUpsertCacheMut       sync.RWMutex
	authAccountKeyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single authAccountKey record from the query.
func (q authAccountKeyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AuthAccountKey, error) {
	o := &AuthAccountKey{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auth_account_key")
	}

	return o, nil
}

// All returns all AuthAccountKey records from the query.
func (q authAccountKeyQuery) All(ctx context.Context, exec boil.ContextExecutor) (AuthAccountKeySlice, error) {
	var o []*AuthAccountKey

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuthAccountKey slice")
	}

	return o, nil
}

// Count returns the count of all AuthAccountKey records in the query.
func (q authAccountKeyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auth_account_key rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q authAccountKeyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auth_account_key exists")
	}

	return count > 0, nil
}

// AuthAccountKeys retrieves all the records using an executor.
func AuthAccountKeys(mods ...qm.QueryMod) authAccountKeyQuery {
	mods = append(mods, qm.From("`auth_account_key`"))
	return authAccountKeyQuery{NewQuery(mods...)}
}

// FindAuthAccountKey retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthAccountKey(ctx context.Context, exec boil.ContextExecutor, iD int16, selectCols ...string) (*AuthAccountKey, error) {
	authAccountKeyObj := &AuthAccountKey{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `auth_account_key` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, authAccountKeyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auth_account_key")
	}

	return authAccountKeyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AuthAccountKey) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no auth_account_key provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(authAccountKeyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	authAccountKeyInsertCacheMut.RLock()
	cache, cached := authAccountKeyInsertCache[key]
	authAccountKeyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			authAccountKeyAllColumns,
			authAccountKeyColumnsWithDefault,
			authAccountKeyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(authAccountKeyType, authAccountKeyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authAccountKeyType, authAccountKeyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `auth_account_key` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `auth_account_key` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `auth_account_key` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, authAccountKeyPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into auth_account_key")
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

	o.ID = int16(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == authAccountKeyMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for auth_account_key")
	}

CacheNoHooks:
	if !cached {
		authAccountKeyInsertCacheMut.Lock()
		authAccountKeyInsertCache[key] = cache
		authAccountKeyInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the AuthAccountKey.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AuthAccountKey) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	authAccountKeyUpdateCacheMut.RLock()
	cache, cached := authAccountKeyUpdateCache[key]
	authAccountKeyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			authAccountKeyAllColumns,
			authAccountKeyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update auth_account_key, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `auth_account_key` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, authAccountKeyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authAccountKeyType, authAccountKeyMapping, append(wl, authAccountKeyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update auth_account_key row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for auth_account_key")
	}

	if !cached {
		authAccountKeyUpdateCacheMut.Lock()
		authAccountKeyUpdateCache[key] = cache
		authAccountKeyUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q authAccountKeyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for auth_account_key")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for auth_account_key")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthAccountKeySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authAccountKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `auth_account_key` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, authAccountKeyPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in authAccountKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all authAccountKey")
	}
	return rowsAff, nil
}

var mySQLAuthAccountKeyUniqueColumns = []string{
	"id",
	"p2pkh_address",
	"p2sh_segwit_address",
	"bech32_address",
	"wallet_import_format",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AuthAccountKey) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no auth_account_key provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(authAccountKeyColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLAuthAccountKeyUniqueColumns, o)

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

	authAccountKeyUpsertCacheMut.RLock()
	cache, cached := authAccountKeyUpsertCache[key]
	authAccountKeyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			authAccountKeyAllColumns,
			authAccountKeyColumnsWithDefault,
			authAccountKeyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			authAccountKeyAllColumns,
			authAccountKeyPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert auth_account_key, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "auth_account_key", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `auth_account_key` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(authAccountKeyType, authAccountKeyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authAccountKeyType, authAccountKeyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for auth_account_key")
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

	o.ID = int16(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == authAccountKeyMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(authAccountKeyType, authAccountKeyMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for auth_account_key")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for auth_account_key")
	}

CacheNoHooks:
	if !cached {
		authAccountKeyUpsertCacheMut.Lock()
		authAccountKeyUpsertCache[key] = cache
		authAccountKeyUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single AuthAccountKey record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthAccountKey) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AuthAccountKey provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authAccountKeyPrimaryKeyMapping)
	sql := "DELETE FROM `auth_account_key` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from auth_account_key")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for auth_account_key")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q authAccountKeyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no authAccountKeyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from auth_account_key")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auth_account_key")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthAccountKeySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authAccountKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `auth_account_key` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, authAccountKeyPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from authAccountKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auth_account_key")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthAccountKey) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAuthAccountKey(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthAccountKeySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AuthAccountKeySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authAccountKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `auth_account_key`.* FROM `auth_account_key` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, authAccountKeyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthAccountKeySlice")
	}

	*o = slice

	return nil
}

// AuthAccountKeyExists checks if the AuthAccountKey row exists.
func AuthAccountKeyExists(ctx context.Context, exec boil.ContextExecutor, iD int16) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `auth_account_key` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auth_account_key exists")
	}

	return exists, nil
}

// InsertAll inserts all rows with the specified column values, using an executor.
func (o AuthAccountKeySlice) InsertAll(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
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

		nzDefaults := queries.NonZeroDefaultSet(authAccountKeyColumnsWithDefault, row)
		wl, _ := columns.InsertColumnSet(
			authAccountKeyAllColumns,
			authAccountKeyColumnsWithDefault,
			authAccountKeyColumnsWithoutDefault,
			nzDefaults,
		)
		if i == 0 {
			sql = "INSERT INTO `auth_account_key` " + "(`" + strings.Join(wl, "`,`") + "`)" + " VALUES "
		}
		sql += strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), len(vals)+1, len(wl))
		if i != len(o)-1 {
			sql += ","
		}
		valMapping, err := queries.BindMapping(authAccountKeyType, authAccountKeyMapping, wl)
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
		return errors.Wrap(err, "models: unable to insert into auth_account_key")
	}

	return nil
}
