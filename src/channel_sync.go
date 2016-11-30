package main

import (
	"fmt"
)

/*Channel synchronization */

var (
	status = make(chan string, 1)
)

func main() {
	go inv_check(status) //Initiate go routine
	m := <-status
	fmt.Println(m)
	
	go pay_complete(status)
	m = <-status
	fmt.Println(m)
	
	go order_confirm(status)
	m = <-status
	fmt.Print(m)

}

func inv_check(status chan string) {
	status <- "Inventory checked"
}

func pay_complete(status chan string) {

	status <- "Payment completed"
}

func order_confirm(status chan string) {

	status <- "Order completed"
}

