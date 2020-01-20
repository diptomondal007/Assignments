package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Data struct {
	ID    int `json:"id"`
	Price int `json:"price"`
}

var dataSlice []Data
var queue []Data
var jobs chan Data
var TotalProfit int

func main() {
	go taskQueue()
	routeHandler()
}

func routeHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/", addData).Methods("POST")
	log.Fatalln(http.ListenAndServe("localhost:8000", router))
}

func addData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic("Sorry json reading error")
	}
	var data Data
	json.Unmarshal(requestBody,&data)
	queue = append(queue, data)
	log.Println("Added")

}

func taskQueue()  {
	jobs = make(chan Data,3)
	go worker(jobs)
	for {
		if len(queue)>0{
			for i:=0;i<len(queue);i++{
				jobs <- queue[i]
				queue = queue[:0]
			}
		}
	}
}
func worker(jobs <-chan Data) {
	for job := range jobs {
		TotalProfit += job.Price
		log.Println(TotalProfit ,"after adding ",job.ID )
		time.Sleep(time.Second * 2)
	}
}

