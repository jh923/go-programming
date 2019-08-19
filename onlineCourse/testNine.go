/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wgt1 sync.WaitGroup
var wgt3 sync.WaitGroup
var countT93 = 0

type t92Person struct 	{
	first string
	last string
}

type t92Human interface {
	speak()
}

func (p *t92Person) speak()	{
	fmt.Println("My name is", p.first, p.last)
}

func t92SaySomething(h t92Human)	{
	h.speak()
}

func main()	{
	wgt1.Add(2)
	fmt.Println("Test 9.1")
	go func()	{
		fmt.Println("Function 1")
		wgt1.Done()
	}()
	go func()	{
		fmt.Println("Function 2")
		wgt1.Done()
	}()
	wgt1.Wait()

	p1 := t92Person{
		first: "Jeff",
		last: "Jones",
	}

	fmt.Println("Test 9.2")
	t92SaySomething(&p1)
	//t92SaySomething(p1) will not work

	fmt.Println("Test 9.3")
	lmt := 100
	wgt3.Add(lmt)
	for i := 0; i < lmt; i++	{
		go func() {
			localCount := countT93
			runtime.Gosched()
			localCount++
			countT93 = localCount
			wgt3.Done()
		}()
	}
	wgt3.Wait()
	fmt.Println(countT93, "This count has a race condition. As shown in the terminal" +
		" with he -race flag")
}