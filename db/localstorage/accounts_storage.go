package localstorage

import (
	"container/list"
	"errors"
	"math/rand"

	"github.com/ISSuh/go-todo/db"
)

func (s *LocalStorage) CreateAccount(account db.Account) error {
	account.AccountNum = rand.Intn(100)
	s.Accounts.PushBack(account)
	return nil
}

func (s *LocalStorage) DeleteAccount(account db.Account) error {
	node := findAccount(s.Accounts, account.User.Email)
	if node == nil {
		return errors.New("Invalid user account")
	}

	s.Accounts.Remove(node)
	return nil
}

func (s *LocalStorage) AccessAccount(user db.User) db.Account {
	for e := s.Accounts.Front(); e != nil; e = e.Next() {
		userAccount := e.Value.(db.Account)
		if user.Email == userAccount.User.Email && user.Password == userAccount.User.Password {
			return userAccount
		}
	}
	return db.Account{}
}

func findAccount(accounts *list.List, email string) *list.Element {
	var node *list.Element
	node = nil

	for e := accounts.Front(); e != nil; e = e.Next() {
		account := e.Value.(db.Account)
		if account.User.Email == email {
			node = e
		}
	}

	return node
}
