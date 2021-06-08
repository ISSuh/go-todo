package localstorage

import (
	"container/list"
	"errors"
	"fmt"

	"github.com/ISSuh/go-todo/db"
)

func (s *LocalStorage) AddSession(session db.Session) error {
	if findSession(s.Sessios, session.User.Email) != nil {
		return errors.New("Already exist")
	}

	s.Sessios.PushBack(session)

	printForDebuggingOnSession(s.Sessios)
	return nil
}

func (s *LocalStorage) DeleteSession(user db.User) error {
	node := findSession(s.Sessios, user.Email)
	if node == nil {
		return errors.New("Invalid user")
	}

	s.ItemLists.Remove(node)

	printForDebuggingOnSession(s.Sessios)
	return nil
}

func (s *LocalStorage) GetSession(user db.User) (*db.Session, error) {
	node := findSession(s.Sessios, user.Email)
	if node == nil {
		return nil, errors.New("Invalid user")
	}

	session := &db.Session{}
	*session = node.Value.(db.Session)
	return session, nil
}

func findSession(sessions *list.List, email string) *list.Element {
	var node *list.Element = nil

	for e := sessions.Front(); e != nil; e = e.Next() {
		session := e.Value.(db.Session)
		if session.User.Email == email {
			node = e
		}
	}
	return node
}

func printForDebuggingOnSession(sessions *list.List) {
	fmt.Println("-----------------------")
	for e := sessions.Front(); e != nil; e = e.Next() {
		session := e.Value.(db.Session)
		fmt.Println(session)
	}
}
