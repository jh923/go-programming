package main

import "fmt"

func main() 	{
	f := func() {
		fmt.Println("This function is assigned to a variable. ")
	}
	f()

	gorwell := orwell()
	fmt.Printf("%T\n", gorwell())
	fmt.Println(gorwell())
	fmt.Println(bar()())
}

func orwell() func() int{
	return func() int {
		return 1984
	}
}

func bar() func() int {
	return func() int {
		return 451
	}
}