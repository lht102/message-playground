package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	uri := flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchangeName := flag.String("exchange", "", "Exchange name")
	routingKey := flag.String("routing-key", "", "Routing key")
	hashHeader := flag.String("hash-header", "", "Hash header")
	numOfMsg := flag.Int("number-of-messages", 1, "Number of messages")

	flag.Parse()

	if err := publish(
		*uri,
		*exchangeName,
		*routingKey,
		*hashHeader,
		*numOfMsg,
	); err != nil {
		log.Fatal("failed to publish messages", err)
	}
}

func publish(
	uri string,
	exchangeName string,
	routingKey string,
	hashHeader string,
	numOfMsg int,
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

	if err := ch.Confirm(false); err != nil {
		return fmt.Errorf("confirm: %w", err)
	}

	publishCh := make(chan uint64, 1)
	confirmCh := ch.NotifyPublish(make(chan amqp.Confirmation, 1))

	go func() {
		for {
			select {
			case _ = <-publishCh:
			case confirmed := <-confirmCh:
				if confirmed.DeliveryTag > 0 {
					if !confirmed.Ack {
						log.Printf("failed delivery of delivery tag: %d\n", confirmed.DeliveryTag)
					}
				}
			}
		}
	}()

	arr := make([]int, 0, numOfMsg)
	for i := 1; i <= numOfMsg; i++ {
		arr = append(arr, i)
	}

	// shuffling
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1) //nolint: gosec
		arr[i], arr[j] = arr[j], arr[i]
	}

	for _, id := range arr {
		seqNo := ch.GetNextPublishSeqNo()

		table := amqp.Table{}
		if hashHeader != "" {
			table[hashHeader] = id
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := ch.PublishWithContext(
			ctx,
			exchangeName,
			routingKey,
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				Headers:      table,
				ContentType:  "text/plain",
				Body:         []byte(fmt.Sprintf("m%d", id)),
				DeliveryMode: amqp.Transient,
			},
		); err != nil {
			return fmt.Errorf("publish: %w", err)
		}

		publishCh <- seqNo
	}

	return nil
}
