package connection_manager

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

// TODO: Try to wrap amqp for unit testing (exploration)

type DefaultConnectionManager struct {
	name  string
	uri   string
	vhost string
	conn  *amqp.Connection
	ch    *amqp.Channel
}

func New(name, uri, vhost string) ConnectionManager {
	return &DefaultConnectionManager{
		name:  name,
		uri:   uri,
		vhost: vhost,
	}
}

func (cm *DefaultConnectionManager) GetChannel() *amqp.Channel {
	if cm.ch == nil {
		conn := cm.getConnection()
		ch, err := conn.Channel()
		failOnError(err, "Failed to open RabbitMQ channel")
		cm.ch = ch
	}

	return cm.ch
}

func (cm *DefaultConnectionManager) getConnection() *amqp.Connection {
	if cm.conn == nil {
		conn, err := amqp.Dial(
			fmt.Sprintf("amqp://%s/%s", cm.uri, cm.vhost))
		failOnError(err, "Failed to connect to RabbitMQ")
		cm.conn = conn
	}

	return cm.conn
}

func (cm *DefaultConnectionManager) Shutdown() {
	cm.closeChannel()
	cm.closeConnection()
}

func (cm *DefaultConnectionManager) closeChannel() {
	err := cm.ch.Close()
	failOnError(err, "Failed to close RabbitMQ channel")
}

func (cm *DefaultConnectionManager) closeConnection() {
	err := cm.conn.Close()
	failOnError(err, "Failed to close to RabbitMQ connection")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
