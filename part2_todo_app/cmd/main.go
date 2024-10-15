package main

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ConsoleDecorateLine = "================================"

var reader = bufio.NewReader(os.Stdin)

func main() {
	printDecoratedTitle("Welcome to the To Do application")
	for {
		menu()
	}
}

func menu() {

	printDecoratedTitle("Main menu")

	fmt.Println("Please select an option: ")
	fmt.Println("1. Create a new To Do item")
	fmt.Println("2. View To Do items")
	fmt.Println("3. Edit a To Do item")
	fmt.Println("4. Delete a To Do item")
	fmt.Println("5. Exit application")

	fmt.Println(ConsoleDecorateLine)

	mainMenuSelection := readAndTrimUserInput("Select a menu item")

	handleMainMenuSelection(mainMenuSelection)
}

func handleMainMenuSelection(userInput string) {

	switch userInput {
	case "1":
		handleCreateNewItem()

	case "2":
		handleViewItem()

	case "3":
		handleEditItem()

	case "4":
		handleDeleteItem()

	case "5":
		fmt.Println("Goodbye!")
		os.Exit(0)

	default:
		fmt.Println("Please enter a valid option")
	}

}

func handleCreateNewItem() {
	printDecoratedTitle("Create a new To Do item")

	fmt.Println("Please input details for your new To Do item")
	itemName := readAndTrimUserInput("Item name")

	newItem := repo.CreateItemFromTitle(itemName)

	fmt.Printf("Added your item:")
	newItem.PrettyPrintToDoItem()
	fmt.Println()

	pauseForInput()
}

func handleViewItem() {
	printDecoratedTitle("View To Do items")
	todoitem.PrettyPrintToDoItems(repo.GetAll()...)
	pauseForInput()
}

func handleEditItem() {
	printDecoratedTitle("Edit To Do items")
	todoitem.PrettyPrintToDoItems(repo.GetAll()...)

	userInput := readAndTrimUserInput("Provide the Id of the item to edit")

	userInputAsInt, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Please enter a number selection")
		return
	}

	activeItem, err := repo.GetById(userInputAsInt)

	if err != nil {
		fmt.Println("Unable to retrieve item from repo", err)
		return
	}

	fmt.Println("Selected To Do Item: ")
	activeItem.PrettyPrintToDoItem()

	var markUnmarkPrompt string

	markAsCompletePrompt := "1. Mark as complete"
	markAsIncompletePrompt := "2. Mark as incomplete"

	if activeItem.IsComplete {
		markUnmarkPrompt = markAsCompletePrompt
	} else {
		markUnmarkPrompt = markAsIncompletePrompt
	}

	fmt.Println(markUnmarkPrompt)
	fmt.Println("2. Update title")

	editOption := readAndTrimUserInput("Select an edit option: ")

	switch editOption {
	case "1":
		break
	case "2":
		break
	default:
		fmt.Println("Please enter a valid option")

	}

}

func handleDeleteItem() {

}

func readAndTrimUserInput(prompt string) string {
	fmt.Printf("%v: ", prompt)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again")
		return ""
	}

	trimmedInput := strings.TrimSpace(input)

	return trimmedInput
}

func pauseForInput() {
	fmt.Println("Press enter to continue...")
	_, _ = reader.ReadByte()
}

func printDecoratedTitle(title string) {
	fmt.Println(ConsoleDecorateLine)
	fmt.Println(title)
	fmt.Println(ConsoleDecorateLine)
}
