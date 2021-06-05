package handle

import (
	"encoding/json"
	"net/http"

	"github.com/ISSuh/go-todo/auth"
)

func (app *TodoHandle) Login(res http.ResponseWriter, req *http.Request) {
	user := auth.User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

}
