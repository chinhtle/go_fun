package main

import (
	"fmt"
	"sync"
)

type element struct {
	value interface{}
	next  *element
}

// Stack represents a LIFO data structure
type Stack struct {
	top  *element
	size int
	m    sync.Mutex
}

// NewStack returns an empty value of type Stack
func NewStack() *Stack { return &Stack{} }

// Push inserts a new value at the top of the stack
func (s *Stack) Push(v interface{}) {
	s.push(&element{v, nil})
}

func (s *Stack) push(e *element) {
	s.m.Lock()
	defer s.m.Unlock()

	e.next = s.top
	s.top = e
	s.size++
}

// Pop removes the top value from the stack
func (s *Stack) Pop() interface{} {
	return s.pop()
}

func (s *Stack) pop() (e *element) {
	s.m.Lock()
	defer s.m.Unlock()

	if s.size < 1 {
		return
	}

	e = s.top
	s.top = e.next
	s.size--
	return
}

// Print displays all values on the stack starting from the top
func (s *Stack) Print() {
	fmt.Println("Contents of stack:")

	s.m.Lock()
	defer s.m.Unlock()

	for top := s.top; top != nil; top = top.next {
		fmt.Println(top.value)
	}
}

// Top returns the value of the top-most item
func (s *Stack) Top() (v interface{}) {
	s.m.Lock()
	defer s.m.Unlock()

	if s.size < 1 {
		return
	}

	v = s.top.value
	return
}

// Len returns the number of elements in the stack
func (s *Stack) Len() int {
	s.m.Lock()
	defer s.m.Unlock()
	return s.size
}

func main() {
	s := NewStack()

	s.Push(1)
	s.Print()
	s.Push(2)
	s.Print()
	s.Push(3)
	s.Print()
	fmt.Println("Size of stack:", s.Len())

	fmt.Println("Top of the stack:", s.Top())
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
	fmt.Println("Top of the stack:", s.Top())

	fmt.Println("Attempting to pop from empty stack.")
	s.Pop()

	fmt.Println("Size of stack:", s.Len())
}
