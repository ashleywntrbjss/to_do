package menu

import (
	"errors"
	"fmt"
)

type Option struct {
	Key   string // Abbreviated name, e.g. deleteUser
	Title string // Friendly print name, e.g. "Delete user
}

type Menu struct {
	Options []Option
}

const ConsoleDecorateLine = "================================"

func (menu *Menu) PrintMenuItems() {
	fmt.Println(ConsoleDecorateLine)
	for index, option := range menu.Options {
		fmt.Println(index, "-", option.Title)
	}
	fmt.Println(ConsoleDecorateLine)
}

func (menu *Menu) MakeMenuSelection(selectionIndex int) (string, error) {
	if len(menu.Options) < selectionIndex {
		return "", errors.New("selected index out of bounds")
	}

	return menu.Options[selectionIndex].Key, nil
}
