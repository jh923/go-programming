package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main()	{
	go count()
	go count()
	wg.Add(2)
	wg.Wait()

	//Getting return data from a go routine rather than dropping it
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}

func doSomething(x int) int {
	return x * 5
}

func count()	{
	for i := 0; i < 100; i++ 	{
		fmt.Println("first go routine count:", i)
	}
	wg.Done()
}
