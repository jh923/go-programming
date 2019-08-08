package main

import (
	"encoding/json"
	"fmt"
)

// Caps are now used otherwise the data cannot be accessed by the JSON martial
type Person struct {
	First string
	Last string
	Age int
}

func main()	{
	marshalExample()
	unmarshalExample()
}

func marshalExample()	{
	p1 := Person{
		First: "Tom",
		Last: "Rose",
		Age: 25,
	}
	p2 := Person{
		First: "Yuki",
		Last: "Nagato",
		Age: 200,
	}

	//JSON Marshal process
	people := []Person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil	{
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

func unmarshalExample()	{
	s := `[{"First":"Tom","Last":"Rose","Age":25},{"First":"Yuki","Last":"Nagato","Age":200}]`

	bs2 := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs2)

	var ES []Person

	err := json.Unmarshal(bs2, &ES)
	if err != nil	{
		fmt.Println(err)
	}
	for i,v := range ES	{
		fmt.Printf("Person number %d.\n", i)
		fmt.Printf("Name: %s %s. Age: %d", v.First, v.Last, v.Age)
	}
}