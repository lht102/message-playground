package rabbitmq

import (
	"context"

	"github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/api"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageBus struct {
	conn *amqp.Connection
}

var _ jobworker.MessageBus = (*MessageBus)(nil)

func (m *MessageBus) PublishJob(
	ctx context.Context,
	job api.JobResponse,
) error {
	panic("implement me")
}

func (m *MessageBus) SubscribeJob(
	ctx context.Context,
	queue string,
	do func(job api.JobResponse) error,
	states ...api.JobState,
) error {
	panic("implement me")
}
