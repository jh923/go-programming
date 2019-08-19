/**
The tests can be found at https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wgt1 sync.WaitGroup
var wgt3 sync.WaitGroup
var wgt4 sync.WaitGroup
var wgt5 sync.WaitGroup
var countT93, countT94 = 0, 0
var counterT95 int64 = 0

type t92Person struct {
	first string
	last  string
}

type t92Human interface {
	speak()
}

func (p *t92Person) speak() {
	fmt.Println("My name is", p.first, p.last)
}

func t92SaySomething(h t92Human) {
	h.speak()
}

func main() {
	wgt1.Add(2)
	fmt.Println("Test 9.1")
	go func() {
		fmt.Println("Function 1")
		wgt1.Done()
	}()
	go func() {
		fmt.Println("Function 2")
		wgt1.Done()
	}()
	wgt1.Wait()

	p1 := t92Person{
		first: "Jeff",
		last:  "Jones",
	}

	fmt.Println("Test 9.2")
	t92SaySomething(&p1)
	//t92SaySomething(p1) will not work

	fmt.Println("Test 9.3")
	lmt := 100
	wgt3.Add(lmt)
	for i := 0; i < lmt; i++ {
		go func() {
			localCount := countT93
			runtime.Gosched()
			localCount++
			countT93 = localCount
			wgt3.Done()
		}()
	}
	wgt3.Wait()
	fmt.Println(countT93, "This count has a race condition. As shown in the terminal"+
		" with he -race flag")

	fmt.Println("Test 9.4")
	wgt4.Add(lmt)
	muT94 := sync.Mutex{}
	for i := 0; i < lmt; i++ {
		go func() {
			muT94.Lock()
			localCount := countT94
			localCount++
			countT94 = localCount
			muT94.Unlock()
			wgt4.Done()
		}()
	}
	wgt4.Wait()
	fmt.Println(countT94)

	fmt.Println("Test 9.5")
	wgt5.Add(lmt)
	for i := 0; i < lmt; i++ {
		go func() {
			atomic.AddInt64(&counterT95, 1)
			//fmt.Println("Counter\t", atomic.LoadInt64(&counterT95))
			wgt5.Done()
		}()
	}
	wgt5.Wait()
	fmt.Println(atomic.LoadInt64(&counterT95))

	fmt.Println("Test 9.6")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println()
}
