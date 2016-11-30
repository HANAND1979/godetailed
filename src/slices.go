// slices
package main

import (
	"fmt"
)

func main() {
	//slices are parts of arrays
	s := make([]string, 5, 10) //declaring a slice
	fmt.Println(s)             //returns a blank matrix
	s[0] = "apple"
	s[1] = "bat"
	s[2] = "cat"
	fmt.Println(len(s)) //this returns the declared length of the slice
	s = append(s, "dog")
	s = append(s, "eel")
	s = append(s, "fox")
	fmt.Println("set:", s)
	fmt.Println(len(s)) //observe how the length changes
	s = append(s, "goat")
	s = append(s, "horse")
	s = append(s, "iguana")
	t := make([]string, len(s))
	copy(t, s)          //copying a slice
	fmt.Println(len(s)) //observe how the length changes beyond size of 10 declared earlier
	fmt.Println("get:", s[5])
	fmt.Println("appended set:", s)
	fmt.Println(s[2:5], s[2:], s[:9])
	fmt.Println("copied slice:", t)

	//Example of use of range for slices
	for q, r := range s {
		fmt.Print(q, r)
	}
}
