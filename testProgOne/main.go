/**
	This document is used to host my notes on a range of topics about go
 */

package main

import "fmt"

var a string = "global scoped var"	//var name type. Type is not needed as the compiler can figure it out. But locks the scope of the variable

func main()  {
	foo()
	convert()
}

/**
	foo contains basic info about formatting and data types and printing
 */
func foo()	{
	var y = "within the scope of function main" //can be declared outside of function body.
	x := 10	//Can't be outside of a function body

	fmt.Println("printing")
	fmt.Println(x*x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Printf("%T\n", a)	//print format type of a, then a new line
	fmt.Printf("%T\t%b\t%#x\n", x, x, x)	//print format type of a, tab, then binary value tab, hex value, new lime
	fmt.Println(`"Allows you to print quote marks"`)	// reference to the ` symbol

	str := fmt.Sprintf("%T\t%b\t%#x\n", x, x, x)	//Sprint converts to a string
	fmt.Println(str)

	//Fprint is used for printing to a file
}

/**
	Making a new type and converting
 */
func convert()	{
	type sInt int
	var num sInt
	num = 100
	fmt.Printf("%T",num) //num is of type sInt not int and thus cannot be converted to an int


}