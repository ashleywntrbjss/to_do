package repo

import "bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"

type repo interface {
	CreateItemFromTitle() todoitem.ToDoItem
	AddNew() todoitem.ToDoItem
	GetById() todoitem.ToDoItem
	GetAll() []todoitem.ToDoItem
	UpdateItemTitleById() error
	UpdateItemCompletionStatusById() error
	DeleteItemById() error
}
