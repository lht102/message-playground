package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	jobworker "github.com/lht102/message-playground/jobworker"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("request_uuid", uuid.UUID{}).
			Unique().
			Immutable(),
		field.Enum("state").
			GoType(jobworker.JobStateQueued),
		field.String("description").
			MaxLen(255).
			NotEmpty(),
		field.Time("completed_at").
			Optional().
			Nillable(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}
