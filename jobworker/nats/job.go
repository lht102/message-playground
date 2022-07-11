package nats

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/api"
	"github.com/nats-io/nats.go"
)

const streamName = "job"

type MessageBus struct {
	nc   *nats.Conn
	js   nats.JetStreamContext
	subs []*nats.Subscription
}

var _ jobworker.MessageBus = (*MessageBus)(nil)

func NewMessageBus(
	url string,
	opts ...nats.Option,
) (*MessageBus, error) {
	nc, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, fmt.Errorf("connect nats server: %w", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, fmt.Errorf("init jet stream: %w", err)
	}

	return &MessageBus{
		nc: nc,
		js: js,
	}, nil
}

func (m *MessageBus) Close() error {
	for _, sub := range m.subs {
		if err := sub.Unsubscribe(); err != nil && !errors.Is(err, nats.ErrConnectionClosed) {
			return fmt.Errorf("unsubscribe: %w", err)
		}
	}

	return nil
}

func (m *MessageBus) PublishJob(
	ctx context.Context,
	job api.JobResponse,
) error {
	if err := m.createStreamIfNotExist(); err != nil {
		return err
	}

	b, err := json.Marshal(job)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	_, err = m.js.Publish(
		fmt.Sprintf("job.%s", strings.ToLower(string(job.State))),
		b,
	)
	if err != nil {
		return fmt.Errorf("publish: %w", err)
	}

	return nil
}

func (m *MessageBus) SubscribeJob(
	ctx context.Context,
	queue string,
	do func(job api.JobResponse) error,
	states ...api.JobState,
) error {
	if err := m.createStreamIfNotExist(); err != nil {
		return err
	}

	var filter string
	if len(states) == 0 || len(states) > 1 {
		filter = "*"
	} else {
		filter = strings.ToLower(string(states[0]))
	}

	filterMap := make(map[api.JobState]struct{})
	if len(states) == 1 {
		filterMap[states[0]] = struct{}{}
	} else {
		for _, state := range jobworker.JobStateQueued.Values() {
			filterMap[api.JobState(state)] = struct{}{}
		}
	}

	sub, err := m.js.QueueSubscribe(
		fmt.Sprintf("job.%s", filter),
		queue,
		func(msg *nats.Msg) {
			var resp api.JobResponse
			if err := json.Unmarshal(msg.Data, &resp); err != nil {
				_ = msg.Ack()

				return
			}

			if _, ok := filterMap[resp.State]; !ok {
				_ = msg.Ack()

				return
			}

			if err := do(resp); err != nil {
				_ = msg.Nak()

				return
			}

			_ = msg.Ack()
		},
	)
	if err != nil {
		return fmt.Errorf("create queue subscription: %w", err)
	}

	m.subs = append(m.subs, sub)

	return nil
}

func (m *MessageBus) createStreamIfNotExist() error {
	_, err := m.js.StreamInfo(streamName)
	if err != nil {
		if _, err := m.js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{fmt.Sprintf("%s.*", streamName)},
		}); err != nil {
			return fmt.Errorf("add stream: %w", err)
		}
	}

	return nil
}
