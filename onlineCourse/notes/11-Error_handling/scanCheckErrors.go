package main

import "fmt"

func main() {
	var first, last, from string
	var age int

	fmt.Println("Please enter your first name")
	_, err := fmt.Scan(&first)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter your last name")
	_, err = fmt.Scan(&last)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter your age")
	_, err = fmt.Scan(&age)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter where you are from")
	_, err = fmt.Scan(&from)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Hello %s %s age %d, from %s.\n", first, last, age, from)
}
