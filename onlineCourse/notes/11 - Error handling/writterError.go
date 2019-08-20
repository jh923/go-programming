package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main()	{
	f, err := os.Create("Evil plan to take over the world.txt")
	if err != nil {
		fmt.Println(err)
		return	//Quit program
	}
	defer f.Close()	//Set the file to close when program is about to exit

	r := strings.NewReader("Step 1: Create generically engineered cat girls.\n" +
								"Step 2: ???\n" +
								"Step 3: World domination.")

	io.Copy(f, r)
}
