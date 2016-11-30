// web-get
package main

import (
	"fmt"
	_ "log"
	"net/http"
)

func main() {

	//Demo 1
	res, err := http.Get("https://www.google.com/nunjucks")

	if res.StatusCode == 404 {
		fmt.Println(res.Status)
		fmt.Printf("404, Page not found !\n")
	}
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

	//Demo 2 with panic and recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic ! Sorry, something went wrong")
		}
	}()

	url := "https:/thisisnotaurl"
	status := get_url(url)
	if status == 404 {
		fmt.Printf("404, Page not found !")
	}
}

func get_url(s string) (status int) {
	res, err := http.Get(s)
	if res == nil {
		panic("Whoa ! Something went wrong !")
	}
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
	status = res.StatusCode
	return status
}
