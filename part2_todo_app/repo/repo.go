package repo

import "bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"

type Repo interface {
	CreateItemFromTitle(title string) (todoitem.ToDoItem, error)
	AddNew(item todoitem.ToDoItem) (int, error)
	GetById(itemId int) (todoitem.ToDoItem, error)
	GetAll() ([]todoitem.ToDoItem, error)
	UpdateItemTitleById(title string, itemId int) error
	UpdateItemCompletionStatusById(completionStatus bool, itemId int) error
	DeleteItemById(itemId int) error
}
