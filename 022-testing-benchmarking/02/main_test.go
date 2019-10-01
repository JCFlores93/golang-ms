package main

import "testing"

// TestMySum test for mysum
func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 22}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 2}, 2},
		test{[]int{1, 0, -1}, 0},
	}
	for _, v := range tests {
		if x := mySum(v.data...); x != v.answer {
			t.Error("Expected", 5, "Got", x)
		}
	}
}
