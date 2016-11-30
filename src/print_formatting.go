package main

import (
	"fmt"
	"math"
	_"errors"
)

//Examples of fmt package functions
func main() {

	
	fmt.Println("Hello, playground")
	
	//Print value and type 
	s:="wipro limited"
	m,err:=fmt.Printf("Hello %v\n",s) 
	fmt.Println(m,"bytes written \n")
	
	m,err=fmt.Printf("Type is %T\n",s) 
	
	//Boolean formatting
	m,err=fmt.Printf("%t\n",7>8) 
	fmt.Println(m,"bytes written \n")
	
	//Numeric formatting 
	m,err=fmt.Printf("%f\n",math.Sqrt(2))
	m,err=fmt.Printf("%E\n",math.Sqrt(2))  
	
	if err!=nil{
		fmt.Print("error %v",err)
	}
	
	
}
