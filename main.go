package main

import (
	"log"
	"net/http"

	"github.com/ISSuh/go-todo/todo"

	"github.com/gorilla/mux"
)

func main() {
	app := todo.TodoHandle{}
	app.Initialize()

	router := mux.NewRouter()
	router.HandleFunc("/login", app.Login)
	router.HandleFunc("/item", app.TodoItem)
	router.HandleFunc("/item/{id}", app.TodoItemById)

	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}
