package main

import (
	"example.com/sender-client/queue"
	"example.com/sender-client/queue/manager"
	"fmt"
	"log"
)

func init() {
	fmt.Printf("Init")
}

func main() {
	fmt.Printf("Starting...")

	qm := manager.New(1)
	q, err := qm.BuildQueue(queue.RabbitMQ, "hello")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to build a queue", err)
	}

	body := "Hello World!"
	q.PublishPlainText([]byte(body))

	log.Printf(" [x] Sent %s\n", body)
}
