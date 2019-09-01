package main

import "testing"

func TestPower(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{1, 3}, 1},
		{[]int{2, 3}, 8},
		{[]int{3, 3}, 27},
	}

	for _, v := range tests {
		x := power(v.data[0], v.data[1])
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}
