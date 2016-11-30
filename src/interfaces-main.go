// interfaces-main
package main

import (
	"fmt"

	"interfaces"
)

var (
	i interfaces.Int
)

func main() {
	r := i.Root(4.5)
	fmt.Println("Square root is", r)
}
