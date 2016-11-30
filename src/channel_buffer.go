/*Channel buffer */

package main

import (
	"fmt"
)



var (
	status = make(chan string, 3) //Buffered channel
)

func main() {
	go pay_complete(status)
	i := <-status
	fmt.Println(i)

	go inv_check(status)
	i = <-status
	fmt.Println(i)

	go order_confirm(status)
	i = <-status
	fmt.Println(i)
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
