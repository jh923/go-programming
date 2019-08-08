/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/

package main

import "fmt"

func main()	{
	//test 7.11
	x := 42
	fmt.Println(&x)

	//test 7.2
	p1 := t7Person{
		first: "Jack",
		last: "Jones",
	}
	
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}


func changeMe(p *t7Person) {
	p.first = "Jeff"
}

type t7Person struct {
	first string
	last string

}