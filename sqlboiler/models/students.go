// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Student is an object representing the database table.
type Student struct {
	ID      int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name    string `boil:"name" json:"name" toml:"name" yaml:"name"`
	ClassID int    `boil:"class_id" json:"class_id" toml:"class_id" yaml:"class_id"`

	R *studentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L studentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StudentColumns = struct {
	ID      string
	Name    string
	ClassID string
}{
	ID:      "id",
	Name:    "name",
	ClassID: "class_id",
}

var StudentTableColumns = struct {
	ID      string
	Name    string
	ClassID string
}{
	ID:      "students.id",
	Name:    "students.name",
	ClassID: "students.class_id",
}

// Generated where

var StudentWhere = struct {
	ID      whereHelperint
	Name    whereHelperstring
	ClassID whereHelperint
}{
	ID:      whereHelperint{field: "\"students\".\"id\""},
	Name:    whereHelperstring{field: "\"students\".\"name\""},
	ClassID: whereHelperint{field: "\"students\".\"class_id\""},
}

// StudentRels is where relationship names are stored.
var StudentRels = struct {
	Class string
}{
	Class: "Class",
}

// studentR is where relationships are stored.
type studentR struct {
	Class *Class `boil:"Class" json:"Class" toml:"Class" yaml:"Class"`
}

// NewStruct creates a new relationship struct
func (*studentR) NewStruct() *studentR {
	return &studentR{}
}

func (r *studentR) GetClass() *Class {
	if r == nil {
		return nil
	}
	return r.Class
}

// studentL is where Load methods for each relationship are stored.
type studentL struct{}

var (
	studentAllColumns            = []string{"id", "name", "class_id"}
	studentColumnsWithoutDefault = []string{"name", "class_id"}
	studentColumnsWithDefault    = []string{"id"}
	studentPrimaryKeyColumns     = []string{"id"}
	studentGeneratedColumns      = []string{}
)

type (
	// StudentSlice is an alias for a slice of pointers to Student.
	// This should almost always be used instead of []Student.
	StudentSlice []*Student
	// StudentHook is the signature for custom Student hook methods
	StudentHook func(context.Context, boil.ContextExecutor, *Student) error

	studentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	studentType                 = reflect.TypeOf(&Student{})
	studentMapping              = queries.MakeStructMapping(studentType)
	studentPrimaryKeyMapping, _ = queries.BindMapping(studentType, studentMapping, studentPrimaryKeyColumns)
	studentInsertCacheMut       sync.RWMutex
	studentInsertCache          = make(map[string]insertCache)
	studentUpdateCacheMut       sync.RWMutex
	studentUpdateCache          = make(map[string]updateCache)
	studentUpsertCacheMut       sync.RWMutex
	studentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var studentAfterSelectHooks []StudentHook

var studentBeforeInsertHooks []StudentHook
var studentAfterInsertHooks []StudentHook

var studentBeforeUpdateHooks []StudentHook
var studentAfterUpdateHooks []StudentHook

var studentBeforeDeleteHooks []StudentHook
var studentAfterDeleteHooks []StudentHook

var studentBeforeUpsertHooks []StudentHook
var studentAfterUpsertHooks []StudentHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Student) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range studentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Student) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range studentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Student) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range studentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Student) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range studentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Student) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range studentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Student) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range studentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Student) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range studentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Student) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range studentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Student) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range studentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStudentHook registers your hook function for all future operations.
func AddStudentHook(hookPoint boil.HookPoint, studentHook StudentHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		studentAfterSelectHooks = append(studentAfterSelectHooks, studentHook)
	case boil.BeforeInsertHook:
		studentBeforeInsertHooks = append(studentBeforeInsertHooks, studentHook)
	case boil.AfterInsertHook:
		studentAfterInsertHooks = append(studentAfterInsertHooks, studentHook)
	case boil.BeforeUpdateHook:
		studentBeforeUpdateHooks = append(studentBeforeUpdateHooks, studentHook)
	case boil.AfterUpdateHook:
		studentAfterUpdateHooks = append(studentAfterUpdateHooks, studentHook)
	case boil.BeforeDeleteHook:
		studentBeforeDeleteHooks = append(studentBeforeDeleteHooks, studentHook)
	case boil.AfterDeleteHook:
		studentAfterDeleteHooks = append(studentAfterDeleteHooks, studentHook)
	case boil.BeforeUpsertHook:
		studentBeforeUpsertHooks = append(studentBeforeUpsertHooks, studentHook)
	case boil.AfterUpsertHook:
		studentAfterUpsertHooks = append(studentAfterUpsertHooks, studentHook)
	}
}

// One returns a single student record from the query.
func (q studentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Student, error) {
	o := &Student{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for students")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Student records from the query.
func (q studentQuery) All(ctx context.Context, exec boil.ContextExecutor) (StudentSlice, error) {
	var o []*Student

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Student slice")
	}

	if len(studentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Student records in the query.
func (q studentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count students rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q studentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if students exists")
	}

	return count > 0, nil
}

// Class pointed to by the foreign key.
func (o *Student) Class(mods ...qm.QueryMod) classQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ClassID),
	}

	queryMods = append(queryMods, mods...)

	return Classes(queryMods...)
}

// LoadClass allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (studentL) LoadClass(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStudent interface{}, mods queries.Applicator) error {
	var slice []*Student
	var object *Student

	if singular {
		var ok bool
		object, ok = maybeStudent.(*Student)
		if !ok {
			object = new(Student)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeStudent)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeStudent))
			}
		}
	} else {
		s, ok := maybeStudent.(*[]*Student)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeStudent)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeStudent))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &studentR{}
		}
		args = append(args, object.ClassID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &studentR{}
			}

			for _, a := range args {
				if a == obj.ClassID {
					continue Outer
				}
			}

			args = append(args, obj.ClassID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`classes`),
		qm.WhereIn(`classes.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Class")
	}

	var resultSlice []*Class
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Class")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for classes")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for classes")
	}

	if len(classAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Class = foreign
		if foreign.R == nil {
			foreign.R = &classR{}
		}
		foreign.R.Students = append(foreign.R.Students, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ClassID == foreign.ID {
				local.R.Class = foreign
				if foreign.R == nil {
					foreign.R = &classR{}
				}
				foreign.R.Students = append(foreign.R.Students, local)
				break
			}
		}
	}

	return nil
}

// SetClass of the student to the related item.
// Sets o.R.Class to related.
// Adds o to related.R.Students.
func (o *Student) SetClass(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Class) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"students\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"class_id"}),
		strmangle.WhereClause("\"", "\"", 2, studentPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ClassID = related.ID
	if o.R == nil {
		o.R = &studentR{
			Class: related,
		}
	} else {
		o.R.Class = related
	}

	if related.R == nil {
		related.R = &classR{
			Students: StudentSlice{o},
		}
	} else {
		related.R.Students = append(related.R.Students, o)
	}

	return nil
}

// Students retrieves all the records using an executor.
func Students(mods ...qm.QueryMod) studentQuery {
	mods = append(mods, qm.From("\"students\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"students\".*"})
	}

	return studentQuery{q}
}

// FindStudent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStudent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Student, error) {
	studentObj := &Student{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"students\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, studentObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from students")
	}

	if err = studentObj.doAfterSelectHooks(ctx, exec); err != nil {
		return studentObj, err
	}

	return studentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Student) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no students provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(studentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	studentInsertCacheMut.RLock()
	cache, cached := studentInsertCache[key]
	studentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			studentAllColumns,
			studentColumnsWithDefault,
			studentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(studentType, studentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(studentType, studentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"students\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"students\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into students")
	}

	if !cached {
		studentInsertCacheMut.Lock()
		studentInsertCache[key] = cache
		studentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Student.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Student) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	studentUpdateCacheMut.RLock()
	cache, cached := studentUpdateCache[key]
	studentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			studentAllColumns,
			studentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update students, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"students\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, studentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(studentType, studentMapping, append(wl, studentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update students row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for students")
	}

	if !cached {
		studentUpdateCacheMut.Lock()
		studentUpdateCache[key] = cache
		studentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q studentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for students")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for students")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StudentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), studentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"students\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, studentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in student slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all student")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Student) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no students provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(studentColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	studentUpsertCacheMut.RLock()
	cache, cached := studentUpsertCache[key]
	studentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			studentAllColumns,
			studentColumnsWithDefault,
			studentColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			studentAllColumns,
			studentPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert students, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(studentPrimaryKeyColumns))
			copy(conflict, studentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"students\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(studentType, studentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(studentType, studentMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert students")
	}

	if !cached {
		studentUpsertCacheMut.Lock()
		studentUpsertCache[key] = cache
		studentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Student record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Student) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Student provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), studentPrimaryKeyMapping)
	sql := "DELETE FROM \"students\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from students")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for students")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q studentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no studentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from students")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for students")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StudentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(studentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), studentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"students\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, studentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from student slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for students")
	}

	if len(studentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Student) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStudent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StudentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StudentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), studentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"students\".* FROM \"students\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, studentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StudentSlice")
	}

	*o = slice

	return nil
}

// StudentExists checks if the Student row exists.
func StudentExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"students\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if students exists")
	}

	return exists, nil
}

// Exists checks if the Student row exists.
func (o *Student) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return StudentExists(ctx, exec, o.ID)
}
