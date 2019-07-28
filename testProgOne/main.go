/**
This document is used to host my notes on a range of topics about go
*/

package main

import "fmt"

var a string = "global scoped var" //var name type. Type is not needed as the compiler can figure it out. But locks the scope of the variable

func main() {
	//intro()
	//convert()
	//test1()
	//con()
	//test2()
	//switchExample()
	//test3()
	//groupData()
	test4()
}

/**
foo contains basic info about formatting and data types and printing
*/
func intro() {
	var y = "within the scope of function main" //can be declared outside of function body.
	x := 10                                     //Can't be outside of a function body

	fmt.Println("printing")
	fmt.Println(x * x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Printf("%T\n", a)                            //print format type of a, then a new line
	fmt.Printf("%T\t%b\t%#x\n", x, x, x)             //print format type of a, tab, then binary value tab, hex value, new lime
	fmt.Println(`"Allows you to print quote marks"`) // reference to the ` symbol

	str := fmt.Sprintf("%T\t%b\t%#x\n", x, x, x) //Sprint converts to a string
	fmt.Println(str)

	//Fprint is used for printing to a file
}

/**
Making a new type and converting
*/
func convert() {
	type sInt int
	var num sInt
	num = 100
	fmt.Printf("%T\n", num) //num is of type sInt not int and thus cannot be converted to an int
	a := 42
	a = int(num) //int(value) converts a value to an int, in the from T(x) where T is the desired type and x is a value which can be converted to type T
	fmt.Println(a)

}

var x int = 42
var y string = "James Bond"
var z bool = true

/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/
func test1() {
	fmt.Println("Test 1.1 - 1.3")
	//x := 42
	//y := "James Bond"
	//z := true
	fmt.Println(x, y, z)
	ansT := fmt.Sprintf("%T\t%T\t%T", x, y, z)
	fmt.Println(ansT)
	ansV := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(ansV)

	fmt.Println("\nTest 1.4")
	type epicInt int
	var x epicInt
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

	fmt.Println("\nTest 1.5")

	var y int
	fmt.Printf("%v\n", x) //is 42 from prev exercise
	fmt.Printf("%T\n", x)
	x = 24
	fmt.Printf("%v\n", x)
	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)

}

/*
defining constants
Constants have their expected type and value are are immutable
*/
func con() {
	//Different ways of defining values as constants
	const (
		jeff    = "my name is"
		notTrue = false
	)

	const scrub bool = true

	fmt.Printf("%T\t%v\n", jeff, jeff)
	fmt.Printf("%T\t%v\n", notTrue, notTrue)
	fmt.Printf("%T\t%v\n", scrub, scrub)

	fmt.Println(jeff)
	fmt.Println(notTrue)
	fmt.Println(scrub)

	//iota - used to count up in constants
	const (
		a0 = iota
		a1
		a2
	)
	fmt.Println(a0, a1, a2)

	//using iota in bit shift to denote memory

	const (
		_  = iota             // _ discard char
		kb = 1 << (iota * 10) // 1 shift left 10
		mb = 1 << (iota * 10) // 1 shift left 20
		gb = 1 << (iota * 10) // 1 shift left 30
	)
	fmt.Println(kb, mb, gb)
	fmt.Printf("%b\n", kb)
	fmt.Printf("%b\n", mb)
	fmt.Printf("%b\n", gb)
}

/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/
func test2() {
	//test 2.1 - print a number in base 10, base 2 and base 16
	num := 100
	fmt.Printf("%d\t%b\t%X\n", num, num, num)

	//test 2.2 - comparing values
	a := 10
	b := 100
	g := a == b
	h := a <= b
	j := a != b
	k := a < b
	l := a > b
	fmt.Println(g, h, j, k, l)

	//test 2.3 - constants
	const t23 = 23
	const t231 int = 23
	fmt.Printf("%v\t%v\n", t23, t231)

	//test 2.4 - bit shift
	a = 7
	fmt.Printf("%d\t%b\t%X\n", a, a, a)
	a = a << 1
	fmt.Printf("%d\t%b\t%X\n", a, a, a)

	//test 2.5 - raw string literal
	c := `Literal string`
	fmt.Println(c)

	//test 2.6 - iota
	const (
		_  = 2019 + iota
		n1 = 2019 + iota
		n2 = 2019 + iota
		n3 = 2019 + iota
		n4 = 2019 + iota
	)
	fmt.Println(n1, n2, n3, n4)
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

func groupData()	{
	//Go has basic arrays which are defined like an array in Java and has a fixed size, so we use slices instead

	sliceX := []int{1,2,3,4,5}
	fmt.Println(sliceX[1])		//print out the value in position 1. Slices also start at 0 like arrays


	//Go loop though a slice you must take the index and value, but _ is used as we don't care about the index and can't make a var without using it
	for _, v := range sliceX {
		fmt.Println(v)
	}

	//Use colon to select start and end points for print where the value at each end is the start / end point, no end point prints to end of the slice
	fmt.Println(sliceX[1:4])	//prints {2 3 4}

	//Append all of sliceX into sliceY
	sliceY := append(sliceX, 6, 7, 8, 9)
	fmt.Println(sliceY)

	//Take parts of sliceX and use them in sliceZ
	sliceZ := append(sliceY[0:2], 11, 12 )
	fmt.Println(sliceZ)

	//Remove unwanted values in a slice by replacing the slice with only the value from the only slice we want to keep
	sliceX = append(sliceX[0:2], sliceX[3:5]...)	//... operator used here to fix typing error as it makes it treat slice x as an int rather than a series of ints
	fmt.Println(sliceX)

	//Make command is used to better manage memory by presetting the size of the array a slice is built one. If we outgrown the array it is rebuilt with an array double the size
	makeX := make([]int, 10, 100)	//slice of 10, underlying array of 100
	fmt.Println(makeX[2])	// Value of 0 as all empty data points in an int slice are set to 0.

	//Multi dimensional slices exist and are denoted by [][]TYPE

	fmt.Println()
	fmt.Println()

	//Maps in Go

	//Define a map, map{key}value { data, ..., data, }
	m := map[string]int{
		"Josh"		:  741,
		"Callum"	: 6000,
		"Jack" 		: 2421,
	}

	//Get values from a map
	fmt.Println(m["Josh"])
	fmt.Println(m["Jack"])
	fmt.Println(m["Callum"])

	//Verify if a key has a value
	// traditionally written as v, ok := ... but we don't care about the value as it will just be the 0 value if the key has no value mapped to it.
	_, ok := m["James"]
	fmt.Println(ok)

	//Use ok as a boolean to decide if we should enter an if statement
	if v, ok := m["James"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("No value found")
	}

	//Import a new value to map and recheck
	m["James"] = 6597
	if v, ok := m["James"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("No value found")
	}

	//For each key and value par in the map m, print them
	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "Callum")	//Deletes Callum from the map, if someone who doesn't exist is deleted no error is thrown

}

func test4()	{

}