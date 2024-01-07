// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-member-rpc/ent/memberrank"
	"github.com/suyuan32/simple-admin-member-rpc/ent/predicate"
)

// MemberRankDelete is the builder for deleting a MemberRank entity.
type MemberRankDelete struct {
	config
	hooks    []Hook
	mutation *MemberRankMutation
}

// Where appends a list predicates to the MemberRankDelete builder.
func (mrd *MemberRankDelete) Where(ps ...predicate.MemberRank) *MemberRankDelete {
	mrd.mutation.Where(ps...)
	return mrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mrd *MemberRankDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mrd.sqlExec, mrd.mutation, mrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mrd *MemberRankDelete) ExecX(ctx context.Context) int {
	n, err := mrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mrd *MemberRankDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(memberrank.Table, sqlgraph.NewFieldSpec(memberrank.FieldID, field.TypeUint64))
	if ps := mrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mrd.mutation.done = true
	return affected, err
}

// MemberRankDeleteOne is the builder for deleting a single MemberRank entity.
type MemberRankDeleteOne struct {
	mrd *MemberRankDelete
}

// Where appends a list predicates to the MemberRankDelete builder.
func (mrdo *MemberRankDeleteOne) Where(ps ...predicate.MemberRank) *MemberRankDeleteOne {
	mrdo.mrd.mutation.Where(ps...)
	return mrdo
}

// Exec executes the deletion query.
func (mrdo *MemberRankDeleteOne) Exec(ctx context.Context) error {
	n, err := mrdo.mrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{memberrank.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mrdo *MemberRankDeleteOne) ExecX(ctx context.Context) {
	if err := mrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
