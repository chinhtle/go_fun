package main

import "fmt"

type element struct {
	value interface{}
	next  *element
}

// List represents a singly linked list
type List struct {
	root element
	len  int
}

// New returns a new list
func New() *List { return new(List) }

// AddValueAt inserts a value at an index
func (l *List) AddValueAt(v interface{}, index int) {
	l.Add(&element{v, nil}, index)
}

// Add inserts an element at an index
func (l *List) Add(e *element, index int) {
	curr := &l.root
	var i int

	for ; curr.next != nil && i < index; i++ {
		curr = curr.next
	}

	if i != index {
		fmt.Println("Invalid index.")
		return
	}

	e.next = curr.next
	curr.next = e
	l.len++
}

// FindValue looks for the value in the list and returns the index.
// If the element is not found then it returns -1
func (l *List) FindValue(v interface{}) int {
	for i, e := 0, l.root.next; e != nil; i, e = i+1, e.next {
		if e.value == v {
			return i
		}
	}

	return -1
}

// DeleteValue deletes an element in the list with the corresponding value.
// If the value was not found then it returns nil; otherwise, returns the
// element
func (l *List) DeleteValue(v interface{}) {
	if l.len < 1 {
		return
	}

	curr := l.root.next

	// Check if the value is at head
	if curr.value == v {
		l.root.next = curr.next
		curr.next = nil
		l.len--
		return
	}

	prev := curr
	curr = prev.next

	for ; curr != nil; curr = curr.next {
		if curr.value == v {
			break
		}
		prev = curr
	}

	// Indicates reached the end of list and value wasn't found
	if curr == nil {
		return
	}

	prev.next = curr.next
	curr.next = nil // Make sure to clear to avoid memory leaks
	l.len--
}

// Print displays all elements in the list
func (l *List) Print() {
	for curr := l.root.next; curr != nil; curr = curr.next {
		fmt.Print(curr.value, " ")
	}
	fmt.Println()
}

func (l *List) front() *element {
	if l.len < 1 {
		return nil
	}

	return l.root.next
}

// FrontValue returns the value of the first element; otherwise, returns nil
func (l *List) FrontValue() interface{} {
	if e := l.front(); e != nil {
		return e.value
	}

	return nil
}

// Back returns the last element in the list
func (l *List) back() *element {
	if l.len < 1 {
		return nil
	}

	curr := l.root.next
	for curr.next != nil {
		curr = curr.next
	}

	return curr
}

// BackValue returns the value of the last element; otherwise, returns nil
func (l *List) BackValue() interface{} {
	if e := l.back(); e != nil {
		return e.value
	}

	return nil
}

func main() {
	l := New()
	l.AddValueAt(1, 0)
	l.AddValueAt(2, 1)
	l.AddValueAt(3, 2)
	l.Print()

	fmt.Println("Front of the list:", l.FrontValue())
	fmt.Println("Back of the list:", l.BackValue())
	fmt.Println("Trying to delete non-existent value:")
	l.DeleteValue(4)
	l.Print()

	fmt.Println("Deleting existing value:")
	l.DeleteValue(2)
	l.Print()

	fmt.Println("Deleting existing value:")
	l.DeleteValue(1)
	l.Print()

	fmt.Println("Trying to delete non-existent value:")
	l.DeleteValue(5)
	l.Print()
}
