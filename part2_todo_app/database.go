package main

import "fmt"

var toDoItemRepo []ToDoItem

func AddItem(item ToDoItem) {
	toDoItemRepo = append(toDoItemRepo, item)
}

func RemoveItemById(itemId int) {
	index, isFound := findIndexById(itemId)
	if !isFound {
		return
	}

	toDoItemRepo = append(toDoItemRepo[:index], toDoItemRepo[index+1:]...)
}

func UpdateItemTitleById(itemId int, newTitle string) {
	index, isFound := findIndexById(itemId)
	if !isFound {
		fmt.Println("Item not found")
		return
	}
	toDoItemRepo[index].Title = newTitle
}

func ToggleItemCompletionStatusById(itemId int) {
	index, isFound := findIndexById(itemId)
	if !isFound {
		fmt.Println("Item not found")
		return
	}
	toDoItemRepo[index].IsComplete = !toDoItemRepo[index].IsComplete
}

func GetById(itemId int) *ToDoItem {
	index, isFound := findIndexById(itemId)
	if !isFound {
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
