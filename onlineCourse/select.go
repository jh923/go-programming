package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// Send data
	go send(even, odd, quit)

	// Receive data
	receive(even, odd, quit)

	fmt.Println("Program will now exit.")
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
	quit <- true
	close(quit)
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v, ok := <-even:
			if ok {
				fmt.Println(v, "was received from the even channel")
			}
		case v, ok := <-odd:
			if ok {
				fmt.Println(v, "Was received from the odd channel")
			}
		case <-quit:
			return
		}
	}
}
