package handle

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ISSuh/go-todo/todo"

	"github.com/gorilla/mux"
)

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

	todoItem, err := app.Contents.GetItem(id)
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
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := app.Contents.AddItem(todoItem)
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

	err := app.Contents.DeleteItem(id)
	if err != nil {
		result, _ := json.MarshalIndent(Result{Status: false, Id: id, Err: err.Error()}, "", "  ")
		res.Write([]byte(result))
		return
	}

	result, _ := json.MarshalIndent(Result{Status: true, Id: id}, "", "  ")
	res.Write([]byte(result))
}

func (app *TodoHandle) GetItemList(res http.ResponseWriter, req *http.Request) {
	list, _ := app.Contents.GetItemList()
	json, _ := json.MarshalIndent(list, "", "  ")
	res.Write([]byte(json))
}
