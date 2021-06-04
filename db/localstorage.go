package db

import (
	"container/list"
	"errors"
	"fmt"

	"github.com/ISSuh/go-todo/todo"
)

type LocalStorage struct {
	ItemLists *list.List
}

func (s *LocalStorage) InitLocalStorage() error {
	s.ItemLists = list.New()
	return nil
}

func (s *LocalStorage) GetItem(itemId int) (*todo.TodoItem, error) {
	node := findItem(s.ItemLists, itemId)
	if node == nil {
		return nil, errors.New("Invalid Item id")
	}

	todoItem := &todo.TodoItem{}
	*todoItem = node.Value.(todo.TodoItem)
	return todoItem, nil
}

func (s *LocalStorage) GetItemList() (todo.TodoItemList, error) {
	var todoItems []todo.TodoItem
	for e := s.ItemLists.Front(); e != nil; e = e.Next() {
		todoItems = append(todoItems, e.Value.(todo.TodoItem))
	}

	return todo.TodoItemList{List: todoItems}, nil
}

func (s *LocalStorage) AddItem(item todo.TodoItem) (int, error) {
	item.Id = s.ItemLists.Len()
	s.ItemLists.PushBack(item)
	return index, nil
}

func (s *LocalStorage) DeleteItem(itemId int) error {
	node := findItem(s.ItemLists, itemId)
	if node == nil {
		return errors.New("Invalid Item id")
	}

	s.ItemLists.Remove(node)

	return nil
}

func findItem(itemList *list.List, id int) *list.Element {
	var node *list.Element
	node = nil

	for e := itemList.Front(); e != nil; e = e.Next() {
		todoItem := e.Value.(todo.TodoItem)
		if todoItem.Id == id {
			node = e
		}
	}

	return node
}

func printItems(itemList *list.List) {
	fmt.Println("---------------------------------")
	for e := itemList.Front(); e != nil; e = e.Next() {
		todoItem := e.Value.(todo.TodoItem)
		fmt.Println(todoItem)
	}
}
