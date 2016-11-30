package main

import (
	"fmt"
	
)

type Name struct {
	firstName string `json:"fname"`
	lastName string `json:"lname"`
}

type Address struct{
	doorNumber string `json:"dnum"`
	building string `json:"building"`
	city	string `json:"city"`
	pin	string `json:"pin"`
}


type Customer struct { //Example of an anonymous struct
	Name
	Address
}

var c = Customer{Name{"Anand", "Hariharan"},Address{"24","Electronic City","Bangalore","560001"}}

func main() {
	
	fmt.Println("Hello, playground")
	fmt.Println(c.Name.firstName)
	fmt.Println(c)
	
}
