package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

//This function can only be used by secret agents
func (s secretAgent) speak() {
	fmt.Printf("I am agent %s %s.\n", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Jack",
			"Black",
		},
		ltk: true,
	}
	sa1.speak()
	sa2.speak()
}
