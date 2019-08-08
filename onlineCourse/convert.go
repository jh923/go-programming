/**
Making a new type and converting
*/

package main

import "fmt"

func main() {
	convert()
}

func convert() {
	type sInt int
	var num sInt
	num = 100
	fmt.Printf("%T\n", num) //num is of type sInt not int and thus cannot be converted to an int
	a := 42
	a = int(num) //int(value) converts a value to an int, in the from T(x) where T is the desired type and x is a value which can be converted to type T
	fmt.Println(a)

}
