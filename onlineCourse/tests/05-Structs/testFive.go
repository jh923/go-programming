/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/

package main

import "fmt"

func main() {
	test5()
}

func test5() {
	type person struct {
		first     string
		last      string
		favFlavor []string
	}

	p1 := person{
		first: "Jeff",
		last:  "Jones",
		favFlavor: []string{
			"Choc",
			"Strawberry",
			"Plain"},
	}

	p2 := person{
		first: "Nadir",
		last:  "Farfa",
		favFlavor: []string{
			"Mint",
			"Lemon",
			"Strawberry"},
	}

	//Done using for loop
	fmt.Printf("%s %s likes the following types of ice creams: ", p1.first, p1.last)
	for i := 0; i < len(p1.favFlavor); i++ {
		fmt.Print(p1.favFlavor[i], " ")
	}

	fmt.Printf("\n%s %s likes the following types of ice creams: ", p2.first, p2.last)
	for i := 0; i < len(p2.favFlavor); i++ {
		fmt.Print(p2.favFlavor[i], " ")
	}

	//Done using range
	fmt.Printf("\n%s %s likes the following types of ice creams: ", p1.first, p1.last)
	for _, v := range p1.favFlavor {
		fmt.Print(v, " ")
	}

	fmt.Printf("\n%s %s likes the following types of ice creams: ", p2.first, p2.last)
	for _, v := range p2.favFlavor {
		fmt.Print(v, " ")
	}

	fmt.Println()

	m1 := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m1 {
		fmt.Println("\n", k)
		for i, v2 := range v.favFlavor {
			fmt.Println(i, v2)
		}
	}

	type vehicle struct {
		doors  int
		colour string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedab struct {
		vehicle
		luxury bool
	}

	car1 := truck{
		vehicle: vehicle{
			doors:  2,
			colour: "blue",
		},
		fourWheel: true,
	}

	car2 := sedab{
		vehicle: vehicle{
			doors:  5,
			colour: "Red",
		},
		luxury: true,
	}

	fmt.Println(car1, car2)
	fmt.Println(car1.doors, car2.doors)

	fmt.Println()

	per := struct {
		first   string
		last    string
		age     int
		married bool
	}{
		first:   "Aivis",
		last:    "Jak",
		age:     21,
		married: false,
	}

	if per.married == true {
		fmt.Printf("%s %s is married.", per.first, per.last)
	} else {
		fmt.Printf("%s %s is not married.", per.first, per.last)
	}
}
