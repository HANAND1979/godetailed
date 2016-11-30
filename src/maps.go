// maps
package main

import (
	"fmt"
)

func main() {
	//illustration of maps [key-value pairs]
	scrip := make(map[string]float32)
	scrip["nifty"] = 7000
	scrip["sensex"] = 22000
	scrip["ITC"] = 230
	fmt.Println(scrip)
	delete(scrip, "nifty")
	fmt.Println(scrip, len(scrip))
	_, is_present := scrip["nifty"]
	fmt.Println("is nifty present?:", is_present)

	//illustration of ranges
	for i, v := range scrip {
		fmt.Println(i, v)
	}
	//Example of use of range in the case of matrices
	//standard variable description
	var even [5]int
	even[0] = 2
	even[1] = 4
	even[2] = 6
	even[3] = 4
	even[4] = 8
	//even[5] = 10 //error will be thrown here. contrast this with how we can expand a slice dynamically later.
	odd := [5]int{1, 3, 5, 7, 9} //implicit array declaration

	matrix := make(map[int]int)
	for _, x := range even {
		for _, y := range odd {
			fmt.Println(x, y)
			matrix[x] = x * y
			fmt.Println(matrix)
		}
	}
	fmt.Println(matrix) //notice from the output that there is no sorted order to a map
}
