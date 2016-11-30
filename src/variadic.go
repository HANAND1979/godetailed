package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1,2,3,4))
	n:=[]int{1,2,3,4} //A slice can also be sent to a variadic function
	fmt.Println(sum(n...))

	s:=[]string{"a", "b", "c"}
	print_func(s...)
}


func sum(num ...int) (t int){

fmt.Println(num)
t=0
for _,num := range num{
	t +=num
}
return t
}

func print_func(s ...string){
	for _,st:=range s{
		fmt.Println(st)
	}
}