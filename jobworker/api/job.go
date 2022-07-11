package api

import (
	"time"

	"github.com/google/uuid"
)

type JobState string

const (
	JobStateQueued     JobState = "QUEUED"
	JobStateProcessing JobState = "PROCESSING"
	JobStateCompleted  JobState = "COMPLETED"
)

type CreateJobRequest struct {
	RequestUUID uuid.UUID `json:"request_uuid"`
	Description string    `json:"description"`
}

type CreateJobResponse struct {
	JobUUID uuid.UUID `json:"job_uuid"`
}

type JobResponse struct {
	UUID        uuid.UUID  `json:"uuid"`
	RequestUUID uuid.UUID  `json:"request_uuid"`
	State       JobState   `json:"state"`
	Description string     `json:"description"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
