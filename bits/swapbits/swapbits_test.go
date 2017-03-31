package main

import "testing"

type test struct {
	num      int
	i, j     uint
	solution int
}

var testPairs = []test{
	{0, 0, 1, 0},
	{4, 0, 2, 1},
	{4, 1, 2, 2},
	{4, 2, 2, 4},
}

func TestSwapBits(t *testing.T) {
	for _, test := range testPairs {
		if v := swapBits(test.num, test.i, test.j); v != test.solution {
			t.Error("Expected", test.solution, "but received", v)
		}
	}
}
