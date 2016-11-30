//Example of array, slices, maps and ranges
package main

import (
	"fmt"
)

func main() {
	//Examples of range is provided for each of the data structures below

	//standard variable description
	var even [5]int
	even[0] = 2
	even[1] = 4
	even[2] = 6
	even[3] = 4
	even[4] = 8
	//even[5] = 10 //error will be thrown here. contrast this with how we can expand a slice dynamically later.
	odd := [5]int{1, 3, 5, 7, 9} //implicit array declaration
	fmt.Println(even, even[3])
	fmt.Println(len(odd), odd)

	var sumproduct int
	for i := 0; i < 5; i++ {
		sumproduct += (even[i] * odd[i])
		fmt.Println(sumproduct)
	}
	fmt.Println(sumproduct)

}
