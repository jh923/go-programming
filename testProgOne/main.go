/**
This document is used to host my notes on a range of topics about go
*/

package main

import "fmt"

var a string = "global scoped var" //var name type. Type is not needed as the compiler can figure it out. But locks the scope of the variable

func main() {
	//foo()
	//convert()
	test1()
}

/**
foo contains basic info about formatting and data types and printing
*/
func foo() {
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
