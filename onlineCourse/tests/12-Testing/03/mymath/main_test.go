package mymath

import (
	"fmt"
	"testing"
)

// This could be changed to use a structure to go over a range of data, answer pairs
func TestCenteredAvg(t *testing.T) {
	xs := []int{1, 2, 3}
	if len(xs) <= 2 {
		t.Error("Your list doesn't have enough elements")
	}
	avg := CenteredAvg(xs)
	if avg != 2 {
		t.Error("Expected 2 got", avg)
	}
}

func ExampleCenteredAvg() {
	xs := []int{1, 2, 3}
	fmt.Println(CenteredAvg(xs))
	// Output:
	// 2
}

func BenchmarkCenteredAvg(b *testing.B) {
	xs := []int{1, 2, 3}
	if len(xs) <= 2 {
		b.Error("Your list doesn't have enough elements")
	}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xs)
	}
}
