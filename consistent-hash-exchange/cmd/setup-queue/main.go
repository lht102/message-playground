package main

import (
	"flag"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	uri := flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	queueName := flag.String("queue", "test-queue", "Queue name")

	flag.Parse()

	if err := setup(*uri, *queueName); err != nil {
		log.Fatal("failed to setup queue", err)
	}
}

func setup(uri string, queueName string) error {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("channel: %w", err)
	}
	defer ch.Close()

	if _, err := ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	); err != nil {
		return fmt.Errorf("queue declare: %w", err)
	}

	if _, err := ch.QueuePurge(queueName, false); err != nil {
		return fmt.Errorf("queue purge: %w", err)
	}

	return nil
}
