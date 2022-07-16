// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/ent/job"
	"github.com/lht102/message-playground/jobworker/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeJob = "Job"
)

// JobMutation represents an operation that mutates the Job nodes in the graph.
type JobMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	request_uuid  *uuid.UUID
	state         *jobworker.JobState
	description   *string
	completed_at  *time.Time
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Job, error)
	predicates    []predicate.Job
}

var _ ent.Mutation = (*JobMutation)(nil)

// jobOption allows management of the mutation configuration using functional options.
type jobOption func(*JobMutation)

// newJobMutation creates new mutation for the Job entity.
func newJobMutation(c config, op Op, opts ...jobOption) *JobMutation {
	m := &JobMutation{
		config:        c,
		op:            op,
		typ:           TypeJob,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withJobID sets the ID field of the mutation.
func withJobID(id uuid.UUID) jobOption {
	return func(m *JobMutation) {
		var (
			err   error
			once  sync.Once
			value *Job
		)
		m.oldValue = func(ctx context.Context) (*Job, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Job.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withJob sets the old Job of the mutation.
func withJob(node *Job) jobOption {
	return func(m *JobMutation) {
		m.oldValue = func(context.Context) (*Job, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m JobMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m JobMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Job entities.
func (m *JobMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *JobMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *JobMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Job.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetRequestUUID sets the "request_uuid" field.
func (m *JobMutation) SetRequestUUID(u uuid.UUID) {
	m.request_uuid = &u
}

// RequestUUID returns the value of the "request_uuid" field in the mutation.
func (m *JobMutation) RequestUUID() (r uuid.UUID, exists bool) {
	v := m.request_uuid
	if v == nil {
		return
	}
	return *v, true
}

// OldRequestUUID returns the old "request_uuid" field's value of the Job entity.
// If the Job object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobMutation) OldRequestUUID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRequestUUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRequestUUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRequestUUID: %w", err)
	}
	return oldValue.RequestUUID, nil
}

// ResetRequestUUID resets all changes to the "request_uuid" field.
func (m *JobMutation) ResetRequestUUID() {
	m.request_uuid = nil
}

// SetState sets the "state" field.
func (m *JobMutation) SetState(js jobworker.JobState) {
	m.state = &js
}

// State returns the value of the "state" field in the mutation.
func (m *JobMutation) State() (r jobworker.JobState, exists bool) {
	v := m.state
	if v == nil {
		return
	}
	return *v, true
}

// OldState returns the old "state" field's value of the Job entity.
// If the Job object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobMutation) OldState(ctx context.Context) (v jobworker.JobState, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldState is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldState requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldState: %w", err)
	}
	return oldValue.State, nil
}

// ResetState resets all changes to the "state" field.
func (m *JobMutation) ResetState() {
	m.state = nil
}

// SetDescription sets the "description" field.
func (m *JobMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *JobMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Job entity.
// If the Job object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *JobMutation) ResetDescription() {
	m.description = nil
}

// SetCompletedAt sets the "completed_at" field.
func (m *JobMutation) SetCompletedAt(t time.Time) {
	m.completed_at = &t
}

// CompletedAt returns the value of the "completed_at" field in the mutation.
func (m *JobMutation) CompletedAt() (r time.Time, exists bool) {
	v := m.completed_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCompletedAt returns the old "completed_at" field's value of the Job entity.
// If the Job object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobMutation) OldCompletedAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCompletedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCompletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCompletedAt: %w", err)
	}
	return oldValue.CompletedAt, nil
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (m *JobMutation) ClearCompletedAt() {
	m.completed_at = nil
	m.clearedFields[job.FieldCompletedAt] = struct{}{}
}

// CompletedAtCleared returns if the "completed_at" field was cleared in this mutation.
func (m *JobMutation) CompletedAtCleared() bool {
	_, ok := m.clearedFields[job.FieldCompletedAt]
	return ok
}

// ResetCompletedAt resets all changes to the "completed_at" field.
func (m *JobMutation) ResetCompletedAt() {
	m.completed_at = nil
	delete(m.clearedFields, job.FieldCompletedAt)
}

// SetCreatedAt sets the "created_at" field.
func (m *JobMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *JobMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Job entity.
// If the Job object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *JobMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *JobMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *JobMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Job entity.
// If the Job object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *JobMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *JobMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the JobMutation builder.
func (m *JobMutation) Where(ps ...predicate.Job) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *JobMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Job).
func (m *JobMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *JobMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.request_uuid != nil {
		fields = append(fields, job.FieldRequestUUID)
	}
	if m.state != nil {
		fields = append(fields, job.FieldState)
	}
	if m.description != nil {
		fields = append(fields, job.FieldDescription)
	}
	if m.completed_at != nil {
		fields = append(fields, job.FieldCompletedAt)
	}
	if m.created_at != nil {
		fields = append(fields, job.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, job.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *JobMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case job.FieldRequestUUID:
		return m.RequestUUID()
	case job.FieldState:
		return m.State()
	case job.FieldDescription:
		return m.Description()
	case job.FieldCompletedAt:
		return m.CompletedAt()
	case job.FieldCreatedAt:
		return m.CreatedAt()
	case job.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *JobMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case job.FieldRequestUUID:
		return m.OldRequestUUID(ctx)
	case job.FieldState:
		return m.OldState(ctx)
	case job.FieldDescription:
		return m.OldDescription(ctx)
	case job.FieldCompletedAt:
		return m.OldCompletedAt(ctx)
	case job.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case job.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Job field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *JobMutation) SetField(name string, value ent.Value) error {
	switch name {
	case job.FieldRequestUUID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRequestUUID(v)
		return nil
	case job.FieldState:
		v, ok := value.(jobworker.JobState)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetState(v)
		return nil
	case job.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case job.FieldCompletedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCompletedAt(v)
		return nil
	case job.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case job.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Job field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *JobMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *JobMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *JobMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Job numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *JobMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(job.FieldCompletedAt) {
		fields = append(fields, job.FieldCompletedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *JobMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *JobMutation) ClearField(name string) error {
	switch name {
	case job.FieldCompletedAt:
		m.ClearCompletedAt()
		return nil
	}
	return fmt.Errorf("unknown Job nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *JobMutation) ResetField(name string) error {
	switch name {
	case job.FieldRequestUUID:
		m.ResetRequestUUID()
		return nil
	case job.FieldState:
		m.ResetState()
		return nil
	case job.FieldDescription:
		m.ResetDescription()
		return nil
	case job.FieldCompletedAt:
		m.ResetCompletedAt()
		return nil
	case job.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case job.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Job field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *JobMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *JobMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *JobMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *JobMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *JobMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *JobMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *JobMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Job unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *JobMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Job edge %s", name)
}
