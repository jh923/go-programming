package main

import (
	"fmt"
	"sort"
)

func main()	{
	ints := []int{1,5,6,4,2,8,99,12,325,654,32,5}
	fmt.Println(ints)
	sort.Ints(ints)
	fmt.Println(ints)

	strings := []string{"Jeff", "Cat", "Dog", "Steel", "Zebra", "Bird", "Sea", "Twelve", "Rules", "For", "Life"}
	fmt.Println(strings)
	sort.Strings(strings)
	fmt.Println(strings)

}
