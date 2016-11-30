
//Design Principle : "Don't communicate by sharing memory, Share memory by communicating" 
//watch concurrency is not parallelism https://www.youtube.com/watch?v=f6kdp27TYZs

package main

import (
	"fmt"
)

var (
	m1 = make(chan string)
	m2 = make(chan bool, 2) //This is a buffered channel

)

func main() {
	go response() //Initiate go routine
	
	message1 := <-m1
	message2 := <-m2
	message3 := <-m2
	fmt.Print(message1, " ", message2, message3)
}

func response() {

	m1 <- "hi"
	m2 <- true
	m2 <- true

}
