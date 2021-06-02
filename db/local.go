package db

import (
	"github.com/ISSuh/go-todo/todo"
)

type LocalStorage struct {
	Items []]todo.Item
}

func (s *LocalStorage) AddItem(item Item) (int32, error) {
	index := len(s.Items)
	s.Items = append(s.Items, item)
	return index, nil
}

func (s *LocalStorage) DeleteItem(itemId int32) error {
	if itemId >= len(s.Items) {
		return error.new("")
	}

	index := len(s.Items)
	s.Items = append(s.Items, item)
	return index, nil
}

func (s *LocalStorage) GetItem(itemId int32) (Item, error) {

}

func (s *LocalStorage) GetItemList() (ItemList, error) {

}
