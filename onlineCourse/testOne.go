/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/

package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main()	{
	test1()
}


func test1() {
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
	fmt.Printf("%v\n", x) //is 42 from prev exercise
	fmt.Printf("%T\n", x)
	x = 24
	fmt.Printf("%v\n", x)
	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)

}
