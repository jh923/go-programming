/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/

package main

import "fmt"

func main()  {
	test3()
}

func test3() {
	//print 1<=x<=10
	for i := 1; i < +10; i++ {
		fmt.Println(i)
	}

	fmt.Println()
	fmt.Println()

	//print utf-8 capital english letters 3 times each and their respective number
	for i := 65; i < 91; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	//loop until age is correctly guessed

	age := 21
	guess := 0
	for {
		if guess == age {
			fmt.Printf("Your ages is: %d", guess)
			break
		}
		guess++
	}

	//find x mod 4 where 10<=x<=25
	for i := 10; i <= 25; i++ {
		fmt.Println(i % 4)
	}

	fmt.Println()
	fmt.Println()

	if age < 18 {
		fmt.Println("You can't drink in the UK or US")
	} else if age >= 18 && age < 21 {
		fmt.Println("You can drink in the UK but not the US")
	} else {
		fmt.Println("You can drink in the UK and the US")
	}

	fmt.Println()
	fmt.Println()

	favSport := "Golf"

	switch favSport {
	case "DMM", "LCS", "Basketball", "Rugby":
		fmt.Printf("I like %v too!", favSport)

	case "Football", "Golf", "Baseball":
		fmt.Printf("I don't like %v", favSport)
	default:
		fmt.Println("I don't know that sport")
	}

	fmt.Println()
	fmt.Println()

	//true, false, true, false, false

}
