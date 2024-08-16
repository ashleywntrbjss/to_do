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

func EditItemById(itemId int) {

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
