package job

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	jobworker "github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/api"
	"github.com/lht102/message-playground/jobworker/ent"
	"github.com/lht102/message-playground/jobworker/ent/job"
	"go.uber.org/zap"
)

type Service struct {
	entClient  *ent.Client
	messageBus jobworker.MessageBus
	logger     *zap.Logger
}

var _ jobworker.Service = (*Service)(nil)

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
	job, err := s.entClient.Job.Query().
		Where(job.IDEQ(uuid)).
		Only(ctx)
	if err != nil {
		return jobworker.Job{}, fmt.Errorf("get job by id: %w", err)
	}

	return ent.ParseJobFromModel(job), nil
}

func (s *Service) RunBackgroundWorker(ctx context.Context) error {
	if err := s.messageBus.SubscribeJob(ctx, "jobworker_queue", func(resp api.JobResponse) error {
		job, err := s.createOrUpdateJob(ctx, func(ctx context.Context, jc *ent.JobClient) (*ent.Job, error) {
			job, err := jc.UpdateOneID(resp.UUID).
				SetState(jobworker.JobStateProcessing).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("update job to processing state: %w", err)
			}

			return job, nil
		})
		if err != nil {
			s.logger.Error("Failed to update job to processing state", zap.Error(err))

			return err
		}

		// doing some work
		job.Execute()

		_, err = s.createOrUpdateJob(ctx, func(ctx context.Context, jc *ent.JobClient) (*ent.Job, error) {
			job, err := jc.UpdateOneID(job.UUID).
				SetState(job.State).
				SetCompletedAt(*job.CompletedAt).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("update job to completed state: %w", err)
			}

			return job, nil
		})
		if err != nil {
			s.logger.Error("Failed to update job to completed state", zap.Error(err))

			return err
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
