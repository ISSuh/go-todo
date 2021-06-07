package todo

import (
	"github.com/ISSuh/go-todo/db"
	"github.com/ISSuh/go-todo/db/localstorage"
)

type Result struct {
	Status int    `json: "status"`
	Id     int    `json: "id, omitempty"`
	Err    string `json: "err, omitempty"`
}

type TodoHandle struct {
	Accounts db.AccountStorage
	Sessions db.SessionStorage
	Contents db.ContentStorage
}

func (app *TodoHandle) Initialize() {
	localStorage := &localstorage.LocalStorage{}

	app.Accounts = localStorage
	// app.Sessions = localStorage
	app.Contents = localStorage

	app.Accounts.InitLocalAccountStorage()
	// app.Sessions.InitLocalSessiontStorage()
	app.Contents.InitLocalContentStorage()

	account := db.Account{
		User: db.User{
			Email:    "test",
			Password: "test",
		},
		AccountNum: 0,
	}

	app.Accounts.CreateAccount(account)
}
