package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak()	{
	fmt.Printf("I am agent %s %s.\n", s.first, s.last)
}

func (p person) speak()	{
	fmt.Printf("I am %s %s.\n", p.first, p.last)
}

func hum(h human){
	switch h.(type) {
	case person:
		fmt.Printf("This is an example of asserting %s is an person", h.(person).first)
	case secretAgent:
		fmt.Printf("This is an example of asserting %s is an person", h.(secretAgent).first)
	}

}

func main()	{
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	p1 := person{
		first: "Mark",
		last: "Zucc",
	}
	sa1.speak()
	sa1.person.speak()
	hum(sa1)
	hum(p1)

}