package job

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	jobworker "github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/api"
	"github.com/lht102/message-playground/jobworker/ent"
	"go.uber.org/zap"
)

type Service struct {
	entClient  *ent.Client
	messageBus jobworker.MessageBus
	logger     *zap.Logger
}

var _ jobworker.JobService = (*Service)(nil)

func NewService(
	entClient *ent.Client,
	messageBus jobworker.MessageBus,
	logger *zap.Logger,
) *Service {
	return &Service{
		entClient:  entClient,
		messageBus: messageBus,
		logger:     logger,
	}
}

func (s *Service) CreateJob(
	ctx context.Context,
	createJobCmd jobworker.CreateJobCommand,
) (jobworker.Job, error) {
	job, err := s.createOrUpdateJob(ctx, func(ctx context.Context, jc *ent.JobClient) (*ent.Job, error) {
		job, err := jc.Create().
			SetRequestUUID(createJobCmd.RequestUUID).
			SetDescription(createJobCmd.Description).
			SetState(jobworker.JobStateQueued).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("create job: %w", err)
		}

		return job, nil
	})
	if err != nil {
		return jobworker.Job{}, err
	}

	return job, nil
}

func (s *Service) GetJob(
	ctx context.Context,
	uuid uuid.UUID,
) (jobworker.Job, error) {
	job, err := s.entClient.Job.Get(ctx, uuid)
	if err != nil {
		return jobworker.Job{}, fmt.Errorf("get job by id: %w", err)
	}

	return ent.ParseJobFromModel(job), nil
}

func (s *Service) ExecuteJob(
	ctx context.Context,
	uuid uuid.UUID,
) error {
	createdJob, err := s.GetJob(ctx, uuid)
	if err != nil {
		return fmt.Errorf("get job: %w", err)
	}

	if createdJob.State == jobworker.JobStateCompleted {
		return nil
	}

	_, err = s.createOrUpdateJob(ctx, func(ctx context.Context, jc *ent.JobClient) (*ent.Job, error) {
		job, err := jc.UpdateOneID(uuid).
			SetState(jobworker.JobStateProcessing).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("update job to processing state: %w", err)
		}

		return job, nil
	})
	if err != nil {
		return err
	}

	// doing some work
	createdJob.Execute()

	_, err = s.createOrUpdateJob(ctx, func(ctx context.Context, jc *ent.JobClient) (*ent.Job, error) {
		job, err := jc.UpdateOneID(createdJob.UUID).
			SetState(createdJob.State).
			SetCompletedAt(*createdJob.CompletedAt).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("update job to completed state: %w", err)
		}

		return job, nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RunBackgroundWorker(ctx context.Context) error {
	if err := s.messageBus.SubscribeJob(ctx, "jobworker_queue", func(resp api.JobResponse) error {
		if err := s.ExecuteJob(ctx, resp.UUID); err != nil {
			return fmt.Errorf("execute job: %w", err)
		}

		return nil
	}, api.JobStateQueued); err != nil {
		return fmt.Errorf("subscribe job: %w", err)
	}

	return nil
}

func (s *Service) createOrUpdateJob(
	ctx context.Context,
	do func(context.Context, *ent.JobClient) (*ent.Job, error),
) (jobworker.Job, error) {
	var job jobworker.Job

	if err := ent.WithTx(ctx, s.entClient, func(tx *ent.Tx) error {
		updatedJob, err := do(ctx, tx.Job)
		if err != nil {
			return err
		}

		job = ent.ParseJobFromModel(updatedJob)

		// FIXME: should not publish mesage if commit transaction fails
		if err := s.messageBus.PublishJob(ctx, jobworker.ParseJobAPIResponse(job)); err != nil {
			return fmt.Errorf("publish job: %w", err)
		}

		return nil
	}); err != nil {
		return jobworker.Job{}, fmt.Errorf("with tx: %w", err)
	}

	return job, nil
}
