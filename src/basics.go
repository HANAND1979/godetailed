// example1
package main

//importing packages
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var hello string

//hello:="hi, this will throw an error"

//Basic arithmetic functions
func main() {
	fmt.Println("Hello Git World!")
	hello = "Hello World!"
	fmt.Println(hello)

	//constant declaration
	const Pi, radius = 3.14159, 2
	fmt.Println(Pi * math.Sqrt(radius))

	//complex maths
	var loc1, loc2 complex128
	loc1 = 5 + 7i
	loc2 = 7 + 5i
	fmt.Println(loc1 * loc2)

	a, b := 8.1, 9.7 //short variable declaration
	fmt.Println(a * b)
	fmt.Println(math.Pi) //notice that Pi is different from the constant Pi declared above
	fmt.Println(rand.Intn(7))

	//type conversions
	var x, y int = 7, 9
	var f float64 = math.Sqrt(float64(x*x + y*y)) //conversion to float64
	var z uint = uint(f)                          //conversion to uint
	fmt.Println("Type conversion", x, y, f, z)

	//Demo of function
	m, n := func_demo("Google")
	fmt.Println(m, n)

	//Demo of control statements
	defer func_for(3) //illustration of defer
	if_output := func_if(9, 9)
	fmt.Println(if_output)
	func_case()

	//Demo of types and methods
	fr := Mynumber(7)
	fmt.Println("The square of my number is", fr.func_square())

	post1 := My_blog{"A Go blog", "Welcome to my blog on Go", "16-Feb-2016"}
	fmt.Println(My_blog.func_postblog(post1))
	post2 := My_blog{"", "", ""}
	My_blog.say_hi(post2)

}

//demo of function
func func_demo(s string) (string, int) {

	a := "Hello" + " " + s
	b := len(s)
	return a, b
}

func func_for(counter int) {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += i
		defer print(sum, "\n") //use of defer inverts the output
	}

}

func func_if(p, q int) bool {

	if p > q {
		return true
	} else {
		return false
	}

}

func func_case() {
	fmt.Print("When is Monday?")
	today := time.Now().Weekday()
	switch time.Monday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("One day away")
	case today + 2:
		fmt.Println("Two days away")
	default:
		fmt.Println("Too far away")

	}
}

//example of type and methods

type Mynumber int

func (n Mynumber) func_square() Mynumber {
	return (n * n)
}

//more complex example of type and methods
//this demonstrates how Go implements object orientation
type My_blog struct { //My_blog is equivalent to a class declaration here
	title     string
	post      string
	posted_on string
}

func (message My_blog) say_hi() { //This is equivalent to a method declaration on the class
	fmt.Println("Hi from my blog")
}

func (blog My_blog) func_postblog() string {
	blog_stream := "<xml><title>" + blog.title + "</title><post>" + blog.post + "</post><time>" + blog.posted_on + "</time></xml>"
	return blog_stream
}
