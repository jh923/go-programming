package word

import (
	"fmt"
	"github.com/jh923/go-programming/onlineCourse/tests/12-Testing/02/quote"
	"testing"
)

func TestUseCount(t *testing.T) {
	x := UseCount("Testing how many unique words are counted")
	if len(x) != 7	{
		t.Error("Expected 7 got", len(x))
	}
}

func TestUseCount2(t *testing.T) {
	x := UseCount("Testing how many unique words are counted." +
		"Testing how many unique words are counted.")
	if len(x) != 7	{
		t.Error("Expected 7 got", len(x))
	}
}


// Could expand this to range over all words but is okay in this instance
func TestUseCount3(t *testing.T) {
	x := UseCount("Testing how many unique words are counted." +
		" Testing how many unique words are counted.")
	if x["Testing"] != 2	{
		t.Error("Expected 2 got", x["Testing"])
	}
}

func TestCount(t *testing.T) {
	x := Count("Hey there big fella")
	if x != 4	{
		t.Error("Expected 4 got", x)
	}
}

func ExampleCount() {
	fmt.Println(Count("There are seven words in this sentence."))
	// Output:
	// 7
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++	{
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++	{
		UseCount(quote.SunAlso)
	}
}