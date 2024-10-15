package repo

import (
	"bjss.com/ashley.winter/to_do/part2_17_todo_app_cmdline/todoitem"
	"errors"
	"fmt"
	"sync"
)

var toDoItemRepo = []todoitem.ToDoItem{
	{Id: 1, Title: "Washing up", IsComplete: true},
	{Id: 2, Title: "Ironing", IsComplete: false},
}

var repoLock = sync.Mutex{}

func AddItemFromTitle(title string) todoitem.ToDoItem {
	repoLock.Lock()
	defer repoLock.Unlock()

	newItem := todoitem.NewToDoItem(title)
	newItem.Id = newIndex()
	toDoItemRepo = append(toDoItemRepo, newItem)

	return newItem
}

func RemoveItemById(itemId int) {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := findIndexById(itemId)
	if !isFound {
		return
	}

	toDoItemRepo = append(toDoItemRepo[:index], toDoItemRepo[index+1:]...)
}

func UpdateItemTitleById(itemId int, newTitle string) {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := findIndexById(itemId)
	if !isFound {
		fmt.Println("Item not found")
		return
	}

	toDoItemRepo[index].Title = newTitle
}

func ToggleItemCompletionStatusById(itemId int) {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := findIndexById(itemId)
	if !isFound {
		fmt.Println("Item not found")
		return
	}

	toDoItemRepo[index].IsComplete = !toDoItemRepo[index].IsComplete
}

func GetById(itemId int) (todoitem.ToDoItem, error) {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := findIndexById(itemId)

	if !isFound {
		return todoitem.ToDoItem{}, errors.New("cannot find item by provided id")
	}

	returnItem := toDoItemRepo[index]

	return returnItem, nil
}

func findIndexById(id int) (int, bool) {
	for index, item := range toDoItemRepo {
		if item.Id == id {
			return index, true
		}
	}
	fmt.Println("Provided item Id not found")
	return -1, false
}

func GetAll() []todoitem.ToDoItem {
	return toDoItemRepo
}

func newIndex() int {
	return len(toDoItemRepo) + 1 // simple ID generation
}
