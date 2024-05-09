package manager

import (
	"fmt"
	"log"
	"testing"
)

type mockQueue struct {
	name string
}

func (m *mockQueue) PublishPlainText(bytes []byte) {
	log.Println(fmt.Sprintf("Sending bytes.... %s", string(bytes)))
}

func (m *mockQueue) GetName() string {
	return m.name
}

func TestCreateQueueManagerWithFiveCapacity(t *testing.T) {
	var CAPACITY uint8 = 5

	qm := New(CAPACITY)
	if qm == nil {
		t.Fatalf("Queue Manager instantiation failed")
	}

	qCapacity := qm.Capacity()
	if qCapacity != CAPACITY {
		t.Fatalf("Queue Manager capacity doesn't match, expected %d got %d", CAPACITY, qCapacity)
	}

	qLength := qm.Length()
	if qLength != 0 {
		t.Fatalf("Queue Manager length doesn't match, expected %d got %d", 0, qLength)
	}

	destroy()
}

func TestDoubleCreateQueueManager(t *testing.T) {
	qm1 := New(1)
	qm2 := New(2)

	if qm1 != qm2 {
		t.Fatalf("Queue managers are differents")
	}

	if qm2.Capacity() != 1 {
		t.Fatalf("Wrong capacity, expected %d got %d", 1, qm2.Capacity())
	}

	destroy()
}

func TestExceedQueueCapacityShouldReturnError(t *testing.T) {
	qm := New(1)
	mq := &mockQueue{"hello"}

	err := qm.addQueue(mq)
	if err != nil {
		t.Fatalf("First queue shouldn't return error: %s", err)
	}

	err = qm.addQueue(mq)
	if err == nil {
		t.Fatalf("Second queue should return an error")
	}

	destroy()
}

func TestGetMissingQueue(t *testing.T) {
	qm := New(1)
	q := qm.GetQueue("hello")
	if q != nil {
		t.Fatalf("Queue should be nil")
	}
	destroy()
}

func TestGeExistingQueue(t *testing.T) {
	qm := New(1)
	mq := &mockQueue{"hello"}
	_ = qm.addQueue(mq)
	q := qm.GetQueue("hello")
	if q == nil {
		t.Fatalf("Queue should not be nil")
	}
	destroy()
}
