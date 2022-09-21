package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	uri := flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	queueName := flag.String("queue", "", "Queue name")
	consumerTag := flag.String("consumer-tag", "", "Consumer tag")

	flag.Parse()

	if err := consume(
		*uri,
		*queueName,
		*consumerTag,
	); err != nil {
		log.Fatal("failed to consume messages", err)
	}
}

func consume(
	uri string,
	queueName string,
	consumerTag string,
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

	if err := ch.Qos(1, 0, false); err != nil {
		return fmt.Errorf("qos: %w", err)
	}

	delivery, err := ch.Consume(
		queueName,
		consumerTag,
		false, // noAck
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("consume: %w", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c

		_ = ch.Cancel(consumerTag, true)
	}()

	freq := map[string]int{}

	for d := range delivery {
		message := string(d.Body)
		freq[message]++
		fmt.Printf("%s = %d\n", message, freq[message])

		_ = d.Ack(false)
	}

	return nil
}
