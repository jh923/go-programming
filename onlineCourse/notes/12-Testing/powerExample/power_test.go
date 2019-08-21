package main

import "testing"

func TestPower(t *testing.T) {
	x := power(3,3)
	if x != 27	{
		t.Error("Expected 27 got", x)
	}
}