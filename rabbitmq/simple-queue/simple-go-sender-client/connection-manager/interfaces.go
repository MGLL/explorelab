package connection_manager

import amqp "github.com/rabbitmq/amqp091-go"

// TODO: change *amqp.Channel to interface for mocking

type ConnectionManager interface {
	GetChannel() *amqp.Channel
	Shutdown()
}
