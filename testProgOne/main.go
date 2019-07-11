/**
This document is used to host my notes on a range of topics about go
*/

package main

import "fmt"

var a string = "global scoped var" //var name type. Type is not needed as the compiler can figure it out. But locks the scope of the variable

func main() {
	//intro()
	//convert()
	//test1()
	//con()
	test2()
}

/**
foo contains basic info about formatting and data types and printing
*/
func intro() {
	var y = "within the scope of function main" //can be declared outside of function body.
	x := 10                                     //Can't be outside of a function body

	fmt.Println("printing")
	fmt.Println(x * x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Printf("%T\n", a)                            //print format type of a, then a new line
	fmt.Printf("%T\t%b\t%#x\n", x, x, x)             //print format type of a, tab, then binary value tab, hex value, new lime
	fmt.Println(`"Allows you to print quote marks"`) // reference to the ` symbol

	str := fmt.Sprintf("%T\t%b\t%#x\n", x, x, x) //Sprint converts to a string
	fmt.Println(str)

	//Fprint is used for printing to a file
}

/**
Making a new type and converting
*/
func convert() {
	type sInt int
	var num sInt
	num = 100
	fmt.Printf("%T\n", num) //num is of type sInt not int and thus cannot be converted to an int
	a := 42
	a = int(num)	//int(value) converts a value to an int, in the from T(x) where T is the desired type and x is a value which can be converted to type T
	fmt.Println(a)

}

var x int = 42
var y string = "James Bond"
var z bool = true
/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
 */
func test1()	{
	fmt.Println("Test 1.1 - 1.3")
	//x := 42
	//y := "James Bond"
	//z := true
	fmt.Println(x, y, z)
	ansT := fmt.Sprintf("%T\t%T\t%T", x, y, z)
	fmt.Println(ansT)
	ansV := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(ansV)

	fmt.Println("\nTest 1.4")
	type epicInt int
	var x epicInt
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

	fmt.Println("\nTest 1.5")

	var y int
	fmt.Printf("%v\n", x)	//is 42 from prev exercise
	fmt.Printf("%T\n", x)
	x = 24
	fmt.Printf("%v\n", x)
	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)


}

/*
defining constants
Constants have their expected type and value are are immutable
 */
func con() {
	//Different ways of defining values as constants
	const (
		jeff= "my name is"
		notTrue= false
	)

	const scrub bool = true

	fmt.Printf("%T\t%v\n", jeff, jeff)
	fmt.Printf("%T\t%v\n", notTrue, notTrue)
	fmt.Printf("%T\t%v\n", scrub, scrub)

	fmt.Println(jeff)
	fmt.Println(notTrue)
	fmt.Println(scrub)

	//iota - used to count up in constants
	const  (
		a0 = iota
		a1
		a2
	)
	fmt.Println(a0, a1, a2)

	//using iota in bit shift to denote memory

	const (
		_ = iota	// _ discard char
		kb = 1 << (iota * 10) // 1 shift left 10
		mb = 1 << (iota * 10) // 1 shift left 20
		gb = 1 << (iota * 10) // 1 shift left 30
	)
	fmt.Println(kb, mb, gb)
	fmt.Printf("%b\n", kb)
	fmt.Printf("%b\n", mb)
	fmt.Printf("%b\n", gb)
}


func test2() {
	//test 2.1 - print a number in base 10, base 2 and base 16
	num := 100
	fmt.Printf("%d\t%b\t%X\n", num, num, num)

	//test 2.2 - comparing values
	a := 10
	b := 100
	g := a == b
	h := a <= b
	j := a != b
	k := a <  b
	l := a >  b
	fmt.Println(g,h,j,k,l)

	//test 2.3 - constants
	const t23 = 23
	const t231 int = 23
	fmt.Printf("%v\t%v\n", t23, t231)

	//test 2.4 - bit shift
	a = 7
	fmt.Printf("%d\t%b\t%X\n", a, a, a)
	a = a << 1
	fmt.Printf("%d\t%b\t%X\n", a, a, a)

	//test 2.5 - raw string literal
	c := `Literal string`
	fmt.Println(c)

	//test 2.6 - iota
	const  (
		_  = 2019 + iota
		n1 = 2019 + iota
		n2 = 2019 + iota
		n3 = 2019 + iota
		n4 = 2019 + iota
	)
	fmt.Println(n1,n2,n3,n4)
}