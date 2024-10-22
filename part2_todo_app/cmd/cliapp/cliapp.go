package cliapp

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/cmd/menu"
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var ConsoleDecorateLine = "================================"

var reader = bufio.NewReader(os.Stdin)

var activeRepo repo.Repo

func RunCli(repo repo.Repo) {

	activeRepo = repo

	if activeRepo == nil {
		panic(" repo not initalised ")
	}

	printDecoratedTitle("Welcome to the To Do application")
	for {
		runMainMenu()
	}
}

func runMainMenu() {

	mainMenu := menu.Menu{
		Title: "Main menu",
		Options: []menu.Option{
			{Key: "create", Title: "Create a new To Do item"},
			{Key: "viewAll", Title: "View all To Do items"},
			{Key: "edit", Title: "Edit a To Do item"},
			{Key: "exit", Title: "Exit To Do app"},
		},
	}

	mainMenu.PrintMenuItems()

	mainMenuSelection := readAndTrimUserInput("Select a menu item")

	selectionKey, err := mainMenu.ParseMenuSelectionString(mainMenuSelection)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch selectionKey {
	case "create":
		handleCreateNewItem()

	case "viewAll":
		handleViewItem()

	case "edit":
		handleEditItem()

	case "exit":
		fmt.Println("Goodbye!")
		os.Exit(0)

	default:
		fmt.Println("Please select a valid option")
	}

}

func handleCreateNewItem() {
	printDecoratedTitle("Create a new To Do item")

	fmt.Println("Please input details for your new To Do item")
	itemName := readAndTrimUserInput("Item name")

	newItem := activeRepo.CreateItemFromTitle(itemName)

	fmt.Printf("Added your item:")
	newItem.PrettyPrintToDoItem()
	fmt.Println()

	pauseForInput()
}

func handleViewItem() {
	printDecoratedTitle("View To Do items")
	todoitem.PrettyPrintToDoItems(activeRepo.GetAll()...)
	pauseForInput()
}

func handleEditItem() {
	printDecoratedTitle("Edit To Do items")

	todoitem.PrettyPrintToDoItems(activeRepo.GetAll()...)

	userInput := readAndTrimUserInput("Provide the Id of the item to edit")

	userInputAsInt, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Please enter a number selection")
		return
	}

	activeItem, err := activeRepo.GetById(userInputAsInt)

	if err != nil {
		fmt.Println("Unable to retrieve item from repo", err)
		return
	}

	handleSelectedItem(activeItem)

}

func handleSelectedItem(activeItem todoitem.ToDoItem) {
	fmt.Println("Selected To Do Item: ")
	activeItem.PrettyPrintToDoItem()

	var markUnmarkPrompt string

	markAsCompletePrompt := "Mark as complete"
	markAsIncompletePrompt := "Mark as incomplete"

	if activeItem.IsComplete {
		markUnmarkPrompt = markAsIncompletePrompt
	} else {
		markUnmarkPrompt = markAsCompletePrompt
	}

	editMenu := menu.Menu{
		Title: "Edit To Do item menu",
		Options: []menu.Option{
			{Key: "markCompletionStatus", Title: markUnmarkPrompt},
			{Key: "updateTitle", Title: "Update title"},
			{Key: "deleteItem", Title: "Delete To Do item"},
			{Key: "exit", Title: "Exit to main menu"},
		},
	}

	editMenu.PrintMenuItems()
	editSelection := readAndTrimUserInput("Select an edit option: ")

	editSelectionKey, err := editMenu.ParseMenuSelectionString(editSelection)
	if err != nil {
		fmt.Println("Error when handling selection", err)
		return
	}

	switch editSelectionKey {
	case "markCompletionStatus":
		activeRepo.UpdateItemCompletionStatusById(!activeItem.IsComplete, activeItem.Id)
	case "updateTitle":
		prompt := "Provide new title for item: '" + activeItem.Title + "'"
		newTitle := readAndTrimUserInput(prompt)
		activeRepo.UpdateItemTitleById(newTitle, activeItem.Id)
	case "deleteItem":
		prompt := "Are you sure you wish to delete: '" + activeItem.Title + "'? (yes/no)"
		decision := readAndTrimUserInput(prompt)

		if decision == "yes" {
			activeRepo.DeleteItemById(activeItem.Id)
			fmt.Println("To Do item deleted")
			return // skip printing logic for deleted item
		}
		fmt.Println("To Do item not deleted")
		return
	case "exit":
		return
	default:
		fmt.Println("Please enter a valid option")
		return
	}

	fmt.Println("Updated item: ")

	activeItem, err = activeRepo.GetById(activeItem.Id)
	if err != nil {
		log.Fatal("Unable to retrieve recently updated item", err)
	}

	activeItem.PrettyPrintToDoItem()
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
