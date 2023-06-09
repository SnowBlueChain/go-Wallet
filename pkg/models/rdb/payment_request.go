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

// PaymentRequest is an object representing the database table.
type PaymentRequest struct { // ID
	ID int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	// coin type code
	Coin string `boil:"coin" json:"coin" toml:"coin" yaml:"coin"`
	// tx table ID for payment action
	PaymentID null.Int64 `boil:"payment_id" json:"payment_id,omitempty" toml:"payment_id" yaml:"payment_id,omitempty"`
	// sender address
	SenderAddress string `boil:"sender_address" json:"sender_address" toml:"sender_address" yaml:"sender_address"`
	// sender account
	SenderAccount string `boil:"sender_account" json:"sender_account" toml:"sender_account" yaml:"sender_account"`
	// receiver address
	ReceiverAddress string `boil:"receiver_address" json:"receiver_address" toml:"receiver_address" yaml:"receiver_address"`
	// amount of coin to send
	Amount types.Decimal `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	// true: unsigned transaction is created
	IsDone bool `boil:"is_done" json:"is_done" toml:"is_done" yaml:"is_done"`
	// updated date
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *paymentRequestR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L paymentRequestL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PaymentRequestColumns = struct {
	ID              string
	Coin            string
	PaymentID       string
	SenderAddress   string
	SenderAccount   string
	ReceiverAddress string
	Amount          string
	IsDone          string
	UpdatedAt       string
}{
	ID:              "id",
	Coin:            "coin",
	PaymentID:       "payment_id",
	SenderAddress:   "sender_address",
	SenderAccount:   "sender_account",
	ReceiverAddress: "receiver_address",
	Amount:          "amount",
	IsDone:          "is_done",
	UpdatedAt:       "updated_at",
}

var PaymentRequestTableColumns = struct {
	ID              string
	Coin            string
	PaymentID       string
	SenderAddress   string
	SenderAccount   string
	ReceiverAddress string
	Amount          string
	IsDone          string
	UpdatedAt       string
}{
	ID:              "payment_request.id",
	Coin:            "payment_request.coin",
	PaymentID:       "payment_request.payment_id",
	SenderAddress:   "payment_request.sender_address",
	SenderAccount:   "payment_request.sender_account",
	ReceiverAddress: "payment_request.receiver_address",
	Amount:          "payment_request.amount",
	IsDone:          "payment_request.is_done",
	UpdatedAt:       "payment_request.updated_at",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var PaymentRequestWhere = struct {
	ID              whereHelperint64
	Coin            whereHelperstring
	PaymentID       whereHelpernull_Int64
	SenderAddress   whereHelperstring
	SenderAccount   whereHelperstring
	ReceiverAddress whereHelperstring
	Amount          whereHelpertypes_Decimal
	IsDone          whereHelperbool
	UpdatedAt       whereHelpernull_Time
}{
	ID:              whereHelperint64{field: "`payment_request`.`id`"},
	Coin:            whereHelperstring{field: "`payment_request`.`coin`"},
	PaymentID:       whereHelpernull_Int64{field: "`payment_request`.`payment_id`"},
	SenderAddress:   whereHelperstring{field: "`payment_request`.`sender_address`"},
	SenderAccount:   whereHelperstring{field: "`payment_request`.`sender_account`"},
	ReceiverAddress: whereHelperstring{field: "`payment_request`.`receiver_address`"},
	Amount:          whereHelpertypes_Decimal{field: "`payment_request`.`amount`"},
	IsDone:          whereHelperbool{field: "`payment_request`.`is_done`"},
	UpdatedAt:       whereHelpernull_Time{field: "`payment_request`.`updated_at`"},
}

// PaymentRequestRels is where relationship names are stored.
var PaymentRequestRels = struct {
}{}

// paymentRequestR is where relationships are stored.
type paymentRequestR struct {
}

// NewStruct creates a new relationship struct
func (*paymentRequestR) NewStruct() *paymentRequestR {
	return &paymentRequestR{}
}

// paymentRequestL is where Load methods for each relationship are stored.
type paymentRequestL struct{}

var (
	paymentRequestAllColumns            = []string{"id", "coin", "payment_id", "sender_address", "sender_account", "receiver_address", "amount", "is_done", "updated_at"}
	paymentRequestColumnsWithoutDefault = []string{"coin", "payment_id", "sender_address", "sender_account", "receiver_address", "amount"}
	paymentRequestColumnsWithDefault    = []string{"id", "is_done", "updated_at"}
	paymentRequestPrimaryKeyColumns     = []string{"id"}
	paymentRequestGeneratedColumns      = []string{}
)

type (
	// PaymentRequestSlice is an alias for a slice of pointers to PaymentRequest.
	// This should almost always be used instead of []PaymentRequest.
	PaymentRequestSlice []*PaymentRequest

	paymentRequestQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	paymentRequestType                 = reflect.TypeOf(&PaymentRequest{})
	paymentRequestMapping              = queries.MakeStructMapping(paymentRequestType)
	paymentRequestPrimaryKeyMapping, _ = queries.BindMapping(paymentRequestType, paymentRequestMapping, paymentRequestPrimaryKeyColumns)
	paymentRequestInsertCacheMut       sync.RWMutex
	paymentRequestInsertCache          = make(map[string]insertCache)
	paymentRequestUpdateCacheMut       sync.RWMutex
	paymentRequestUpdateCache          = make(map[string]updateCache)
	paymentRequestUpsertCacheMut       sync.RWMutex
	paymentRequestUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single paymentRequest record from the query.
func (q paymentRequestQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PaymentRequest, error) {
	o := &PaymentRequest{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for payment_request")
	}

	return o, nil
}

// All returns all PaymentRequest records from the query.
func (q paymentRequestQuery) All(ctx context.Context, exec boil.ContextExecutor) (PaymentRequestSlice, error) {
	var o []*PaymentRequest

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PaymentRequest slice")
	}

	return o, nil
}

// Count returns the count of all PaymentRequest records in the query.
func (q paymentRequestQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count payment_request rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q paymentRequestQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if payment_request exists")
	}

	return count > 0, nil
}

// PaymentRequests retrieves all the records using an executor.
func PaymentRequests(mods ...qm.QueryMod) paymentRequestQuery {
	mods = append(mods, qm.From("`payment_request`"))
	return paymentRequestQuery{NewQuery(mods...)}
}

// FindPaymentRequest retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPaymentRequest(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*PaymentRequest, error) {
	paymentRequestObj := &PaymentRequest{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `payment_request` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, paymentRequestObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from payment_request")
	}

	return paymentRequestObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PaymentRequest) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no payment_request provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(paymentRequestColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	paymentRequestInsertCacheMut.RLock()
	cache, cached := paymentRequestInsertCache[key]
	paymentRequestInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			paymentRequestAllColumns,
			paymentRequestColumnsWithDefault,
			paymentRequestColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(paymentRequestType, paymentRequestMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(paymentRequestType, paymentRequestMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `payment_request` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `payment_request` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `payment_request` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, paymentRequestPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into payment_request")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == paymentRequestMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for payment_request")
	}

CacheNoHooks:
	if !cached {
		paymentRequestInsertCacheMut.Lock()
		paymentRequestInsertCache[key] = cache
		paymentRequestInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PaymentRequest.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PaymentRequest) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	paymentRequestUpdateCacheMut.RLock()
	cache, cached := paymentRequestUpdateCache[key]
	paymentRequestUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			paymentRequestAllColumns,
			paymentRequestPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update payment_request, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `payment_request` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, paymentRequestPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(paymentRequestType, paymentRequestMapping, append(wl, paymentRequestPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update payment_request row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for payment_request")
	}

	if !cached {
		paymentRequestUpdateCacheMut.Lock()
		paymentRequestUpdateCache[key] = cache
		paymentRequestUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q paymentRequestQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for payment_request")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for payment_request")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PaymentRequestSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), paymentRequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `payment_request` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, paymentRequestPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in paymentRequest slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all paymentRequest")
	}
	return rowsAff, nil
}

var mySQLPaymentRequestUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PaymentRequest) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no payment_request provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(paymentRequestColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPaymentRequestUniqueColumns, o)

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

	paymentRequestUpsertCacheMut.RLock()
	cache, cached := paymentRequestUpsertCache[key]
	paymentRequestUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			paymentRequestAllColumns,
			paymentRequestColumnsWithDefault,
			paymentRequestColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			paymentRequestAllColumns,
			paymentRequestPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert payment_request, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`payment_request`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `payment_request` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(paymentRequestType, paymentRequestMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(paymentRequestType, paymentRequestMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for payment_request")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == paymentRequestMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(paymentRequestType, paymentRequestMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for payment_request")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for payment_request")
	}

CacheNoHooks:
	if !cached {
		paymentRequestUpsertCacheMut.Lock()
		paymentRequestUpsertCache[key] = cache
		paymentRequestUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single PaymentRequest record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PaymentRequest) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PaymentRequest provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), paymentRequestPrimaryKeyMapping)
	sql := "DELETE FROM `payment_request` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from payment_request")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for payment_request")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q paymentRequestQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no paymentRequestQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from payment_request")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for payment_request")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PaymentRequestSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), paymentRequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `payment_request` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, paymentRequestPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from paymentRequest slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for payment_request")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PaymentRequest) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPaymentRequest(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PaymentRequestSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PaymentRequestSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), paymentRequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `payment_request`.* FROM `payment_request` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, paymentRequestPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PaymentRequestSlice")
	}

	*o = slice

	return nil
}

// PaymentRequestExists checks if the PaymentRequest row exists.
func PaymentRequestExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `payment_request` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if payment_request exists")
	}

	return exists, nil
}
