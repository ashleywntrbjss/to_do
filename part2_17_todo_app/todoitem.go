package main

type ToDoItem struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"is_complete"`
}

func NewToDoItem(title string) *ToDoItem {
	return &ToDoItem{Title: title, IsComplete: false}
}
