package localstorage

import (
	"container/list"
)

type LocalStorage struct {
	Accounts  *list.List
	ItemLists *list.List
}

func (s *LocalStorage) InitLocalAccountStorage() error {
	s.Accounts = list.New()
	return nil
}

func (s *LocalStorage) InitLocalContentStorage() error {
	s.ItemLists = list.New()
	return nil
}
