// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/database64128/proxy-sharing-go/ent/predicate"
	"github.com/database64128/proxy-sharing-go/ent/registrationtoken"
)

// RegistrationTokenDelete is the builder for deleting a RegistrationToken entity.
type RegistrationTokenDelete struct {
	config
	hooks    []Hook
	mutation *RegistrationTokenMutation
}

// Where appends a list predicates to the RegistrationTokenDelete builder.
func (rtd *RegistrationTokenDelete) Where(ps ...predicate.RegistrationToken) *RegistrationTokenDelete {
	rtd.mutation.Where(ps...)
	return rtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rtd *RegistrationTokenDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rtd.sqlExec, rtd.mutation, rtd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rtd *RegistrationTokenDelete) ExecX(ctx context.Context) int {
	n, err := rtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rtd *RegistrationTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(registrationtoken.Table, sqlgraph.NewFieldSpec(registrationtoken.FieldID, field.TypeInt))
	if ps := rtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rtd.mutation.done = true
	return affected, err
}

// RegistrationTokenDeleteOne is the builder for deleting a single RegistrationToken entity.
type RegistrationTokenDeleteOne struct {
	rtd *RegistrationTokenDelete
}

// Where appends a list predicates to the RegistrationTokenDelete builder.
func (rtdo *RegistrationTokenDeleteOne) Where(ps ...predicate.RegistrationToken) *RegistrationTokenDeleteOne {
	rtdo.rtd.mutation.Where(ps...)
	return rtdo
}

// Exec executes the deletion query.
func (rtdo *RegistrationTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := rtdo.rtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{registrationtoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rtdo *RegistrationTokenDeleteOne) ExecX(ctx context.Context) {
	if err := rtdo.Exec(ctx); err != nil {
		panic(err)
	}
}
