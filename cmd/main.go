package main

import (
	"fmt"
	"github.com/AnaSi95/todo-list-app/configs"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CounterHandler struct {
	counter int
}

func (ct *CounterHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(ct.counter)
	ct.counter++
	fmt.Fprintf(w, "Counter:", ct.counter)
}

func main() {
	log.Println("Starting the HTTP server on port 8080")
	router := mux.NewRouter().StrictSlash(true)
	log.Fatal(http.ListenAndServe("localhost:8080", router))

	configs.ConnectDB()
}
