package main

import "fmt"

func main()  {
	structExamples()
}

func structExamples() {
	type person struct {
		first string
		last string
	}

	p1 := person{
		first: "Callum",
		last: "Russel",
	}

	p2 := person{
		first: "James",
		last: "Welsh",
	}

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2)
	fmt.Println(p2.first, p2.last)


}
