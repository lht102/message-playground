package main

import (
	"flag"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	uri := flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchangeName := flag.String("exchange", "test-exchange", "Exchange name")
	queuePrefix := flag.String("queue-prefix", "test-queue", "Queue prefix")
	hashHeader := flag.String("hash-header", "", "Hash header")
	numOfQueue := flag.Int("number-of-queues", 1, "Number of queues")

	flag.Parse()

	if err := setup(
		*uri,
		*exchangeName,
		*queuePrefix,
		*numOfQueue,
		*hashHeader,
	); err != nil {
		log.Fatal("failed to setup consistent hash exchange", err)
	}
}

func setup(
	uri string,
	exchangeName string,
	queuePrefix string,
	numOfQueue int,
	hashHeader string,
) error {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("channel: %w", err)
	}
	defer ch.Close()

	table := amqp.Table{}
	if hashHeader != "" {
		table["hash-header"] = hashHeader
	}

	if err := ch.ExchangeDeclare(
		exchangeName,
		"x-consistent-hash",
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // noWait
		table,
	); err != nil {
		return fmt.Errorf("exchange declare: %w", err)
	}

	for i := 1; i <= numOfQueue; i++ {
		queueName := fmt.Sprintf("%s%03d", queuePrefix, i)

		if _, err := ch.QueueDeclare(
			queueName,
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // noWait
			amqp.Table{
				"x-single-active-consumer": true,
			},
		); err != nil {
			return fmt.Errorf("queue declare: %w", err)
		}

		if err := ch.QueueBind(
			queueName,
			"1", // key
			exchangeName,
			false, // noWait
			nil,
		); err != nil {
			return fmt.Errorf("queue bind: %w", err)
		}

		if _, err := ch.QueuePurge(queueName, false); err != nil {
			return fmt.Errorf("queue purge: %w", err)
		}
	}

	return nil
}
