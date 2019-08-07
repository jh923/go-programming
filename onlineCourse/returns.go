package main

import "fmt"


func main() {
	defer goodbye()		//A deferred function will run when the final swe are going to exit the current function
	a := hello("Jeff", "Jones")
	fmt.Println(a)
	s, b := multiReturn("Jeff", 21)
	fmt.Println(s, b)
	sum1 := sum(1,2,3,4,5,6)
	fmt.Println(sum1)
	//Rather than passing a list of numbers to sum, we are passing a slice of numbers and using the ... to treat them
	// as individual ints rather than a slice of ints
	xi := []int{1,2,3,4,5,6,7,8,9}
	sum2 := sum(xi...)
	fmt.Println(sum2)
}

//This function has a list of parameters and then a return type, similar to Java
//We can also use the parameters directly rather than assigning variables to the parameters.
func hello(fn string, ln string) string {
	rt := fmt.Sprint("Hello ", fn, " ", ln)
	return rt
}

//Go can return more than one piece of data even data of varying types
func multiReturn(s string, n int) (string, bool){
	b := false

	if n%2 == 0	{
		b = false
	} else {
		b = true
	}

	return s, b
}

//Example of a variadic parameter
//takes an unlimited amount of ints and will store them in a new slice called x
// the variadic parameter must be the final parameter or the code will not run.
func sum( x ...int) int {
	sum := 0
	for  v := range x {
		sum += v
	}
	return sum
}


func goodbye()	{
	fmt.Printf("\nGoodbye")
}


