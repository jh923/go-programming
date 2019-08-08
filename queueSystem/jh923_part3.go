package main

/**
**	This project was not done using raptor so I wrote the local version of the random input for my use and
**	the git link for the benefit of others, however, you will have to uncomment it fot the program to comoile and
**  comment out my local version
**/
import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"strconv"
	"time"
)

//replaced m and n with 10 for testing purposes
func main() {
	lect := make(chan bool)           // creates a synchronous channel
	wait := make(chan chan string, 3) // creates an asynchronous
	// channel of size n, where n = 3
	for i := 0; i < 2; i++ { // A loop is ued so the amount of lecturers can be easily changed and to reduce duplicate code
		go lecturer(wait, lect, randomdata.SillyName())
	}
	for i := 0; i < 20; i++ { //create m number students, where m = 20.
		go student(wait, lect, randomdata.SillyName())
		time.Sleep(1 * time.Second)
	}
	time.Sleep(1000 * time.Second)
}

//Check the queue, and either brings a person who has been waiting or does nothing if the queue is empty
func lecturer(wait chan chan string, lect chan bool, name string) {
	for {
		<-wait //Check waiting queue
		s := strconv.Itoa(len(wait))
		fmt.Println("There are " + s + " students waiting")
		lect <- true
		fmt.Println("New student is called by: " + name + ".")
		time.Sleep(4 * time.Second) //Simulate meeting time
	}
}

//The student goes to the meeting room, waits until the lecturer class them, they have their meeting, then the student leaves
//The student checks the current length of the queue,if there is a space they will join the queue, if it is full they will not.
func student(wait chan chan string, lect chan bool, name string) {
	myself := make(chan string)
	fmt.Println(name + " wants a meeting.")
	if len(wait) <= 2 { // where they are <= n-1
		wait <- myself //Goes to the waiting room
		fmt.Println(name + "is waiting.")
		<-lect //Gets called in by the lecturer
		fmt.Println(name + " is in a meeting!")
	} else {
		fmt.Println(name + " says bye!")
	}
}
