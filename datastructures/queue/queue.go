package main

import (
	"fmt"
	"sync"
)

type element struct {
	data interface{}
	next *element
}

// Queue represents a FIFO data structure
type Queue struct {
	head, tail *element
	count      int
	m          sync.Mutex
}

// NewQueue returns a new queue
func NewQueue() *Queue { return &Queue{} }

// Enqueue appends a new item to the queue
func (q *Queue) Enqueue(v interface{}) {
	e := &element{v, nil}

	q.m.Lock()
	defer q.m.Unlock()

	if q.head == nil {
		q.head = e
	} else {
		q.tail.next = e
	}

	q.tail = e
	q.count++
}

// Dequeue removes the first item in the queue
func (q *Queue) Dequeue() (v interface{}) {
	q.m.Lock()
	defer q.m.Unlock()

	if q.count < 1 {
		return
	}

	head := q.head
	v = head.data
	q.head = head.next
	head.next = nil

	if q.count--; q.count == 0 {
		q.tail = nil
	}

	return
}

// Len returns the number of items in the queue
func (q *Queue) Len() int {
	q.m.Lock()
	defer q.m.Unlock()
	return q.count
}

// Print displays all data in the queue
func (q *Queue) Print() {
	fmt.Println("Contents in the queue:")
	for n := q.head; n != nil; n = n.next {
		fmt.Print(n.data, " ")
	}
	fmt.Println()
}

func main() {
	q := NewQueue()

	q.Enqueue(1)
	q.Print()
	q.Enqueue(2)
	q.Print()
	q.Enqueue(3)
	q.Print()
	fmt.Println("Size of the queue:", q.Len())

	q.Dequeue()
	q.Print()
	q.Dequeue()
	q.Print()
	q.Dequeue()
	q.Print()
	fmt.Println("Size of the queue:", q.Len())

	fmt.Println("Attempting to pop from empty queue..")
	q.Dequeue()
	fmt.Println("Size of the queue:", q.Len())
}
