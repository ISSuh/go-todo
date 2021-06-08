package todo

import (
	"net/http"

	"github.com/ISSuh/go-todo/auth"
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
	app.Sessions = localStorage
	app.Contents = localStorage

	app.Accounts.InitAccountStorage()
	app.Sessions.InitSessionStorage()
	app.Contents.InitContentStorage()

	account := db.Account{
		User: db.User{
			Email:    "test",
			Password: "test",
		},
		AccountNum: 0,
	}

	app.Accounts.CreateAccount(account)
}

func extractTokenString(r *http.Request) (string, string) {
	var err error = nil

	accessTokenCookie, err := r.Cookie("access_token")
	if err != nil {
		return "", ""
	}

	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		return "", ""
	}

	return accessTokenCookie.Value, refreshTokenCookie.Value
}

func updateAccessToken(session *db.Session) error {
	var err error = nil
	session.Token.AccessToken, err = auth.CreateAccessToken(session.User.Email)
	if err != nil {
		return err
	}
	return nil
}

func updateTokenPair(session *db.Session) error {
	var err error = nil

	if auth.ValidateToken(session.Token.RefreshToken) != nil {
		session.Token.RefreshToken, err = auth.CreateAccessToken(session.User.Email)
		if err != nil {
			return err
		}
	}

	return updateAccessToken(session)
}
