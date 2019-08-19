package main

import (
	"fmt"
)

func main()	{
	c := make (chan int)
	for i := 0; i < 10; i++	{
		go func(){
			for i := 0; i < 10; i++	{
				c <- i
			}
		}()
	}
	receive(c)

}

func receive(c <-chan int)  {
	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}
}