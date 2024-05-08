package queue_manager

import (
	"example.com/sender-client/queue-manager/queue"
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type DefaultQueueManager struct {
	queues    []queue.Queue
	maxQueues uint8
}

var queueManager QueueManager

func New(maxQueues uint8) QueueManager {

	if queueManager == nil {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("Creating queue manager.")
		queueManager = &DefaultQueueManager{
			[]queue.Queue{},
			maxQueues,
		}
	}

	return queueManager
}

// TODO: finish

func (qm *DefaultQueueManager) BuildQueue(name string) queue.Queue {
	return nil
}

func (qm *DefaultQueueManager) GetQueue(name string) queue.Queue {
	return nil
}

func (qm *DefaultQueueManager) AddQueue(q queue.Queue) {
	_ = append(qm.queues, q)
}

func (qm *DefaultQueueManager) Capacity() uint8 {
	return qm.maxQueues
}

func (qm *DefaultQueueManager) Length() int {
	return len(qm.queues)
}
