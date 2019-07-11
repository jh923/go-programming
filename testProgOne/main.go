package main

import "fmt"

var a string = "global scoped var"	//var name type. Type is not needed as the compiler can figure it out. But locks the scope of the variable

func main()  {
	fmt.Println("printing")

	x := 10	//Can't be outside of a function body
	fmt.Println(x*x)

	var y = "within the scope of function main" //can be declared outside of function body.
	fmt.Println(y)
}

func foo()	{
	fmt.Println(a)
}