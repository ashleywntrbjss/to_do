package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func menu() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(ConsoleDecorateLine)
	fmt.Println("Welcome to the To Do Application")
	fmt.Println(ConsoleDecorateLine)

	fmt.Println("Please select an option: ")
	fmt.Println("1. Create a new To Do item")
	fmt.Println("2. View To Do items")
	fmt.Println("3. Edit a To Do item")
	fmt.Println("4. Delete a To Do item")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	handleMainMenuSelection(input)
}

func handleMainMenuSelection(userInput string) {
	trimmedInput := strings.TrimSpace(userInput)

	trimmedInputAsInt, err := strconv.Atoi(trimmedInput)
	if err != nil {
		fmt.Println("Please enter a number selection", err)
		return
	}

	switch {
	case trimmedInputAsInt == 1:

	}
}

func Main() {
	for {
		menu()
	}
}
