package jobworker

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/lht102/message-playground/jobworker/api"
)

type JobService interface {
	CreateJob(ctx context.Context, createJobCmd CreateJobCommand) (Job, error)
	GetJob(ctx context.Context, uuid uuid.UUID) (Job, error)
	ExecuteJob(ctx context.Context, uuid uuid.UUID) error
}

type MessageBus interface {
	Producer
	Consumer
}

type Producer interface {
	PublishJob(ctx context.Context, job api.JobResponse) error
}

type Consumer interface {
	SubscribeJob(
		ctx context.Context,
		queue string,
		do func(job api.JobResponse) error,
		states ...api.JobState,
	) error
}

type JobState string

const (
	JobStateQueued     JobState = "QUEUED"
	JobStateProcessing JobState = "PROCESSING"
	JobStateCompleted  JobState = "COMPLETED"
)

func (state JobState) Values() []string {
	return []string{string(JobStateQueued), string(JobStateProcessing), string(JobStateCompleted)}
}

type CreateJobCommand struct {
	RequestUUID uuid.UUID
	Description string
}

type Job struct {
	UUID        uuid.UUID
	RequestUUID uuid.UUID
	State       JobState
	Description string
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (j *Job) Execute() {
	if j.State != JobStateCompleted {
		// simulate running job with processing time
		//nolint: gosec
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

		j.State = JobStateCompleted
		timeNow := time.Now()
		j.CompletedAt = &timeNow
	}
}

func ParseJobAPIResponse(job Job) api.JobResponse {
	return api.JobResponse{
		UUID:        job.UUID,
		RequestUUID: job.RequestUUID,
		State:       api.JobState(job.State),
		Description: job.Description,
		CompletedAt: job.CompletedAt,
		CreatedAt:   job.CreatedAt,
		UpdatedAt:   job.UpdatedAt,
	}
}

type NopMessageBus struct {
}

var _ MessageBus = (*NopMessageBus)(nil)

func (nop *NopMessageBus) PublishJob(
	ctx context.Context,
	job api.JobResponse,
) error {
	return nil
}

func (nop *NopMessageBus) SubscribeJob(
	ctx context.Context,
	queue string,
	do func(job api.JobResponse) error,
	states ...api.JobState,
) error {
	return nil
}
