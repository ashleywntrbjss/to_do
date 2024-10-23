package inMemory

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"errors"
	"fmt"
	"sync"
)

type InMemory struct {
	store []todoitem.ToDoItem
	lock  sync.RWMutex
}

func (r *InMemory) InitTestData() {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.store = append(r.store, todoitem.ToDoItem{
		Id:         1,
		Title:      "Washing up",
		IsComplete: false,
	}, todoitem.ToDoItem{
		Id:         2,
		Title:      "Walk the dog",
		IsComplete: true,
	})
}

func (r *InMemory) CreateItemFromTitle(title string) (todoitem.ToDoItem, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	newItem := todoitem.NewToDoItem(title)
	newItem.Id = r.newIndex()
	r.store = append(r.store, newItem)

	return newItem, nil
}

func (r *InMemory) AddNew(item todoitem.ToDoItem) (int, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	item.Id = r.newIndex()

	_, isFound := r.findIndexById(item.Id)

	if isFound {
		fmt.Println("Item already exists")
		return -1, errors.New("item already exists")
	}

	r.store = append(r.store, item)

	fmt.Println("Created new item", item)

	return item.Id, nil
}

func (r *InMemory) GetById(itemId int) (todoitem.ToDoItem, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	index, isFound := r.findIndexById(itemId)

	if !isFound {
		return todoitem.ToDoItem{}, errors.New("cannot find item by provided id")
	}

	returnItem := r.store[index]

	return returnItem, nil
}

func (r *InMemory) GetAll() ([]todoitem.ToDoItem, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.store, nil
}

func (r *InMemory) UpdateItemTitleById(newTitle string, itemId int) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	index, isFound := r.findIndexById(itemId)
	if !isFound {
		return errors.New("cannot find item by provided id")
	}

	r.store[index].Title = newTitle
	return nil
}

func (r *InMemory) UpdateItemCompletionStatusById(completionStatus bool, itemId int) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	index, isFound := r.findIndexById(itemId)
	if !isFound {
		fmt.Println("item not found")
		return errors.New("item not found")
	}

	r.store[index].IsComplete = completionStatus
	return nil
}

func (r *InMemory) DeleteItemById(itemId int) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	index, isFound := r.findIndexById(itemId)
	if !isFound {
		return errors.New("unable to find item to delete")
	}

	r.store = append(r.store[:index], r.store[index+1:]...)
	return nil
}

func (r *InMemory) findIndexById(id int) (int, bool) {
	for index, item := range r.store {
		if item.Id == id {
			return index, true
		}
	}
	return -1, false
}

func (r *InMemory) newIndex() int {
	return len(r.store) + 1 // simple ID generation
}
