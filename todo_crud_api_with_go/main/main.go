package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	todo_crud_api_with_go "todo_crud.com"
)

var Todos []todo_crud_api_with_go.ToDo

type Res struct {
	ResBody string `json:"response"`
}

var errorRes = Res{ResBody: "Operation Failed"}
var successRes = Res{ResBody: "Operation successful"}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	keys := r.URL.Query()
	if len(keys.Get("id")) > 0 {
		for _, todo := range Todos {
			if todo.ID == keys.Get("id") {
				json.NewEncoder(w).Encode(todo)
			}
		}
	} else {
		json.NewEncoder(w).Encode(Todos)
	}

}

func postTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("Json Reading failed")
	}
	var todo todo_crud_api_with_go.ToDo
	todo.CreatedOn = time.Now()
	json.Unmarshal(responseBody, &todo)
	Todos = append(Todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := r.URL.Query()
	if len(key.Get("id")) > 0 {
		found := false
		for _, todo := range Todos {
			if todo.ID == key.Get("id") {
				found = true
			}
		}
		if found {
			keyInt, err := strconv.Atoi(key.Get("id"))
			if err != nil {
				json.NewEncoder(w).Encode(errorRes)
			}
			Todos = append(Todos[:keyInt], Todos[keyInt+1:]...)
			json.NewEncoder(w).Encode(successRes)
		} else {
			json.NewEncoder(w).Encode(errorRes)
		}
	} else {
		json.NewEncoder(w).Encode(Todos)
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := r.URL.Query()
	if len(key.Get("id")) > 0 {
		found := false
		for _, todo := range Todos {
			if todo.ID == key.Get("id") {
				found = true
			}
		}
		if found {
			keyInt, err := strconv.Atoi(key.Get("id"))
			if err != nil {
				json.NewEncoder(w).Encode(errorRes)
			}

			var anotherTodo todo_crud_api_with_go.ToDo
			requestBody, err := ioutil.ReadAll(r.Body)
			json.Unmarshal(requestBody, &anotherTodo)
			Todos[keyInt] = anotherTodo
			json.NewEncoder(w).Encode(successRes)
		} else {
			json.NewEncoder(w).Encode(errorRes)
		}
	} else {
		json.NewEncoder(w).Encode(Todos)
	}
}

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", getTodo).Methods("GET")
	router.HandleFunc("/", postTodo).Methods("POST")
	router.HandleFunc("/", deleteTodo).Methods("DELETE")
	router.HandleFunc("/", updateTodo).Methods("PATCH")

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func main() {
	Todos = append(Todos, todo_crud_api_with_go.ToDo{
		ID:          "0",
		Title:       "First Todo",
		Body:        "This is first Todo",
		IsCompleted: false,
		CreatedOn:   time.Now(),
	})
	handleRequest()
}
