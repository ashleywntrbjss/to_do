package repo

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"errors"
	"fmt"
	"sync"
)

var toDoItemRepo = []todoitem.ToDoItem{
	{Id: 1, Title: "Washing up", IsComplete: true},
	{Id: 2, Title: "Ironing", IsComplete: false},
}

var repoLock = sync.Mutex{}

func CreateItemFromTitle(title string) todoitem.ToDoItem {
	repoLock.Lock()
	defer repoLock.Unlock()

	newItem := todoitem.NewToDoItem(title)
	newItem.Id = newIndex()
	toDoItemRepo = append(toDoItemRepo, newItem)

	return newItem
}

func AddNew(item todoitem.ToDoItem) {
	repoLock.Lock()
	defer repoLock.Unlock()

	item.Id = newIndex()
	toDoItemRepo = append(toDoItemRepo, item)
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

func GetAll() []todoitem.ToDoItem {
	return toDoItemRepo
}

func UpdateItemTitleById(newTitle string, itemId int) error {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := findIndexById(itemId)
	if !isFound {
		return errors.New("cannot find item by provided id")
	}

	toDoItemRepo[index].Title = newTitle
	return nil
}

func UpdateItemCompletionStatusById(completionStatus bool, itemId int) error {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := findIndexById(itemId)
	if !isFound {
		fmt.Println("item not found")
		return errors.New("item not found")
	}

	toDoItemRepo[index].IsComplete = completionStatus
	return nil
}

func DeleteItemById(itemId int) {
	repoLock.Lock()
	defer repoLock.Unlock()

	index, isFound := findIndexById(itemId)
	if !isFound {
		return
	}

	toDoItemRepo = append(toDoItemRepo[:index], toDoItemRepo[index+1:]...)
}

func findIndexById(id int) (int, bool) {
	for index, item := range toDoItemRepo {
		if item.Id == id {
			return index, true
		}
	}
	fmt.Println("provided item Id not found")
	return -1, false
}

func newIndex() int {
	return len(toDoItemRepo) + 1 // simple ID generation
}
