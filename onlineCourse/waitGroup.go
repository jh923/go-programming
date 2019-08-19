package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

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

	counter := 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			//Locks when entering a critical section of a shared variable and unlocks when exiting it
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter value:", counter)
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
