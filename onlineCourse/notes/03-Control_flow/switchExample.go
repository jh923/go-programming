package main

import "fmt"

func main() {
	switchExample()
}

func switchExample() {
	//Switch by default has no fall though. Meaning if multiple statements are true only the first true one will be executed.
	switch {
	case (2%2 == 1):
		fmt.Println("This will not print ever")
	case (2 == 2):
		fmt.Println("this will print")
	case (true == true):
		fmt.Println("although this is true this statement is unreachable.")
	}

	//reason toi not use fall though
	switch {
	case (2%2 == 1):
		fmt.Println("This will not print ever")
		fallthrough
	case (2 == 2):
		fmt.Println("this will print")
		fallthrough
	case (true == true):
		fmt.Println("This statement is reached due to fall thought")
		fallthrough
	case (false == true):
		fmt.Println("Due to fallthrough this statement is printed although it is false")
	default:
		fmt.Println("Default")
	}

}
