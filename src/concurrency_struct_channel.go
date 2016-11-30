package main

import (
	"fmt"
)

/*Sending Struct through channels */

type SKU struct {
	StockCode string  `json:"stock_code"`
	Size      string  `json:"size"`
	PIN       int     `json:"pin_code"`
	Price     float32 `json:"price"`
	InStock   bool    `json:"in_stock"`
}

var (
	m1 = make(chan SKU)
)

func main() {
	go response() //Initiate go routine

	message1 := <-m1
	fmt.Print(message1)
}

func response() {

	m1 <- SKU{StockCode: "AB123456", Size: "M", PIN: 400001}

}
