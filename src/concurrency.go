package main

import (
	"fmt"
	"time"
)

func f(word string) {
	fmt.Println("I am going first")
	for i := 0; i < 3; i++ {
		fmt.Println(word, i)
	}

}

func g(word string, done chan bool) {
	fmt.Println("Go routine initiated")
	for i := 0; i < 3; i++ {
		fmt.Println(word, i)
	}
	done <- true

}

func main() { //the function main() itself is an implicit goroutine

	messages := make(chan string)

	f("direct") //here control passes on to next line only once this completes execution

	done := make(chan bool)
	go g("goroutine", done) // here control is passed on to the next line while this executes in the background

	go func(msg string) { //implicit go routine declaration
		fmt.Println(msg)
		messages <- "I am finished"
	}("I am yet another go routine")

	fmt.Println("I am running from the main")
	for j := 0; j < 3; j++ {
		fmt.Println(time.Now().Month(), j)
	}

	msg := <-messages
	fmt.Println(msg)
	//Block until we receive a notification from the worker on the channel.
	//return
	<-done
	//<-messages
	//If you removed the <- done line from this program, the program would exit before the worker even started.

}
