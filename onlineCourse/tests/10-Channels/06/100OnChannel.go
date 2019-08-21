package main

import "fmt"

func main()	{
	c := make (chan int)

	go func(){
		for i := 0; i < 100; i++	{
			c <- i
		}
		close(c)
	}()

	receive(c)
	fmt.Println("about to exit")
}


func receive(c <- chan int)	{
	for v:= range c	{
		fmt.Println(v)
	}
}