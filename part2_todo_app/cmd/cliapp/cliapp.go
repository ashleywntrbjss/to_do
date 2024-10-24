package cliapp

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/cmd/menu"
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var ConsoleDecorateLine = "================================"

var reader = bufio.NewReader(os.Stdin)

var activeRepo repo.Repo

func RunCli(ctx context.Context, repo repo.Repo) {

	activeRepo = repo

	if activeRepo == nil {
		panic("repo not initialised")
	}

	printDecoratedTitle("Welcome to the To Do application")
	for {
		runMainMenu(ctx)
	}
}

func runMainMenu(ctx context.Context) {

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
		handleCreateNewItem(ctx)

	case "viewAll":
		handleViewItem(ctx)

	case "edit":
		handleEditItem(ctx)

	case "exit":
		fmt.Println("Goodbye!")
		os.Exit(0)

	default:
		fmt.Println("Please select a valid option")
	}

}

func handleCreateNewItem(ctx context.Context) {
	printDecoratedTitle("Create a new To Do item")

	fmt.Println("Please input details for your new To Do item")
	itemName := readAndTrimUserInput("Item name")

	newItem, err := activeRepo.CreateItemFromTitle(ctx, itemName)

	if err != nil {
		fmt.Println("error adding item", err)
	}

	fmt.Println("Added your item:")
	newItem.PrettyPrintToDoItem()
	fmt.Println()

	pauseForInput()
}

func handleViewItem(ctx context.Context) {
	printDecoratedTitle("View To Do items")
	items, err := activeRepo.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoitem.PrettyPrintToDoItems(items...)
	pauseForInput()
}

func handleEditItem(ctx context.Context) {
	printDecoratedTitle("Edit To Do items")

	items, err := activeRepo.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoitem.PrettyPrintToDoItems(items...)

	userInput := readAndTrimUserInput("Provide the Id of the item to edit")

	userInputAsInt, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Please enter a number selection")
		return
	}

	activeItem, err := activeRepo.GetById(ctx, userInputAsInt)

	if err != nil {
		fmt.Println("Unable to retrieve item from repo", err)
		return
	}

	handleSelectedItem(ctx, activeItem)

}

func handleSelectedItem(ctx context.Context, activeItem todoitem.ToDoItem) {
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
		err := activeRepo.UpdateItemCompletionStatusById(ctx, !activeItem.IsComplete, activeItem.Id)
		if err != nil {
			return
		}

	case "updateTitle":
		prompt := "Provide new title for item: '" + activeItem.Title + "'"
		newTitle := readAndTrimUserInput(prompt)
		err := activeRepo.UpdateItemTitleById(ctx, newTitle, activeItem.Id)
		if err != nil {
			return
		}
	case "deleteItem":
		prompt := "Are you sure you wish to delete: '" + activeItem.Title + "'? (yes/no)"
		decision := readAndTrimUserInput(prompt)

		if decision == "yes" {
			err := activeRepo.DeleteItemById(ctx, activeItem.Id)
			if err != nil {
				return
			}
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

	activeItem, err = activeRepo.GetById(ctx, activeItem.Id)
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
