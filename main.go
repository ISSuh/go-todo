package main

import (
	"log"
	"net/http"

	"github.com/ISSuh/go-todo/todo/handle"

	"github.com/gorilla/mux"
)

func main() {
	app := handle.TodoHandle{}
	app.Initialize()

	router := mux.NewRouter()
	router.HandleFunc("/login", app.Login)
	router.HandleFunc("/item", app.TodoItem)
	router.HandleFunc("/item/{id}", app.TodoItemById)

	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}
