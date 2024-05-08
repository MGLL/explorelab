package queue

import (
	"context"
	"example.com/sender-client/connection-manager"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type DefaultQueue struct {
	name        string
	connManager connection_manager.ConnectionManager
}

func New(name string, connManager connection_manager.ConnectionManager) *DefaultQueue {
	q := &DefaultQueue{
		name,
		connManager,
	}
	q.declare()
	return q
}

func (q *DefaultQueue) declare() {
	ch := q.connManager.GetChannel()

	_, err := ch.QueueDeclare(
		q.name,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare Queue")
}

func (q *DefaultQueue) PublishPlainText(body []byte) {
	ch := q.connManager.GetChannel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ch.PublishWithContext(ctx,
		"",
		q.name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
