package main

/**
**	This project was not done using raptor so I wrote the local version of the random input for my use and
**	the git link for the benefit of others, however, you will have to uncomment it fot the program to comoile and
**  comment out my local version
**/
import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"time"
)

//replaced m and n with 10 for testing purposes
func main() {
	lect := make(chan bool)            // creates a synchronous channel
	wait := make(chan chan string, 10) // creates an asynchronous
	// channel of size n
	go lecturer(wait, lect)
	for i := 0; i < 20; i++ {
		go student(wait, lect, randomdata.SillyName())
	}
	time.Sleep(1000 * time.Second)
}

//Check the queue, and either brings a person who has been waiting or does nothing if the queue is empty
func lecturer(wait chan chan string, lect chan bool) {
	for {
		<-wait //Check waiting queue
		lect <- true
		fmt.Println("New student is called.")
		time.Sleep(2 * time.Second) //Simulate meeting time
	}
}

//The student goes to the meeting room, waits until the lecturer class them, they have their meeting, then the student leaves
func student(wait chan chan string, lect chan bool, name string) {
	myself := make(chan string)
	wait <- myself //Goes to the waiting room
	fmt.Println(name + "is waiting.")
	<-lect //Gets called in by the lecturer
	fmt.Println(name + " is in a meeting!")
}
