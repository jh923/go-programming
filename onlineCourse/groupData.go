package main

import "fmt"

func main()  {
	groupData()
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
