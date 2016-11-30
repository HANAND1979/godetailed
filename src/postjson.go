// postjson

package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type SKU struct {
	StockCode string  `json:"stock_code"`
	Size      string  `json:"size"`
	PIN       int     `json:"pin_code"`
	Price     float32 `json:"price"`
	InStock   bool    `json:"in_stock"`
}

type Metrics struct {
	ReqProcessed int           `json:"req_processed`
	AvgReqTime   float64       `json:"art"`
	TotTime      time.Duration `json:tottime`
}

var (
	metrics      Metrics
	counttoCheck int
	m1           = make(chan string, counttoCheck)
	m2           = make(chan time.Duration, counttoCheck)
	start        time.Time
)

func main() {

	start = time.Now()
	fmt.Println(start)

	counttoCheck = 10
	s := make([]SKU, counttoCheck)

	defer func() {
		recover()
		//fmt.Println(r)
	}()
	for i := 0; i < counttoCheck; i++ {
		stockCode := "AB" + strconv.Itoa(i)
		s[i] = SKU{StockCode: stockCode, Size: "M", PIN: 400001}
		//fmt.Println(s[i])
		go postJSON(s[i], m1, m2)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/metrics", writeMetrics)
	http.Handle("/", router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}

func writeMetrics(w http.ResponseWriter, r *http.Request) {

	avg_resp_time := make([]time.Duration, counttoCheck)
	cum_art := 0.0

	for i := 0; i < counttoCheck; i++ {
		<-m1
		message2 := <-m2
		avg_resp_time[i] = message2
		cum_art += avg_resp_time[i].Seconds()
		//fmt.Fprintf(w, "%v %v \n", message1, message2)
		fmt.Fprintf(w, "%v requests processed\n", i+1)
		metrics.ReqProcessed = i
	}
	elapsed := time.Since(start)
	art := cum_art / float64(counttoCheck)

	metrics.AvgReqTime = art
	metrics.TotTime = elapsed
	fmt.Fprintf(w, "Avg API response time %vs \n", metrics.AvgReqTime)
	fmt.Fprintf(w, "Total elapsed time %v \n", elapsed)
	//fmt.Printf("Avg API response time %vs", art)

}

func postJSON(toCheck SKU, m1 chan string, m2 chan time.Duration) {
	start := time.Now()
	json_input := toCheck
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(json_input)
	defer func() {

		recover()
		//fmt.Println("Unexpected error")

	}()
	response, _ := http.Post("http://130.211.25.226:8080/json", "application/json; charset=utf-8", b)
	//fmt.Println("JSON requested posted...awaiting response")
	//io.Copy(os.Stdout, res.Body) //You can use this for standard output in console

	// fmt.Printf(response.Body)This would fail because Body = io.ReadCloser
	// So we convert it to a string by passing it through a buffer first.
	// Refer http://golangcode.com/convert-io-readcloser-to-a-string/
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	jsonStr := buf.String()

	//fmt.Printf(jsonStr)

	m1 <- jsonStr
	elapsed := time.Since(start)
	fmt.Printf("Time taken to check %v is %v \n", toCheck.StockCode, elapsed)
	m2 <- elapsed
}
