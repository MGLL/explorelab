package queue_manager

import "testing"

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
}
