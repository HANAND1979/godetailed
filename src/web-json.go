// web-json
package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"strconv" //Package to convert String to other Go lang types

	"github.com/gorilla/mux"
)

type SKU struct {
	StockCode string  `json:"stock_code"`
	Size      string  `json:"size"`
	PIN       int     `json:"pin_code"`
	Price     float32 `json:"price"`
	InStock   bool    `json:"in_stock"`
}

type stocks SKU

func main() {
	r := mux.NewRouter().StrictSlash(true) //When true, this ensures that your application will always see the path as specified in the route
	r.HandleFunc("/stock", stockHandler)   //If strict slash is not enabled, this call will result in a 404 page not found error
	r.HandleFunc("/stock/{stock_code}", stockHandler)
	r.HandleFunc("/stock/{stock_code:[a-z]{2}[0-9]{6}}/{size:(?:S|M|L)}/{pin_code:[1-9]{1}[0-9]{5}}", stockHandlerPIN)
	r.HandleFunc("/json", stockHandlerJSON).Methods("POST") //Handle a request posted as a JSON - a REST API
	http.Handle("/", r)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func stockHandlerJSON(w http.ResponseWriter, r *http.Request) {

	//Example of posting JSON request and decoding request
	fmt.Fprintf(w, "Thank you for your enquiry. Processing your request...stand by...\n")
	var my_stock stocks
	if r.Body == nil {
		http.Error(w, "Invalid request. Please send a JSON formatted request", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&my_stock)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//Standadrd query response
	my_stock.InStock = true
	if my_stock.InStock {
		fmt.Fprintf(w, "Yes, SKU %v of size %v is AVAILABLE at location %v \n", my_stock.StockCode, my_stock.Size, my_stock.PIN)
	} else {
		fmt.Fprintf(w, "Sorry ! SKU %v of size %v at location %v is not available \n", my_stock.StockCode, my_stock.Size, my_stock.PIN)
	}

	//JSON response
	json_response := stocks{my_stock.StockCode, my_stock.Size, my_stock.PIN, 300, true}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	err = enc.Encode(json_response)
	if err != nil {
		// if encoding fails, create an error page with code 500.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func stockHandlerPIN(w http.ResponseWriter, r *http.Request) {
	stock_code := mux.Vars(r)["stock_code"]
	size := mux.Vars(r)["size"]
	pin_code, err := strconv.Atoi(mux.Vars(r)["pin_code"]) //String conversion example
	if err != nil {
		log.Fatal(err)
	}
	if stock_code != "" {
		fmt.Fprintf(w, "Checking availability for %v size %v at PIN code %v \n", stock_code, size, pin_code)

		//query result as JSON
		sku := SKU{stock_code, size, pin_code, 30, true}

		// set the Content-Type header.
		w.Header().Set("Content-Type", "application/json")

		enc := json.NewEncoder(w)
		err := enc.Encode(sku)
		//fmt.Printf("Hi %v", w)
		if err != nil {
			// if encoding fails, create an error page with code 500.
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		fmt.Fprintln(w, "Nothing to check")
	}

}

func stockHandler(w http.ResponseWriter, r *http.Request) {
	stock_code := mux.Vars(r)["stock_code"]

	if stock_code != "" {
		fmt.Fprintf(w, "Checking availability for %v  \n", stock_code)

		type sku_basic struct {
			StockCode string `json:"stock_code"`

			Price   float32 `json:"price"`
			InStock bool    `json:"in_stock"`
		}

		//query result as JSON
		sku := sku_basic{stock_code, 30, true}

		// set the Content-Type header.
		w.Header().Set("Content-Type", "application/json")

		enc := json.NewEncoder(w)
		err := enc.Encode(sku)

		if err != nil {
			// if encoding fails, create an error page with code 500.
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		fmt.Fprintln(w, "Nothing to check")
	}

}
