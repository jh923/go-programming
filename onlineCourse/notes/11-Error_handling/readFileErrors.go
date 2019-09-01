package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Open a file
	f, err := os.Open("Evil plan to take over the world.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	//Read the file
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Print contents to terminal
	fmt.Println(string(bs))
}
