package main

import (
	"encoding/json"
	"log"
	"os"
)

type ToDoItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func marshallToDoItems(args ...ToDoItem) []byte {
	toDoItems := []ToDoItem{}

	toDoItems = append(toDoItems, args...)
	marshalledItems, _ := json.Marshal(toDoItems)
	return marshalledItems
}

func bytesToFile(encodedData []byte) {

	err := os.WriteFile("file.txt", encodedData, 0666)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	toDoItems := []ToDoItem{{1, "Washing up"}, {2, "Ironing"}, {3, "Food shop"}, {4, "Wash car"}, {5, "Buy fish and chips"}, {6, "Renew car insurance"}, {7, "Complete Go tasks"}, {8, "return parcel"}, {9, "Organise cutlery drawer"}, {10, "Dust the mantle"}}

	marshalledToDoItems := marshallToDoItems(toDoItems...)

	bytesToFile(marshalledToDoItems)
}
