package manager

import (
	"errors"
	"example.com/sender-client/queue"
	"example.com/sender-client/queue/rabbitmq"
	"fmt"
	"slices"
	"sync"
)

type DefaultQueueManager struct {
	queues    []Queue
	maxQueues uint8
}

var lock = &sync.Mutex{}
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

func destroy() {
	lock.Lock()
	defer lock.Unlock()
	queueManager.queues = nil
	queueManager = nil
}

func (qm *DefaultQueueManager) BuildQueue(queueType queue.Type, name string) (Queue, error) {
	if qm.hasAvailableCapacity() {
		return nil, errors.New("manager is at max capacity")
	}
	q := qm.GetQueue(name)
	if q != nil {
		return q, nil
	}

	if queueType == queue.RabbitMQ {
		q = rabbitmq.New(name)
		_ = qm.addQueue(q)
		return q, nil
	}

	return nil, fmt.Errorf("wrong queue type passed")
}

func (qm *DefaultQueueManager) GetQueue(name string) Queue {
	idx := slices.IndexFunc(qm.queues, func(q Queue) bool { return q.GetName() == name })
	if idx == -1 {
		return nil
	}
	return qm.queues[idx]
}

func (qm *DefaultQueueManager) addQueue(q Queue) error {
	if qm.hasAvailableCapacity() {
		return errors.New("manager is at max capacity")
	}
	qm.queues = append(qm.queues, q)
	return nil
}

func (qm *DefaultQueueManager) hasAvailableCapacity() bool {
	return qm.Capacity() == qm.Length()
}

func (qm *DefaultQueueManager) Capacity() uint8 {
	return qm.maxQueues
}

func (qm *DefaultQueueManager) Length() uint8 {
	return uint8(len(qm.queues))
}
