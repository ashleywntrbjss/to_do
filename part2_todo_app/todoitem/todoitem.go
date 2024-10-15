package todoitem

import "fmt"

type ToDoItem struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"is_complete"`
}

func NewToDoItem(title string) ToDoItem {
	return ToDoItem{Title: title, IsComplete: false}
}

func (item ToDoItem) PrettyPrintToDoItem() {

	var completionIndicator string

	if item.IsComplete {
		completionIndicator = "[X]"
	} else {
		completionIndicator = "[ ]"
	}

	fmt.Println(item.Id, "-", item.Title, "-", completionIndicator)
}

func PrettyPrintToDoItems(item ...ToDoItem) {
	for _, item := range item {
		item.PrettyPrintToDoItem()
	}
}
