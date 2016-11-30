package main

import (
	"fmt"
	_ "strconv" //Package to convert String to other Go lang types
	"strings"
)

//String operations example

func main() {
	fmt.Println("Hello, playground")
	//Contains function
	fmt.Println(strings.Contains("India", "in"))
	fmt.Println(strings.Contains("India", "dia"))

	//ContainsAny

	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "y || z")) //OR condition
	fmt.Println(strings.ContainsAny("failure", "i & e"))  //AND condition

	//Count
	fmt.Println(strings.Count("cheese", "e"))

	//HasPrefix
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))

	//HasSuffix
	fmt.Println(strings.HasSuffix("McClatchy", "Mx"))
	fmt.Println(strings.HasSuffix("Amigo", "O"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))

	//Index
	fmt.Println(strings.Index("wipro", "ro"))
	fmt.Println(strings.Index("wipro", "dmr"))

	//Join - func Join(a []string, sep string) string
	s := []string{"This", "is an", "example"}
	fmt.Println(strings.Join(s, " "))

	//func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("This is an emmom", "m", "r", 2))
	fmt.Println(strings.Replace("This is an emmom", "m", "r", -1)) //There is no limit on number of replacements - find and replace all

	//func Split(s, sep string) []string
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))

	//ToLower, ToUpper
	fmt.Println(strings.ToLower("Gopher"))
	fmt.Println(strings.ToUpper("Gopher"))
}
