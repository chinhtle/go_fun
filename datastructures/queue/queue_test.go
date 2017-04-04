package main

import "testing"

func testEnqueue(q *Queue, t *testing.T, eHead, eTail interface{}, eLen int) {
	if q.head == nil || q.head.data != eHead {
		t.Error("Expected", eHead, "at head instead of", q.head.data)
	}

	if q.tail == nil || q.tail.data != eTail {
		t.Error("Expected", eTail, "at tail instead of", q.tail.data)
	}

	if q.Len() != eLen {
		t.Error("Expected len of", eLen, "but received", q.Len())
	}
}

func testDequeue(q *Queue, t *testing.T, eVal interface{}, eLen int) {
	if q.Dequeue() != eVal {
		t.Error("Did not receive expected value", eVal, "during dequeue")
	}

	if q.Len() != eLen {
		t.Error("Length of queue should be", eLen, "- Received", q.Len())
	}
}

func TestQueue_Len(t *testing.T) {
	q := NewQueue()

	if q.Enqueue(1); q.Len() != 1 {
		t.Error("Expecting length of 1, but got", q.Len())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1);
	testDequeue(q, t, 1, 0)

	q.Enqueue(1)
	q.Enqueue(2)
	testDequeue(q, t, 1, 1)

	q.Dequeue()
	testDequeue(q, t, nil, 0)
}

func TestQueue_Enqueue(t *testing.T) {
	q := NewQueue()

	headData := 1
	tailData := headData

	q.Enqueue(headData)
	testEnqueue(q, t, headData, tailData, 1)

	tailData = 2
	q.Enqueue(tailData)
	testEnqueue(q, t, headData, tailData, 2)

	tailData = 3
	q.Enqueue(tailData)
	testEnqueue(q, t, headData, tailData, 3)
}
