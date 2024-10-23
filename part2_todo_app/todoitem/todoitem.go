package todoitem

import "fmt"

type ToDoItem struct {
	tableName  struct{} `pg:"to_do_items"`
	Id         int      `json:"id"`
	Title      string   `json:"title"`
	IsComplete bool     `json:"isComplete"`
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
