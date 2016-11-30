package main

import "fmt"

type SKU struct {
	StockID string
	pinCode string
}

var (
	s = SKU{"A123","560001"}
	p = &s
)

func main() {
	fmt.Print(checkAvailability(p))
}

func checkAvailability(p *SKU) (avl bool){
	fmt.Println("Checking availability for", p.StockID)
	return true
}