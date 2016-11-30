/* Concurrent goroutines with elapsed time measurement */

package main

import (
	"fmt"
	"time"
)

var (
	status = make(chan string)
)

func main() {
	count_of_orders := 5
	sku := make([]string, count_of_orders)
	sku[0] = "alpha"
	sku[1] = "beta"
	sku[2] = "gamma"
	sku[3] = "eps"
	sku[4] = "omg"
	elapsed := inv_check(status, count_of_orders,sku)
	print_status(status, count_of_orders, elapsed)
}

func inv_check(status chan<- string, count_of_orders int, sku []string) (elapsed time.Duration) {
	start := time.Now()
	for i := 0; i < count_of_orders; i++ {
		m := sku[i]
		//fmt.Print(m)
		go func() {
			fmt.Printf("Goroutine initiated for SKU %v\n", m)
			msg := "Inventory checked " + m
			time.Sleep(2*1e4)
			status <- msg

		}()
	}

	elapsed = time.Since(start)
	return elapsed

}

func print_status(status <-chan string, n int, elapsed time.Duration) {
	for i := 1; i <= n; i++ {
		msg := <-status
		fmt.Println(msg)
	}

	fmt.Printf("Elapsed time %s\n", elapsed)
}
