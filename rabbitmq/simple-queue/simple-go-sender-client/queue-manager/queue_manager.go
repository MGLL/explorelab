package queue_manager

import (
	"fmt"
	"sync"
)

type Queue interface {
	PublishPlainText([]byte)
}

var lock = &sync.Mutex{}

type DefaultQueueManager struct {
	queues    []Queue
	maxQueues uint8
}

var queueManager *DefaultQueueManager

func New(maxQueues uint8) *DefaultQueueManager {

	if queueManager == nil {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("Creating queue manager.")
		queueManager = &DefaultQueueManager{
			[]Queue{},
			maxQueues,
		}
	}

	return queueManager
}

// TODO: finish

func (qm *DefaultQueueManager) BuildQueue(name string) Queue {
	return nil
}

func (qm *DefaultQueueManager) GetQueue(name string) Queue {
	return nil
}

func (qm *DefaultQueueManager) AddQueue(q Queue) {
	_ = append(qm.queues, q)
}

func (qm *DefaultQueueManager) Capacity() uint8 {
	return qm.maxQueues
}

func (qm *DefaultQueueManager) Length() int {
	return len(qm.queues)
}
