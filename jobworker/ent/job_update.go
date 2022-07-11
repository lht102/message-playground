// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/ent/job"
	"github.com/lht102/message-playground/jobworker/ent/predicate"
)

// JobUpdate is the builder for updating Job entities.
type JobUpdate struct {
	config
	hooks    []Hook
	mutation *JobMutation
}

// Where appends a list predicates to the JobUpdate builder.
func (ju *JobUpdate) Where(ps ...predicate.Job) *JobUpdate {
	ju.mutation.Where(ps...)
	return ju
}

// SetState sets the "state" field.
func (ju *JobUpdate) SetState(js jobworker.JobState) *JobUpdate {
	ju.mutation.SetState(js)
	return ju
}

// SetDescription sets the "description" field.
func (ju *JobUpdate) SetDescription(s string) *JobUpdate {
	ju.mutation.SetDescription(s)
	return ju
}

// SetCompletedAt sets the "completed_at" field.
func (ju *JobUpdate) SetCompletedAt(t time.Time) *JobUpdate {
	ju.mutation.SetCompletedAt(t)
	return ju
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (ju *JobUpdate) SetNillableCompletedAt(t *time.Time) *JobUpdate {
	if t != nil {
		ju.SetCompletedAt(*t)
	}
	return ju
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (ju *JobUpdate) ClearCompletedAt() *JobUpdate {
	ju.mutation.ClearCompletedAt()
	return ju
}

// SetCreatedAt sets the "created_at" field.
func (ju *JobUpdate) SetCreatedAt(t time.Time) *JobUpdate {
	ju.mutation.SetCreatedAt(t)
	return ju
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ju *JobUpdate) SetNillableCreatedAt(t *time.Time) *JobUpdate {
	if t != nil {
		ju.SetCreatedAt(*t)
	}
	return ju
}

// SetUpdatedAt sets the "updated_at" field.
func (ju *JobUpdate) SetUpdatedAt(t time.Time) *JobUpdate {
	ju.mutation.SetUpdatedAt(t)
	return ju
}

// Mutation returns the JobMutation object of the builder.
func (ju *JobUpdate) Mutation() *JobMutation {
	return ju.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ju *JobUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ju.defaults()
	if len(ju.hooks) == 0 {
		if err = ju.check(); err != nil {
			return 0, err
		}
		affected, err = ju.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ju.check(); err != nil {
				return 0, err
			}
			ju.mutation = mutation
			affected, err = ju.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ju.hooks) - 1; i >= 0; i-- {
			if ju.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ju.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ju.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ju *JobUpdate) SaveX(ctx context.Context) int {
	affected, err := ju.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ju *JobUpdate) Exec(ctx context.Context) error {
	_, err := ju.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ju *JobUpdate) ExecX(ctx context.Context) {
	if err := ju.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ju *JobUpdate) defaults() {
	if _, ok := ju.mutation.UpdatedAt(); !ok {
		v := job.UpdateDefaultUpdatedAt()
		ju.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ju *JobUpdate) check() error {
	if v, ok := ju.mutation.State(); ok {
		if err := job.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Job.state": %w`, err)}
		}
	}
	if v, ok := ju.mutation.Description(); ok {
		if err := job.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Job.description": %w`, err)}
		}
	}
	return nil
}

func (ju *JobUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   job.Table,
			Columns: job.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: job.FieldID,
			},
		},
	}
	if ps := ju.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ju.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: job.FieldState,
		})
	}
	if value, ok := ju.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldDescription,
		})
	}
	if value, ok := ju.mutation.CompletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: job.FieldCompletedAt,
		})
	}
	if ju.mutation.CompletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: job.FieldCompletedAt,
		})
	}
	if value, ok := ju.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: job.FieldCreatedAt,
		})
	}
	if value, ok := ju.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: job.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ju.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{job.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// JobUpdateOne is the builder for updating a single Job entity.
type JobUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JobMutation
}

// SetState sets the "state" field.
func (juo *JobUpdateOne) SetState(js jobworker.JobState) *JobUpdateOne {
	juo.mutation.SetState(js)
	return juo
}

// SetDescription sets the "description" field.
func (juo *JobUpdateOne) SetDescription(s string) *JobUpdateOne {
	juo.mutation.SetDescription(s)
	return juo
}

// SetCompletedAt sets the "completed_at" field.
func (juo *JobUpdateOne) SetCompletedAt(t time.Time) *JobUpdateOne {
	juo.mutation.SetCompletedAt(t)
	return juo
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (juo *JobUpdateOne) SetNillableCompletedAt(t *time.Time) *JobUpdateOne {
	if t != nil {
		juo.SetCompletedAt(*t)
	}
	return juo
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (juo *JobUpdateOne) ClearCompletedAt() *JobUpdateOne {
	juo.mutation.ClearCompletedAt()
	return juo
}

// SetCreatedAt sets the "created_at" field.
func (juo *JobUpdateOne) SetCreatedAt(t time.Time) *JobUpdateOne {
	juo.mutation.SetCreatedAt(t)
	return juo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (juo *JobUpdateOne) SetNillableCreatedAt(t *time.Time) *JobUpdateOne {
	if t != nil {
		juo.SetCreatedAt(*t)
	}
	return juo
}

// SetUpdatedAt sets the "updated_at" field.
func (juo *JobUpdateOne) SetUpdatedAt(t time.Time) *JobUpdateOne {
	juo.mutation.SetUpdatedAt(t)
	return juo
}

// Mutation returns the JobMutation object of the builder.
func (juo *JobUpdateOne) Mutation() *JobMutation {
	return juo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (juo *JobUpdateOne) Select(field string, fields ...string) *JobUpdateOne {
	juo.fields = append([]string{field}, fields...)
	return juo
}

// Save executes the query and returns the updated Job entity.
func (juo *JobUpdateOne) Save(ctx context.Context) (*Job, error) {
	var (
		err  error
		node *Job
	)
	juo.defaults()
	if len(juo.hooks) == 0 {
		if err = juo.check(); err != nil {
			return nil, err
		}
		node, err = juo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = juo.check(); err != nil {
				return nil, err
			}
			juo.mutation = mutation
			node, err = juo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(juo.hooks) - 1; i >= 0; i-- {
			if juo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = juo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, juo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (juo *JobUpdateOne) SaveX(ctx context.Context) *Job {
	node, err := juo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (juo *JobUpdateOne) Exec(ctx context.Context) error {
	_, err := juo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (juo *JobUpdateOne) ExecX(ctx context.Context) {
	if err := juo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (juo *JobUpdateOne) defaults() {
	if _, ok := juo.mutation.UpdatedAt(); !ok {
		v := job.UpdateDefaultUpdatedAt()
		juo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (juo *JobUpdateOne) check() error {
	if v, ok := juo.mutation.State(); ok {
		if err := job.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Job.state": %w`, err)}
		}
	}
	if v, ok := juo.mutation.Description(); ok {
		if err := job.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Job.description": %w`, err)}
		}
	}
	return nil
}

func (juo *JobUpdateOne) sqlSave(ctx context.Context) (_node *Job, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   job.Table,
			Columns: job.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: job.FieldID,
			},
		},
	}
	id, ok := juo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Job.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := juo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, job.FieldID)
		for _, f := range fields {
			if !job.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != job.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := juo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := juo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: job.FieldState,
		})
	}
	if value, ok := juo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldDescription,
		})
	}
	if value, ok := juo.mutation.CompletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: job.FieldCompletedAt,
		})
	}
	if juo.mutation.CompletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: job.FieldCompletedAt,
		})
	}
	if value, ok := juo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: job.FieldCreatedAt,
		})
	}
	if value, ok := juo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: job.FieldUpdatedAt,
		})
	}
	_node = &Job{config: juo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, juo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{job.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
