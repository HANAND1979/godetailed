/* This package contains basic mathematical functions
(c) Anand Hariharan 2016
*/

package interfaces

import (
	"fmt"
	"math"
)

type sq_root interface {
	root() (float64, string)
	speak() string
}

type real_number struct {
	real float64
}

type Int struct {
	I int
}

//Returns root of a real number
func (r real_number) root() (float64, string) {
	y := r.real * r.real
	return y, "hello real"
}
func (r real_number) speak() string {
	return "You sent me a real number"
}

func (intr Int) Root(x float64) (y float64) {
	//y:="I don't work with integers !"
	y = math.Sqrt(x)
	return y
}
func (intr Int) speak() string {
	return "You sent me an integer"
}

func find_root(r sq_root) {
	fmt.Println(r.root())
	fmt.Println(r.speak())
}
