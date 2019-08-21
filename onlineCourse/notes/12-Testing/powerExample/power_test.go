package main

import (
	"fmt"
	"testing"
)

func TestPower(t *testing.T) {
	x := power(3, 3)
	if x != 27 {
		t.Error("Expected 27 got", x)
	}
}

func ExamplePower() {
	fmt.Println(power(3, 3))
	// Output:
	// 27
}

func BenchmarkPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		power(3, 3)
	}
}
