package queue_manager

import "example.com/sender-client/queue-manager/queue"

type QueueManager interface {
	BuildQueue(name string) queue.Queue
	GetQueue(name string) queue.Queue
	AddQueue(queue.Queue)
	Capacity() uint8
	Length() int
}
