/*
1. Concurrent goroutines
2. Illustration of Empty Struct Channel for synchronization
3. Illustration of channel closure
   Illustration of Empty Struct Channel - Refer http://guzalexander.com/2013/12/06/golang-channels-tutorial.html
*/

package main

import (
	"fmt"
	"time"
)

var (
	status       = make(chan string)
	payment_done = make(chan struct{})
)

func main() {
	count_of_orders := 3
	sku := make([]string, count_of_orders)
	sku[0] = "alpha"
	sku[1] = "beta"
	sku[2] = "gamma"
		
	elapsed := inv_check(status, count_of_orders, sku)
	defer print_status(status, count_of_orders, elapsed)
	go payment(payment_done)
	<-payment_done
	
	
}

func inv_check(status chan<- string, count_of_orders int, sku []string) (elapsed time.Duration) {
	start := time.Now()
	for i := 0; i < count_of_orders; i++ {
		m := sku[i]
		go func() {
			fmt.Printf("Goroutine initiated for SKU %v\n", m)
			msg := "Inventory checked " + m
			time.Sleep(2 * 1e4)
			status <- msg

		}()
	}
	
	elapsed = time.Since(start)
	return elapsed

}

func payment(payment_done chan struct{}) {

	time.Sleep(4 * 1e9)
	close(payment_done)
}

func print_status(status <-chan string, n int, elapsed time.Duration) {

	select {
	case <-status:
		for i := 1; i <= n; i++ {
			msg := <-status
			fmt.Println(msg)
		}

	case <-payment_done:
		fmt.Println("Payment successful")

	default:
		fmt.Println("No message received. Non-blocking channel")
	}
	fmt.Printf("Elapsed time %s\n", elapsed)
}
