package db

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ISSuh/go-todo/todo"
)

var todoItemJson = string(`{"item": {"time": "20210603", "title": "Test", "works": [{"content": "test1", "done": false}, {"content": "test2", "done": true}]}}`)

var todoItem = todo.TodoItem{}

func TestAddItem(t *testing.T) {
	storage := LocalStorage{}
	storage.InitLocalStorage()

	json.Unmarshal([]byte(todoItemJson), &todoItem)

	id, err := storage.AddItem(todoItem)
	if err != nil {
		t.Fatal(err.Error())
	}

	if storage.ItemLists.Len() != 1 {
		t.Fatal("Storage size invalid")
	}

	if storage.ItemLists.Len() != id+1 {
		t.Fatal("Invalid item id")
	}
}

func TestDeleteItem(t *testing.T) {
	storage := LocalStorage{}
	storage.InitLocalStorage()

	json.Unmarshal([]byte(todoItemJson), &todoItem)

	id, err := storage.AddItem(todoItem)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = storage.DeleteItem(id)
	if err != nil {
		t.Fatal(err.Error())
	}

	if storage.ItemLists.Len() != 0 {
		t.Fatal("Storage size invalid")
	}

	node := FindItem(storage.ItemLists, id)
	if node != nil {
		t.Fatal("find already deleted item")
	}
}

func TestGetItem(t *testing.T) {
	storage := LocalStorage{}
	storage.InitLocalStorage()

	json.Unmarshal([]byte(todoItemJson), &todoItem)

	id, err := storage.AddItem(todoItem)
	if err != nil {
		t.Fatal(err.Error())
	}

	item, err := storage.GetItem(id)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !reflect.DeepEqual(todoItem, *item) {
		t.Fatal("Invalid returned item")
	}
}

// func TestGetItemList(t *testing.T) {

// }
