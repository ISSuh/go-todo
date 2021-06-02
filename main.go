package main

import (
	"log"
	"net/http"

	"github.com/ISSuh/go-todo/todo"
	"github.com/gorilla/mux"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}
