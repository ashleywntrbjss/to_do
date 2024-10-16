package menu

import (
	"errors"
	"fmt"
	"strconv"
)

type Option struct {
	Key   string // Abbreviated name, e.g. deleteUser
	Title string // Friendly print name, e.g. "Delete user
}

type Menu struct {
	Title   string
	Options []Option
}

const ConsoleDecorateLine = "================================"

func (menu *Menu) PrintMenuItems() {
	fmt.Println(ConsoleDecorateLine)
	fmt.Println(menu.Title)
	fmt.Println(ConsoleDecorateLine)
	for index, option := range menu.Options {
		fmt.Println(index, "-", option.Title)
	}
	fmt.Println(ConsoleDecorateLine)
}

func (menu *Menu) ParseMenuSelectionString(selectionInput string) (string, error) {
	selectionAsInt, err := strconv.Atoi(selectionInput)
	if err != nil {
		return "", errors.New("input not a valid integer selection")
	}

	return menu.ParseMenuSelection(selectionAsInt)
}

func (menu *Menu) ParseMenuSelection(selectionIndex int) (string, error) {
	if len(menu.Options) < selectionIndex {
		return "", errors.New("selected index out of bounds")
	}

	return menu.Options[selectionIndex].Key, nil
}
