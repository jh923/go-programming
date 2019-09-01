package word

import (
	"strings"
)

// UseCount returns a map of all the words used and how many time each word was used in a string
// This code has a bug where punctuation turns a word into a different word.
// E.g. "There." and "There" are not the same word.
// The same applies to capitals such as: "There" and "there"
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns how many words are used in a string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
