package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ToDoItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func parseTxtFile(fileName string) ([]byte, error) {
	file, err := os.ReadFile(fileName)

	return file, err
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

func mainMenu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("To Do - JSON Read and Print")
	fmt.Println("Please enter a file location: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	trimmedInput := strings.TrimSpace(input)

	parsedText, parseErr := parseTxtFile(trimmedInput)
	if parseErr != nil {
		fmt.Println("The file cannot be located. Please try again", parseErr)
		return
	}

	toDoItems := bytesToToDoItem(parsedText)

	printToDoItems(toDoItems)
}

func main() {
	for {
		mainMenu()
	}
}
