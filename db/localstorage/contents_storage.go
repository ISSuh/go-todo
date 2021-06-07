package localstorage

import (
	"container/list"
	"errors"

	"github.com/ISSuh/go-todo/db"
)

func (s *LocalStorage) GetItem(itemId int) (*db.TodoItem, error) {
	node := findContents(s.ItemLists, itemId)
	if node == nil {
		return nil, errors.New("Invalid Item id")
	}

	todoItem := &db.TodoItem{}
	*todoItem = node.Value.(db.TodoItem)
	return todoItem, nil
}

func (s *LocalStorage) GetItemList() (db.TodoItemList, error) {
	var todoItems []db.TodoItem

	for e := s.ItemLists.Front(); e != nil; e = e.Next() {
		todoItems = append(todoItems, e.Value.(db.TodoItem))
	}

	return db.TodoItemList{List: todoItems}, nil
}

func (s *LocalStorage) AddItem(item db.TodoItem) (int, error) {
	item.Id = s.ItemLists.Len()
	s.ItemLists.PushBack(item)
	return item.Id, nil
}

func (s *LocalStorage) DeleteItem(itemId int) error {
	node := findContents(s.ItemLists, itemId)
	if node == nil {
		return errors.New("Invalid Item id")
	}

	s.ItemLists.Remove(node)
	return nil
}

func findContents(itemList *list.List, id int) *list.Element {
	var node *list.Element
	node = nil

	for e := itemList.Front(); e != nil; e = e.Next() {
		todoItem := e.Value.(db.TodoItem)
		if todoItem.Id == id {
			node = e
		}
	}

	return node
}
