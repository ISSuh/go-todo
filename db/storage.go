package db

import (
	"github.com/ISSuh/go-todo/todo"
)

// Storage for user account
type AccountStorage interface {
	InitLocalAccountStorage()

	CreateAccount(user auth.User) error
	DeleteAccount(user auth.User) error
	AccessAccount(user auth.User) error
}

// Storage for active session 
type SessionStorage interface() {
	InitLocalSessiontStorage()

	AddSession(session auth.Session) error
	DeleteSession(user auth.User) error
	FindSession(user auth.User) error
}

// 
type ContentStorage interface {
	InitLocalContentStorage() error
	GetItem(itemId int) (*todo.TodoItem, error)
	GetItemList() (todo.TodoItemList, error)
	AddItem(item todo.TodoItem) (int, error)
	DeleteItem(itemId int) error
}
