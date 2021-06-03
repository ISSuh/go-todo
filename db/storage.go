package db

import (
	"github.com/ISSuh/go-todo/todo"
)

type Storage interface {
	InitLocalStorage() error
	GetItem(itemId int) (*todo.TodoItem, error)
	GetItemList() (todo.TodoItemList, error)
	AddItem(item todo.TodoItem) (int, error)
	DeleteItem(itemId int) error
}
