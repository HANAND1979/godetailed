package main

import "errors"
import "fmt"

func main() {

	out := create_panic()
	fmt.Print(out)
	defer write_string(out)
}

func create_panic() (t string) {
	err := errors.New("This is an error")
	if err != nil {
		panic(err)

	}
	return "Panic Example"
}
func write_string(s string) {
	fmt.Printf("%v \n",s)
}