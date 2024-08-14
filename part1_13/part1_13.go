package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ToDoItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func parseTxtFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func bytesToToDoItem(bytes []byte) []ToDoItem {
	toDoItems := []ToDoItem{}
	err := json.Unmarshal(bytes, &toDoItems)
	if err != nil {
		log.Fatal(err)
	}
	return toDoItems
}

func printToDoItems(toDoItems []ToDoItem) {
	fmt.Println("To do:")
	for _, item := range toDoItems {
		fmt.Printf("%v: %v\n", strconv.Itoa(item.Id), item.Title)
	}
}

func main() {
	fileName := "./part1_13/toDoItemJson.txt"

	parsedText := parseTxtFile(fileName)

	toDoItems := bytesToToDoItem(parsedText)

	printToDoItems(toDoItems)

}
