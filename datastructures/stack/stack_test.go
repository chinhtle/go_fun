package main

import "testing"

func testPush(s *Stack, v interface{}, t *testing.T) {
	expectedLen := s.Len() + 1
	s.Push(v)

	if s.Top() != v || expectedLen != s.Len() {
		t.Error("Error ")
	}
}

func testPop(s *Stack, t *testing.T, eTop interface{}, eLen int) {
	if s.Pop(); s.Top() != eTop {
		t.Error("Top of the stack should be", eTop, ", got=", s.Top())
	} else if s.Len() != eLen {
		t.Error("Length of stack should be", eLen, ", got=", s.Len())
	}
}

func TestTop(t *testing.T) {
	s := NewStack()

	if s.Top() != nil {
		t.Error("Top value for empty stack should be nil")
	}

	v1 := 1

	if s.Push(v1); s.Top() != v1 {
		t.Error("Top value should be", v1, ", got=", s.Top())
	}
}

func TestStackPush(t *testing.T) {
	s := NewStack()

	testPush(s, 1, t)
	testPush(s, 2, t)
	testPush(s, 3, t)
}

func TestStackPopEmpty(t *testing.T) {
	s := NewStack()
	testPop(s, t, nil, 0)
}

func TestStackPopOneElement(t *testing.T) {
	s := NewStack()
	s.Push(1)
	testPop(s, t, nil, 0)
}

func TestStackPop(t *testing.T) {
	s := NewStack()

	v1, v2 := 1, 2
	s.Push(v1)
	s.Push(v2)

	testPop(s, t, v1, 1)
}
