package rabbitmq

import (
	"context"
	"example.com/sender-client/queue/rabbitmq/connection"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

// TODO: change *amqp.Channel to interface for mocking

type ConnectionManager interface {
	GetChannel() *amqp.Channel
	Shutdown()
}

// TODO: replace connManager by connection
// TODO: connection should be injected by ConnectionManager
// TODO: QueueConfig

type Queue struct {
	name        string
	connManager ConnectionManager
}

func New(name string) *Queue {
	uri := fmt.Sprintf("%s:%s@%s:%s", "sender", "sender", "localhost", "5672")
	conn := connection.New("rabbitmq-1", uri, "vhost")
	q := &Queue{
		name,
		conn,
	}
	q.declare()
	return q
}

func (q *Queue) declare() {
	// TODO: move elsewhere?
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

// TODO: verify if connection is not closed

func (q *Queue) PublishPlainText(body []byte) {
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

func (q *Queue) GetName() string {
	return q.name
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
