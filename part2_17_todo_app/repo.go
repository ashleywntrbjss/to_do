package main

import (
	"fmt"
	"sync"
)

var toDoItemRepo []ToDoItem

var repoLock = sync.Mutex{}

func AddItemFromTitle(title string) {
	newItem := NewToDoItem(title)
	addItem(*newItem)
}

func addItem(item ToDoItem) {
	repoLock.Lock()
	defer repoLock.Unlock()

	toDoItemRepo = append(toDoItemRepo, item)
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
