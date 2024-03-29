// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// MailMagazineLog is an object representing the database table.
type MailMagazineLog struct {
	ID             int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	MailMagazineID string    `boil:"mail_magazine_id" json:"mail_magazine_id" toml:"mail_magazine_id" yaml:"mail_magazine_id"`
	UserID         string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Email          string    `boil:"email" json:"email" toml:"email" yaml:"email"`
	SentAt         null.Time `boil:"sent_at" json:"sent_at,omitempty" toml:"sent_at" yaml:"sent_at,omitempty"`
	CreateAt       time.Time `boil:"create_at" json:"create_at" toml:"create_at" yaml:"create_at"`
	UpdateAt       time.Time `boil:"update_at" json:"update_at" toml:"update_at" yaml:"update_at"`

	R *mailMagazineLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mailMagazineLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MailMagazineLogColumns = struct {
	ID             string
	MailMagazineID string
	UserID         string
	Email          string
	SentAt         string
	CreateAt       string
	UpdateAt       string
}{
	ID:             "id",
	MailMagazineID: "mail_magazine_id",
	UserID:         "user_id",
	Email:          "email",
	SentAt:         "sent_at",
	CreateAt:       "create_at",
	UpdateAt:       "update_at",
}

var MailMagazineLogTableColumns = struct {
	ID             string
	MailMagazineID string
	UserID         string
	Email          string
	SentAt         string
	CreateAt       string
	UpdateAt       string
}{
	ID:             "mail_magazine_log.id",
	MailMagazineID: "mail_magazine_log.mail_magazine_id",
	UserID:         "mail_magazine_log.user_id",
	Email:          "mail_magazine_log.email",
	SentAt:         "mail_magazine_log.sent_at",
	CreateAt:       "mail_magazine_log.create_at",
	UpdateAt:       "mail_magazine_log.update_at",
}

// Generated where

var MailMagazineLogWhere = struct {
	ID             whereHelperint64
	MailMagazineID whereHelperstring
	UserID         whereHelperstring
	Email          whereHelperstring
	SentAt         whereHelpernull_Time
	CreateAt       whereHelpertime_Time
	UpdateAt       whereHelpertime_Time
}{
	ID:             whereHelperint64{field: "\"mail_magazine_log\".\"id\""},
	MailMagazineID: whereHelperstring{field: "\"mail_magazine_log\".\"mail_magazine_id\""},
	UserID:         whereHelperstring{field: "\"mail_magazine_log\".\"user_id\""},
	Email:          whereHelperstring{field: "\"mail_magazine_log\".\"email\""},
	SentAt:         whereHelpernull_Time{field: "\"mail_magazine_log\".\"sent_at\""},
	CreateAt:       whereHelpertime_Time{field: "\"mail_magazine_log\".\"create_at\""},
	UpdateAt:       whereHelpertime_Time{field: "\"mail_magazine_log\".\"update_at\""},
}

// MailMagazineLogRels is where relationship names are stored.
var MailMagazineLogRels = struct {
	MailMagazine string
	User         string
}{
	MailMagazine: "MailMagazine",
	User:         "User",
}

// mailMagazineLogR is where relationships are stored.
type mailMagazineLogR struct {
	MailMagazine *MailMagazine `boil:"MailMagazine" json:"MailMagazine" toml:"MailMagazine" yaml:"MailMagazine"`
	User         *UserDatum    `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*mailMagazineLogR) NewStruct() *mailMagazineLogR {
	return &mailMagazineLogR{}
}

func (r *mailMagazineLogR) GetMailMagazine() *MailMagazine {
	if r == nil {
		return nil
	}
	return r.MailMagazine
}

func (r *mailMagazineLogR) GetUser() *UserDatum {
	if r == nil {
		return nil
	}
	return r.User
}

// mailMagazineLogL is where Load methods for each relationship are stored.
type mailMagazineLogL struct{}

var (
	mailMagazineLogAllColumns            = []string{"id", "mail_magazine_id", "user_id", "email", "sent_at", "create_at", "update_at"}
	mailMagazineLogColumnsWithoutDefault = []string{"mail_magazine_id", "user_id", "email"}
	mailMagazineLogColumnsWithDefault    = []string{"id", "sent_at", "create_at", "update_at"}
	mailMagazineLogPrimaryKeyColumns     = []string{"id"}
	mailMagazineLogGeneratedColumns      = []string{}
)

type (
	// MailMagazineLogSlice is an alias for a slice of pointers to MailMagazineLog.
	// This should almost always be used instead of []MailMagazineLog.
	MailMagazineLogSlice []*MailMagazineLog
	// MailMagazineLogHook is the signature for custom MailMagazineLog hook methods
	MailMagazineLogHook func(context.Context, boil.ContextExecutor, *MailMagazineLog) error

	mailMagazineLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mailMagazineLogType                 = reflect.TypeOf(&MailMagazineLog{})
	mailMagazineLogMapping              = queries.MakeStructMapping(mailMagazineLogType)
	mailMagazineLogPrimaryKeyMapping, _ = queries.BindMapping(mailMagazineLogType, mailMagazineLogMapping, mailMagazineLogPrimaryKeyColumns)
	mailMagazineLogInsertCacheMut       sync.RWMutex
	mailMagazineLogInsertCache          = make(map[string]insertCache)
	mailMagazineLogUpdateCacheMut       sync.RWMutex
	mailMagazineLogUpdateCache          = make(map[string]updateCache)
	mailMagazineLogUpsertCacheMut       sync.RWMutex
	mailMagazineLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var mailMagazineLogAfterSelectHooks []MailMagazineLogHook

var mailMagazineLogBeforeInsertHooks []MailMagazineLogHook
var mailMagazineLogAfterInsertHooks []MailMagazineLogHook

var mailMagazineLogBeforeUpdateHooks []MailMagazineLogHook
var mailMagazineLogAfterUpdateHooks []MailMagazineLogHook

var mailMagazineLogBeforeDeleteHooks []MailMagazineLogHook
var mailMagazineLogAfterDeleteHooks []MailMagazineLogHook

var mailMagazineLogBeforeUpsertHooks []MailMagazineLogHook
var mailMagazineLogAfterUpsertHooks []MailMagazineLogHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MailMagazineLog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineLogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MailMagazineLog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineLogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MailMagazineLog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineLogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MailMagazineLog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineLogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MailMagazineLog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineLogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MailMagazineLog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineLogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MailMagazineLog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineLogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MailMagazineLog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineLogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MailMagazineLog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mailMagazineLogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMailMagazineLogHook registers your hook function for all future operations.
func AddMailMagazineLogHook(hookPoint boil.HookPoint, mailMagazineLogHook MailMagazineLogHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		mailMagazineLogAfterSelectHooks = append(mailMagazineLogAfterSelectHooks, mailMagazineLogHook)
	case boil.BeforeInsertHook:
		mailMagazineLogBeforeInsertHooks = append(mailMagazineLogBeforeInsertHooks, mailMagazineLogHook)
	case boil.AfterInsertHook:
		mailMagazineLogAfterInsertHooks = append(mailMagazineLogAfterInsertHooks, mailMagazineLogHook)
	case boil.BeforeUpdateHook:
		mailMagazineLogBeforeUpdateHooks = append(mailMagazineLogBeforeUpdateHooks, mailMagazineLogHook)
	case boil.AfterUpdateHook:
		mailMagazineLogAfterUpdateHooks = append(mailMagazineLogAfterUpdateHooks, mailMagazineLogHook)
	case boil.BeforeDeleteHook:
		mailMagazineLogBeforeDeleteHooks = append(mailMagazineLogBeforeDeleteHooks, mailMagazineLogHook)
	case boil.AfterDeleteHook:
		mailMagazineLogAfterDeleteHooks = append(mailMagazineLogAfterDeleteHooks, mailMagazineLogHook)
	case boil.BeforeUpsertHook:
		mailMagazineLogBeforeUpsertHooks = append(mailMagazineLogBeforeUpsertHooks, mailMagazineLogHook)
	case boil.AfterUpsertHook:
		mailMagazineLogAfterUpsertHooks = append(mailMagazineLogAfterUpsertHooks, mailMagazineLogHook)
	}
}

// One returns a single mailMagazineLog record from the query.
func (q mailMagazineLogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MailMagazineLog, error) {
	o := &MailMagazineLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for mail_magazine_log")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MailMagazineLog records from the query.
func (q mailMagazineLogQuery) All(ctx context.Context, exec boil.ContextExecutor) (MailMagazineLogSlice, error) {
	var o []*MailMagazineLog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MailMagazineLog slice")
	}

	if len(mailMagazineLogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MailMagazineLog records in the query.
func (q mailMagazineLogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count mail_magazine_log rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mailMagazineLogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if mail_magazine_log exists")
	}

	return count > 0, nil
}

// MailMagazine pointed to by the foreign key.
func (o *MailMagazineLog) MailMagazine(mods ...qm.QueryMod) mailMagazineQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.MailMagazineID),
	}

	queryMods = append(queryMods, mods...)

	return MailMagazines(queryMods...)
}

// User pointed to by the foreign key.
func (o *MailMagazineLog) User(mods ...qm.QueryMod) userDatumQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"user_id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return UserData(queryMods...)
}

// LoadMailMagazine allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (mailMagazineLogL) LoadMailMagazine(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMailMagazineLog interface{}, mods queries.Applicator) error {
	var slice []*MailMagazineLog
	var object *MailMagazineLog

	if singular {
		var ok bool
		object, ok = maybeMailMagazineLog.(*MailMagazineLog)
		if !ok {
			object = new(MailMagazineLog)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMailMagazineLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMailMagazineLog))
			}
		}
	} else {
		s, ok := maybeMailMagazineLog.(*[]*MailMagazineLog)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMailMagazineLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMailMagazineLog))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &mailMagazineLogR{}
		}
		args = append(args, object.MailMagazineID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &mailMagazineLogR{}
			}

			for _, a := range args {
				if a == obj.MailMagazineID {
					continue Outer
				}
			}

			args = append(args, obj.MailMagazineID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`mail_magazine`),
		qm.WhereIn(`mail_magazine.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load MailMagazine")
	}

	var resultSlice []*MailMagazine
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice MailMagazine")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for mail_magazine")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for mail_magazine")
	}

	if len(mailMagazineAfterSelectHooks) != 0 {
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
		object.R.MailMagazine = foreign
		if foreign.R == nil {
			foreign.R = &mailMagazineR{}
		}
		foreign.R.MailMagazineLogs = append(foreign.R.MailMagazineLogs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MailMagazineID == foreign.ID {
				local.R.MailMagazine = foreign
				if foreign.R == nil {
					foreign.R = &mailMagazineR{}
				}
				foreign.R.MailMagazineLogs = append(foreign.R.MailMagazineLogs, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (mailMagazineLogL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMailMagazineLog interface{}, mods queries.Applicator) error {
	var slice []*MailMagazineLog
	var object *MailMagazineLog

	if singular {
		var ok bool
		object, ok = maybeMailMagazineLog.(*MailMagazineLog)
		if !ok {
			object = new(MailMagazineLog)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMailMagazineLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMailMagazineLog))
			}
		}
	} else {
		s, ok := maybeMailMagazineLog.(*[]*MailMagazineLog)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMailMagazineLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMailMagazineLog))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &mailMagazineLogR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &mailMagazineLogR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_data`),
		qm.WhereIn(`user_data.user_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserDatum")
	}

	var resultSlice []*UserDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserDatum")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_data")
	}

	if len(userDatumAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userDatumR{}
		}
		foreign.R.UserMailMagazineLogs = append(foreign.R.UserMailMagazineLogs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userDatumR{}
				}
				foreign.R.UserMailMagazineLogs = append(foreign.R.UserMailMagazineLogs, local)
				break
			}
		}
	}

	return nil
}

// SetMailMagazine of the mailMagazineLog to the related item.
// Sets o.R.MailMagazine to related.
// Adds o to related.R.MailMagazineLogs.
func (o *MailMagazineLog) SetMailMagazine(ctx context.Context, exec boil.ContextExecutor, insert bool, related *MailMagazine) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"mail_magazine_log\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"mail_magazine_id"}),
		strmangle.WhereClause("\"", "\"", 2, mailMagazineLogPrimaryKeyColumns),
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

	o.MailMagazineID = related.ID
	if o.R == nil {
		o.R = &mailMagazineLogR{
			MailMagazine: related,
		}
	} else {
		o.R.MailMagazine = related
	}

	if related.R == nil {
		related.R = &mailMagazineR{
			MailMagazineLogs: MailMagazineLogSlice{o},
		}
	} else {
		related.R.MailMagazineLogs = append(related.R.MailMagazineLogs, o)
	}

	return nil
}

// SetUser of the mailMagazineLog to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserMailMagazineLogs.
func (o *MailMagazineLog) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *UserDatum) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"mail_magazine_log\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, mailMagazineLogPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.UserID
	if o.R == nil {
		o.R = &mailMagazineLogR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userDatumR{
			UserMailMagazineLogs: MailMagazineLogSlice{o},
		}
	} else {
		related.R.UserMailMagazineLogs = append(related.R.UserMailMagazineLogs, o)
	}

	return nil
}

// MailMagazineLogs retrieves all the records using an executor.
func MailMagazineLogs(mods ...qm.QueryMod) mailMagazineLogQuery {
	mods = append(mods, qm.From("\"mail_magazine_log\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"mail_magazine_log\".*"})
	}

	return mailMagazineLogQuery{q}
}

// FindMailMagazineLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMailMagazineLog(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*MailMagazineLog, error) {
	mailMagazineLogObj := &MailMagazineLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"mail_magazine_log\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, mailMagazineLogObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from mail_magazine_log")
	}

	if err = mailMagazineLogObj.doAfterSelectHooks(ctx, exec); err != nil {
		return mailMagazineLogObj, err
	}

	return mailMagazineLogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MailMagazineLog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mail_magazine_log provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreateAt.IsZero() {
			o.CreateAt = currTime
		}
		if o.UpdateAt.IsZero() {
			o.UpdateAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mailMagazineLogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mailMagazineLogInsertCacheMut.RLock()
	cache, cached := mailMagazineLogInsertCache[key]
	mailMagazineLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mailMagazineLogAllColumns,
			mailMagazineLogColumnsWithDefault,
			mailMagazineLogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mailMagazineLogType, mailMagazineLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mailMagazineLogType, mailMagazineLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"mail_magazine_log\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"mail_magazine_log\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into mail_magazine_log")
	}

	if !cached {
		mailMagazineLogInsertCacheMut.Lock()
		mailMagazineLogInsertCache[key] = cache
		mailMagazineLogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MailMagazineLog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MailMagazineLog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdateAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	mailMagazineLogUpdateCacheMut.RLock()
	cache, cached := mailMagazineLogUpdateCache[key]
	mailMagazineLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mailMagazineLogAllColumns,
			mailMagazineLogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update mail_magazine_log, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"mail_magazine_log\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, mailMagazineLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mailMagazineLogType, mailMagazineLogMapping, append(wl, mailMagazineLogPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update mail_magazine_log row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for mail_magazine_log")
	}

	if !cached {
		mailMagazineLogUpdateCacheMut.Lock()
		mailMagazineLogUpdateCache[key] = cache
		mailMagazineLogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q mailMagazineLogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for mail_magazine_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for mail_magazine_log")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MailMagazineLogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mailMagazineLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"mail_magazine_log\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, mailMagazineLogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in mailMagazineLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all mailMagazineLog")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MailMagazineLog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mail_magazine_log provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreateAt.IsZero() {
			o.CreateAt = currTime
		}
		o.UpdateAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mailMagazineLogColumnsWithDefault, o)

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

	mailMagazineLogUpsertCacheMut.RLock()
	cache, cached := mailMagazineLogUpsertCache[key]
	mailMagazineLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mailMagazineLogAllColumns,
			mailMagazineLogColumnsWithDefault,
			mailMagazineLogColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			mailMagazineLogAllColumns,
			mailMagazineLogPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert mail_magazine_log, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(mailMagazineLogPrimaryKeyColumns))
			copy(conflict, mailMagazineLogPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"mail_magazine_log\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(mailMagazineLogType, mailMagazineLogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mailMagazineLogType, mailMagazineLogMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert mail_magazine_log")
	}

	if !cached {
		mailMagazineLogUpsertCacheMut.Lock()
		mailMagazineLogUpsertCache[key] = cache
		mailMagazineLogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MailMagazineLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MailMagazineLog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MailMagazineLog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mailMagazineLogPrimaryKeyMapping)
	sql := "DELETE FROM \"mail_magazine_log\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from mail_magazine_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for mail_magazine_log")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mailMagazineLogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no mailMagazineLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mail_magazine_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mail_magazine_log")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MailMagazineLogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(mailMagazineLogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mailMagazineLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"mail_magazine_log\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mailMagazineLogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mailMagazineLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mail_magazine_log")
	}

	if len(mailMagazineLogAfterDeleteHooks) != 0 {
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
func (o *MailMagazineLog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMailMagazineLog(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MailMagazineLogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MailMagazineLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mailMagazineLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"mail_magazine_log\".* FROM \"mail_magazine_log\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mailMagazineLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MailMagazineLogSlice")
	}

	*o = slice

	return nil
}

// MailMagazineLogExists checks if the MailMagazineLog row exists.
func MailMagazineLogExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"mail_magazine_log\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if mail_magazine_log exists")
	}

	return exists, nil
}

// Exists checks if the MailMagazineLog row exists.
func (o *MailMagazineLog) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MailMagazineLogExists(ctx, exec, o.ID)
}
