package handle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ISSuh/go-todo/db"
	"github.com/ISSuh/go-todo/todo"

	"github.com/gorilla/mux"
)

type Result struct {
	Status bool   `json: "status"`
	Id     int    `json: id, omitempty`
	Err    string `json: "err, omitempty"`
}

type TodoHandle struct {
	Storage db.Storage
}

func (app *TodoHandle) Initialize() {
	localStorage := &db.LocalStorage{}
	localStorage.InitLocalStorage()

	app.Storage = localStorage
}

func (app *TodoHandle) TodoItem(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		app.GetItemList(res, req)
	case "POST":
		app.PostItem(res, req)
	default:
		res.Write([]byte("Invalid HTTP method"))
	}
}

func (app *TodoHandle) TodoItemById(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		app.GetItem(res, req)
	case "DELETE":
		app.DeleteItem(res, req)
	default:
		res.Write([]byte("Invalid HTTP method"))
	}
}

func (app *TodoHandle) GetItem(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	todoItem, err := app.Storage.GetItem(id)
	if err != nil {
		result, _ := json.MarshalIndent(Result{Status: false, Err: err.Error()}, "", "  ")
		res.Write([]byte(result))
		return
	}

	result, _ := json.MarshalIndent(todoItem, "", "  ")
	res.Write([]byte(result))
}

func (app *TodoHandle) PostItem(res http.ResponseWriter, req *http.Request) {
	todoItem := todo.TodoItem{}
	err := json.NewDecoder(req.Body).Decode(&todoItem)

	fmt.Println(todoItem)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := app.Storage.AddItem(todoItem)
	if err != nil {
		result, _ := json.MarshalIndent(Result{Status: false, Id: id, Err: err.Error()}, "", "  ")
		res.Write([]byte(result))
		return
	}

	result, _ := json.MarshalIndent(Result{Status: true, Id: id}, "", "  ")
	res.Write([]byte(result))
}

func (app *TodoHandle) DeleteItem(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	err := app.Storage.DeleteItem(id)
	if err != nil {
		result, _ := json.MarshalIndent(Result{Status: false, Id: id, Err: err.Error()}, "", "  ")
		res.Write([]byte(result))
		return
	}

	result, _ := json.MarshalIndent(Result{Status: true, Id: id}, "", "  ")
	res.Write([]byte(result))
}

func (app *TodoHandle) GetItemList(res http.ResponseWriter, req *http.Request) {
	list, _ := app.Storage.GetItemList()
	json, _ := json.MarshalIndent(list, "", "  ")
	res.Write([]byte(json))
}
