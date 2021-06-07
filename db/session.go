package db

import (
	"github.com/ISSuh/go-todo/auth"
)

type Session struct {
	User  User
	Token auth.Token
}
