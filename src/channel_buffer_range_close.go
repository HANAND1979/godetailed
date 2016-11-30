/* Range over channel and channel closure */
/*the Playground always returns timestamps to be 1257894000 (Tue, 10 Nov 2009 23:00:00, the date Go became a public open source project!) so the difference between two timestamps is always 0. They do this so programs behave deterministically and can be cached. You'll have to run it locally to see the real execution times.*/
// Also note the Channel Direction specification below in the function calls 

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	status = make(chan string)
)

func main() {

	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	count_of_orders := 5

	go inv_check(status, count_of_orders)

	elapsed := time.Since(start)

	print_status(status, count_of_orders, elapsed)

}

func inv_check(status chan<- string, n int) {
	for i := 1; i <= n; i++ {
		msg := "Inventory checked " + strconv.Itoa(i)
		time.Sleep(5 * 1e9)
		status <- msg
	}
	close(status) //If channel is not explicitly closed, the print_status function where range over channel is being done will result in a deadlock
}

func print_status(status <-chan string, n int, elapsed time.Duration) {
	for s := range status {
		fmt.Println(s)
	}
	
	fmt.Println(elapsed)
}
