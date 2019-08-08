/*
defining constants
Constants have their expected type and value are are immutable
*/

package main

import "fmt"

func main() {
	con()
}

func con() {
	//Different ways of defining values as constants
	const (
		jeff    = "my name is"
		notTrue = false
	)

	const scrub bool = true

	fmt.Printf("%T\t%v\n", jeff, jeff)
	fmt.Printf("%T\t%v\n", notTrue, notTrue)
	fmt.Printf("%T\t%v\n", scrub, scrub)

	fmt.Println(jeff)
	fmt.Println(notTrue)
	fmt.Println(scrub)

	//iota - used to count up in constants
	const (
		a0 = iota
		a1
		a2
	)
	fmt.Println(a0, a1, a2)

	//using iota in bit shift to denote memory

	const (
		_  = iota             // _ discard char
		kb = 1 << (iota * 10) // 1 shift left 10
		mb = 1 << (iota * 10) // 1 shift left 20
		gb = 1 << (iota * 10) // 1 shift left 30
	)
	fmt.Println(kb, mb, gb)
	fmt.Printf("%b\n", kb)
	fmt.Printf("%b\n", mb)
	fmt.Printf("%b\n", gb)
}
