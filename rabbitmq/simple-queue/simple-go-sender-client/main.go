package main

import (
	connManager "example.com/sender-client/connection-manager"
	"example.com/sender-client/queue-manager/queue"
	"fmt"
	"log"
)

var q queue.Queue

func init() {
	fmt.Printf("Init")
	uri := fmt.Sprintf("%s:%s@%s:%s", "sender", "sender", "localhost", "5672")

	// TODO: new is glue, change for loose coupling
	conn := connManager.New("rabbitmq-1", uri, "vhost")
	q = queue.New("hello", conn)
}

func main() {
	fmt.Printf("Starting")
	body := "Hello World!"
	q.PublishPlainText([]byte(body))
	log.Printf(" [x] Sent %s\n", body)
}
