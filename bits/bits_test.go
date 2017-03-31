package main

import "testing"
import "math"

type tests struct {
	num, solution int
}

var testPairs = []tests{
	{0, 0},
	{1, 1},
	{5, 2},
	{math.MaxInt32, 31},
	{math.MaxUint32, 32},
	{math.MaxInt64, 63},
}

func TestCountBits(t *testing.T) {
	for _, v := range testPairs {
		if countBits(v.num) != v.solution {
			t.Error("Expected", v.solution, "got", v.num)
		}
	}
}
