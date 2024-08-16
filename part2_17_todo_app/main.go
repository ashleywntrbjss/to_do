package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ConsoleDecorateLine = "================================"

var reader = bufio.NewReader(os.Stdin)

func menu() {

	printDecoratedTitle("Welcome to the To Do application")

	fmt.Println("Please select an option: ")
	fmt.Println("1. Create a new To Do item")
	fmt.Println("2. View To Do items")
	fmt.Println("3. Edit a To Do item")
	fmt.Println("4. Delete a To Do item")
	fmt.Println("5. Exit application")

	fmt.Println(ConsoleDecorateLine)

	mainMenuSelection := readAndTrimUserInput("Select a menu item: ")

	handleMainMenuSelection(mainMenuSelection)
}

func handleMainMenuSelection(userInput string) {
	userInputAsInt, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Please enter a number selection")
		return
	}

	switch {
	case userInputAsInt == 1:
		handleCreateNewItem()

	case userInputAsInt == 2:
		handleViewItem()

	case userInputAsInt == 3:
		handleEditItem()

	case userInputAsInt == 4:
		handleDeleteItem()

	case userInputAsInt == 5:
		fmt.Println("Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("Please enter a valid option")
	}

}

func handleCreateNewItem() {
	printDecoratedTitle("Create a new To Do item")

	fmt.Println("Please input details for your new To Do item")
	itemName := readAndTrimUserInput("Item name: ")

	AddItemFromTitle(itemName)

}

func handleViewItem() {
	printDecoratedTitle("View To Do items")
	fmt.Println(GetAll())
}

func handleEditItem() {

}

func handleDeleteItem() {

}

func main() {
	for {
		menu()
	}
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

func printDecoratedTitle(title string) {
	fmt.Println(ConsoleDecorateLine)
	fmt.Println(title)
	fmt.Println(ConsoleDecorateLine)
}
