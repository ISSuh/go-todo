package handle

import (
	"github.com/ISSuh/go-todo/db"
)

type Result struct {
	Status int    `json: "status"`
	Id     int    `json: id, omitempty`
	Err    string `json: "err, omitempty"`
}

type TodoHandle struct {
	Account  db.UserStorage
	Contents db.ContentStorage
}

func (app *TodoHandle) Initialize() {
	localStorage := &db.LocalStorage{}
	localStorage.InitLocalAccountStorage()
	localStorage.InitLocalContentStorage()

	app.Account = localStorage
	app.Contents = localStorage
}
