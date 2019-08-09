package main

import (
"fmt"
"sort"
)

type Person struct {
	First string
	Age   int
}

// We define a new type which impediments the sort interface sao it requires these three functions: Len, Swap and Less
//This example is for sorting by ages
type ByAge []Person
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

//This is another example of implementing the sort interface, but sorts by name
type ByFirst []Person
func (f ByFirst) Len() int           { return len(f) }
func (f ByFirst) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ByFirst) Less(i, j int) bool { return f[i].First < f[j].First }

func main() {
	p1 := Person{"Donald", 32}
	p2 := Person{"Alex", 27}
	p3 := Person{"Callum", 64}
	p4 := Person{"Oly", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)

	//sorting an interface doesn't return anything, it orders the struct. so no assignment is needed
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByFirst(people))
	fmt.Println(people)
}