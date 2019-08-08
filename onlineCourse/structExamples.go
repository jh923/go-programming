package main

import "fmt"

func main() {
	structExamples()
}

func structExamples() {
	// A defined struct
	type person struct {
		first string
		last  string
	}
	// A struct with inheritance
	type programmer struct {
		person
		langs []string
	}

	p1 := programmer{
		person: person{
			first: "Callum",
			last:  "Russel",
		},
		langs: []string{"Java", "PHP", "HTML", "CSS"},
	}

	p2 := person{
		first: "James",
		last:  "Welsh",
	}

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, "knows: ", p1.langs)
	fmt.Println(p2)
	fmt.Println(p2.first, p2.last)

	//A anon struct
	p3 := struct {
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

	fmt.Println(p3)

}
