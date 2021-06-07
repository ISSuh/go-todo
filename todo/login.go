package todo

import (
	"encoding/json"
	"net/http"

	"github.com/ISSuh/go-todo/auth"
	"github.com/ISSuh/go-todo/db"
)

func (app *TodoHandle) Login(res http.ResponseWriter, req *http.Request) {
	user := db.User{}

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	account := app.Accounts.AccessAccount(user)

	token, _ := auth.CreateToken(account.User.Email)
	session := db.Session{User: user, Token: *token}

	accessTokenCookie := http.Cookie{
		Name:     "access_token",
		Value:    session.Token.AccessToken,
		HttpOnly: true,
	}

	refreshTokenCookie := http.Cookie{
		Name:     "access_token",
		Value:    session.Token.RefreshToken,
		HttpOnly: true,
	}

	res.Header().Set("Set-Cookie", accessTokenCookie.String())
	res.Header().Add("Set-Cookie", refreshTokenCookie.String())

	result, _ := json.MarshalIndent(Result{Status: 200}, "", "  ")
	res.Write([]byte(result))
}
