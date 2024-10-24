package repo

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo/inMemory"
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo/sql"
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"context"
	"flag"
)

type Repo interface {
	CreateItemFromTitle(ctx context.Context, title string) (todoitem.ToDoItem, error)
	AddNew(ctx context.Context, item todoitem.ToDoItem) (int, error)
	GetById(ctx context.Context, itemId int) (todoitem.ToDoItem, error)
	GetAll(ctx context.Context) ([]todoitem.ToDoItem, error)
	UpdateItemTitleById(ctx context.Context, title string, itemId int) error
	UpdateItemCompletionStatusById(ctx context.Context, completionStatus bool, itemId int) error
	DeleteItemById(ctx context.Context, itemId int) error
}

func InitRepo(ctx context.Context) Repo {
	var repoType string

	var connectionString string

	var sharedStore Repo

	flag.StringVar(&repoType, "r", "memory", "type of repository")
	flag.StringVar(&connectionString, "cs", "", "connection string for postgres db")

	flag.Parse()

	switch repoType {
	case "memory":
		inMemoryStore := new(inMemory.InMemory)
		sharedStore = inMemoryStore

	case "sql":
		if connectionString == "" {
			panic("connectionString is required")
		}

		dbStore := new(sql.PostgresStore)
		err := dbStore.InitDB(ctx, connectionString)
		if err != nil {
			panic(err.Error())
		}

		sharedStore = dbStore
	default:
		sharedStore = new(inMemory.InMemory)
	}
	return sharedStore
}
