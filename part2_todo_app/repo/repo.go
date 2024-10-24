package repo

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo/inMemory"
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo/sql"
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"flag"
)

type Repo interface {
	CreateItemFromTitle(title string) (todoitem.ToDoItem, error)
	AddNew(item todoitem.ToDoItem) (int, error)
	GetById(itemId int) (todoitem.ToDoItem, error)
	GetAll() ([]todoitem.ToDoItem, error)
	UpdateItemTitleById(title string, itemId int) error
	UpdateItemCompletionStatusById(completionStatus bool, itemId int) error
	DeleteItemById(itemId int) error
}

func InitRepo() Repo {
	var repoType string

	var connectionString string

	var sharedStore Repo

	flag.StringVar(&repoType, "r", "memory", "type of repository")
	flag.StringVar(&connectionString, "cs", "", "connection string for postgres db")

	flag.Parse()

	switch repoType {
	case "memory":
		sharedStore = new(inMemory.InMemory)
	case "sql":
		if connectionString == "" {
			panic("connectionString is required")
		}

		dbStore := new(sql.PostgresStore)
		err := dbStore.InitDB(connectionString)
		if err != nil {
			panic(err.Error())
		}

		sharedStore = dbStore
	default:
		sharedStore = new(inMemory.InMemory)
	}
	return sharedStore
}
