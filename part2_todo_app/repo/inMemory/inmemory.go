package inMemory

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"context"
	"errors"
	"fmt"
	"sync"
)

var (
	NotFoundError      = errors.New("entity not found")
	AlreadyExistsError = errors.New("entity already exists")
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

func (r *InMemory) CreateItemFromTitle(ctx context.Context, title string) (todoitem.ToDoItem, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	newItem := todoitem.NewToDoItem(title)
	newItem.Id = r.newIndex()
	r.store = append(r.store, newItem)

	return newItem, nil
}

func (r *InMemory) AddNew(ctx context.Context, item todoitem.ToDoItem) (int, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	item.Id = r.newIndex()

	_, isFound := r.findIndexById(ctx, item.Id)

	if isFound {
		fmt.Println("Item already exists")
		return -1, AlreadyExistsError
	}

	r.store = append(r.store, item)

	fmt.Println("Created new item", item)

	return item.Id, nil
}

func (r *InMemory) GetById(ctx context.Context, itemId int) (todoitem.ToDoItem, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	index, isFound := r.findIndexById(ctx, itemId)

	if !isFound {
		return todoitem.ToDoItem{}, NotFoundError
	}

	returnItem := r.store[index]

	return returnItem, nil
}

func (r *InMemory) GetAll(ctx context.Context) ([]todoitem.ToDoItem, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.store, nil
}

func (r *InMemory) UpdateItemTitleById(ctx context.Context, newTitle string, itemId int) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	index, isFound := r.findIndexById(ctx, itemId)
	if !isFound {
		return NotFoundError
	}

	r.store[index].Title = newTitle
	return nil
}

func (r *InMemory) UpdateItemCompletionStatusById(ctx context.Context, completionStatus bool, itemId int) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	index, isFound := r.findIndexById(ctx, itemId)
	if !isFound {
		fmt.Println("item not found")
		return NotFoundError
	}

	r.store[index].IsComplete = completionStatus
	return nil
}

func (r *InMemory) DeleteItemById(ctx context.Context, itemId int) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	index, isFound := r.findIndexById(ctx, itemId)
	if !isFound {
		return NotFoundError
	}

	r.store = append(r.store[:index], r.store[index+1:]...)
	return nil
}

func (r *InMemory) findIndexById(ctx context.Context, id int) (int, bool) {
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
