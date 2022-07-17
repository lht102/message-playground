package job_test

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lht102/message-playground/jobworker"
)

func (s *JobTestSuite) TestCreateJob() {
	ctx := context.TODO()
	requestUUID := uuid.MustParse("c0be5e3f-4e1d-4d15-b427-8f21a1fe5e9c")
	description := "Some Text"
	job, err := s.jobService.CreateJob(ctx, jobworker.CreateJobCommand{
		RequestUUID: requestUUID,
		Description: description,
	})
	s.NoError(err)
	s.Equal(jobworker.Job{
		UUID:        job.UUID,
		RequestUUID: requestUUID,
		State:       jobworker.JobStateQueued,
		Description: description,
		CreatedAt:   job.CreatedAt,
		UpdatedAt:   job.UpdatedAt,
	}, job)
}

func (s *JobTestSuite) TestGetJob() {
	ctx := context.TODO()
	createdJob, err := s.entClient.Job.
		Create().
		SetID(uuid.MustParse("f07e26c1-9a4a-4721-89e3-e7793bfff33f")).
		SetRequestUUID(uuid.MustParse("42cfd3f8-dac0-4be0-a15b-4e60dc8964e4")).
		SetState(jobworker.JobStateCompleted).
		SetDescription("todo").
		SetCompletedAt(time.Date(2022, 1, 1, 3, 0, 0, 0, time.UTC)).
		SetCreatedAt(time.Date(2022, 1, 1, 1, 0, 0, 0, time.UTC)).
		SetUpdatedAt(time.Date(2022, 1, 1, 3, 0, 0, 0, time.UTC)).
		Save(ctx)
	s.NoError(err)
	job, err := s.jobService.GetJob(ctx, createdJob.ID)
	s.NoError(err)
	s.Equal(jobworker.Job{
		UUID:        createdJob.ID,
		RequestUUID: createdJob.RequestUUID,
		State:       createdJob.State,
		Description: createdJob.Description,
		CompletedAt: createdJob.CompletedAt,
		CreatedAt:   createdJob.CreatedAt,
		UpdatedAt:   createdJob.UpdatedAt,
	}, job)
}
