package subscription

import (
	"context"
	"fmt"

	"github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/api"
)

type Worker struct {
	queueName  string
	jobService jobworker.JobService
	consumer   jobworker.Consumer
}

func NewWorker(
	queueName string,
	jobService jobworker.JobService,
	consumer jobworker.Consumer,
) *Worker {
	return &Worker{
		queueName:  queueName,
		jobService: jobService,
		consumer:   consumer,
	}
}

func (s *Worker) Run(ctx context.Context) error {
	if err := s.consumer.SubscribeJob(ctx, s.queueName, func(resp api.JobResponse) error {
		if err := s.jobService.ExecuteJob(ctx, resp.UUID); err != nil {
			return fmt.Errorf("execute job: %w", err)
		}

		return nil
	}, api.JobStateQueued); err != nil {
		return fmt.Errorf("subscribe job: %w", err)
	}

	return nil
}
