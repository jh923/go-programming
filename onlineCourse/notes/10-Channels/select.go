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

// v, ok idiom used as when channel is closed, we match with the first two cases with a value of 0 when the channel
// should no longer be receiving data to match with.
// To fix this we could either keep the channels open witch is bad practice, or close the channels as we have done
// and check if the value of the chanel is ok, and if it is, we know it is a legitimate value, if it is not ok
// we can discard it like done blow.
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
