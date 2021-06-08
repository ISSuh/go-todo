package localstorage

import (
	"container/list"
)

type LocalStorage struct {
	Accounts  *list.List
	Sessios   *list.List
	ItemLists *list.List
}

func (s *LocalStorage) InitAccountStorage() error {
	s.Accounts = list.New()
	return nil
}

func (s *LocalStorage) InitSessionStorage() error {
	s.Sessios = list.New()
	return nil
}

func (s *LocalStorage) InitContentStorage() error {
	s.ItemLists = list.New()
	return nil
}
