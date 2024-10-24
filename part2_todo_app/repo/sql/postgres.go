package sql

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log/slog"
)

type PostgresStore struct {
	db *pg.DB
}

func (r *PostgresStore) InitDB(ctx context.Context, connectionString string) error {
	opt, err := pg.ParseURL(connectionString)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return err
	}

	r.db = pg.Connect(opt)

	err = createSchema(r.db)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return err
	}
	slog.InfoContext(ctx, "Postgres db initialised")
	return nil
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*todoitem.ToDoItem)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{Temp: false, IfNotExists: true})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresStore) CreateItemFromTitle(ctx context.Context, title string) (todoitem.ToDoItem, error) {
	newItem := todoitem.NewToDoItem(title)
	result, err := r.db.Model(&newItem).Insert()

	if err != nil {
		return todoitem.ToDoItem{}, err
	}
	slog.InfoContext(ctx, fmt.Sprintf("Added item, rows affected: %v", result.RowsAffected()))
	return newItem, nil
}

func (r *PostgresStore) AddNew(ctx context.Context, item todoitem.ToDoItem) (int, error) {
	result, err := r.db.Model(&item).Insert()
	if err != nil {
		return -1, err
	}
	slog.InfoContext(ctx, fmt.Sprintf("Added item, rows affected: %v", result.RowsAffected()))
	return item.Id, nil
}
func (r *PostgresStore) GetById(ctx context.Context, itemId int) (todoitem.ToDoItem, error) {
	var item todoitem.ToDoItem
	err := r.db.Model(&item).Where("id = ?", itemId).Select()
	if err != nil {
		return todoitem.ToDoItem{}, err
	}
	slog.InfoContext(ctx, fmt.Sprintf("Found item by id: %v", item))
	return item, nil
}

func (r *PostgresStore) GetAll(ctx context.Context) ([]todoitem.ToDoItem, error) {
	var items []todoitem.ToDoItem
	err := r.db.Model(&items).Order("id DESC").Select()
	if err != nil {
		return nil, err
	}
	slog.InfoContext(ctx, "Finding all items")
	return items, nil
}

func (r *PostgresStore) UpdateItemTitleById(ctx context.Context, newTitle string, itemId int) error {
	result, err := r.db.Model(&todoitem.ToDoItem{}).Set("title = ?", newTitle).Where("id = ?", itemId).Update()
	if err != nil {
		return err
	}
	slog.InfoContext(ctx, fmt.Sprintf("Updating item title, rows affected: %v", result.RowsAffected()))
	return nil
}

func (r *PostgresStore) UpdateItemCompletionStatusById(ctx context.Context, completionStatus bool, itemId int) error {
	result, err := r.db.Model(&todoitem.ToDoItem{}).Set("is_complete = ?", completionStatus).Where("id = ?", itemId).Update()
	if err != nil {
		return err
	}

	slog.InfoContext(ctx, fmt.Sprintf("Updating item completion status, rows affected: %v", result.RowsAffected()))
	return err
}

func (r *PostgresStore) DeleteItemById(ctx context.Context, itemId int) error {
	result, err := r.db.Model(&todoitem.ToDoItem{}).Where("id = ?", itemId).Delete()
	if err != nil {
		return err
	}
	slog.InfoContext(ctx, fmt.Sprintf("Deleting item, rows affected: %v", result.RowsAffected()))
	return err
}
