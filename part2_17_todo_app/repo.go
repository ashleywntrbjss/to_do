package main

import (
	"fmt"
	"sync"
)

var toDoItemRepo = []ToDoItem{
	{Id: 1, Title: "Washing up", IsComplete: true},
	{Id: 2, Title: "Ironing", IsComplete: false},
}

var repoLock = sync.Mutex{}

func AddItemFromTitle(title string) *ToDoItem {
	repoLock.Lock()
	defer repoLock.Unlock()

	newItem := NewToDoItem(title)
	newItem.Id = newIndex()
	toDoItemRepo = append(toDoItemRepo, *newItem)

	return newItem
}

func RemoveItemById(itemId int) {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := findIndexById(itemId)
	if !isFound {
		return
	}

	toDoItemRepo = append(toDoItemRepo[:index], toDoItemRepo[index+1:]...)
}

func UpdateItemTitleById(itemId int, newTitle string) {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := findIndexById(itemId)
	if !isFound {
		fmt.Println("Item not found")
		return
	}

	toDoItemRepo[index].Title = newTitle
}

func ToggleItemCompletionStatusById(itemId int) {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := findIndexById(itemId)
	if !isFound {
		fmt.Println("Item not found")
		return
	}

	toDoItemRepo[index].IsComplete = !toDoItemRepo[index].IsComplete
}

func GetById(itemId int) *ToDoItem {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := findIndexById(itemId)
	if !isFound {
		repoLock.Unlock()
		return nil
	}

	return &toDoItemRepo[index]
}

func GetAll() *[]ToDoItem {
	return &toDoItemRepo
}

func findIndexById(id int) (int, bool) {
	for index, item := range toDoItemRepo {
		if item.Id == id {
			return index, true
		}
	}
	fmt.Println("Provided item Id not found")
	return -1, false
}

func newIndex() int {
	return len(toDoItemRepo) + 1 // simple ID generation
}
