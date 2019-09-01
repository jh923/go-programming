package main

import (
	"fmt"
	"github.com/jh923/go-programming/onlineCourse/tests/12-Testing/02/quote"
	"github.com/jh923/go-programming/onlineCourse/tests/12-Testing/02/word"
)

func main() {
	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
	fmt.Println("Is a list of every word in the quote")
	fmt.Println("There are", word.Count(quote.SunAlso), "words in this quote")
	fmt.Println("There are", len(word.UseCount(quote.SunAlso)), "unique words in this quote")

}
