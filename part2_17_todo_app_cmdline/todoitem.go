package main

import "fmt"

type ToDoItem struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"is_complete"`
}

func NewToDoItem(title string) *ToDoItem {
	return &ToDoItem{Title: title, IsComplete: false}
}

func (item *ToDoItem) PrettyPrintToDoItem() {
	if item.IsComplete {
		fmt.Printf("%v - %v - [X]", item.Id, item.Title)
		return
	}
	fmt.Printf("\n%v - %v - [ ]", item.Id, item.Title)
}

func PrettyPrintToDoItems(items *[]ToDoItem) {
	for _, item := range *items {
		item.PrettyPrintToDoItem()
	}
	fmt.Println()
}
