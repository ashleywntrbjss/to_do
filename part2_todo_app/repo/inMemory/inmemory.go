package inMemory

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"errors"
	"fmt"
	"sync"
)

type InMemory struct {
	store []todoitem.ToDoItem
}

var repoLock = sync.Mutex{}

func (r *InMemory) CreateItemFromTitle(title string) todoitem.ToDoItem {
	repoLock.Lock()
	defer repoLock.Unlock()

	newItem := todoitem.NewToDoItem(title)
	newItem.Id = r.newIndex()
	r.store = append(r.store, newItem)

	return newItem
}

func (r *InMemory) AddNew(item todoitem.ToDoItem) (int, error) {
	repoLock.Lock()
	defer repoLock.Unlock()

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
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := r.findIndexById(itemId)

	if !isFound {
		return todoitem.ToDoItem{}, errors.New("cannot find item by provided id")
	}

	returnItem := r.store[index]

	return returnItem, nil
}

func (r *InMemory) GetAll() []todoitem.ToDoItem {
	return r.store
}

func (r *InMemory) UpdateItemTitleById(newTitle string, itemId int) error {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := r.findIndexById(itemId)
	if !isFound {
		return errors.New("cannot find item by provided id")
	}

	r.store[index].Title = newTitle
	return nil
}

func (r *InMemory) UpdateItemCompletionStatusById(completionStatus bool, itemId int) error {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := r.findIndexById(itemId)
	if !isFound {
		fmt.Println("item not found")
		return errors.New("item not found")
	}

	r.store[index].IsComplete = completionStatus
	return nil
}

func (r *InMemory) DeleteItemById(itemId int) error {
	repoLock.Lock()
	defer repoLock.Unlock()

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
