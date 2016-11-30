// web-app
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	//handling requests with mux
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/hello/{name}", helloHandler)
	http.Handle("/", r)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//Simple handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if name != "" {
		fmt.Fprintf(w, "Hello %v", name)
	} else {
		fmt.Fprintln(w, "Hello World")
	}

}
