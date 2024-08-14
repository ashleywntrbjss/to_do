package part1_11

import (
    "fmt"
    "encoding/json"
)

type ToDoItem struct{
	Id int `json:"id"`
	Title string `json:"title"`
}

func ToDoItemsAsJson(arg ...ToDoItem){
	toDoItems := []ToDoItem{}

	toDoItems = append(toDoItems, arg...)
	marshalledItems, _ := json.Marshal(toDoItems)
	fmt.Println(string(marshalledItems))
}