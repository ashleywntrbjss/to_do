package sql

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"errors"
	"github.com/go-pg/pg/v10"
)

type PostgresStore struct {
	db *pg.DB
}

func (r *PostgresStore) InitDB() {
	r.db = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "1234",
	})
}

func (r *PostgresStore) CreateItemFromTitle(title string) (todoitem.ToDoItem, error) {
	newItem := todoitem.NewToDoItem(title)
	_, err := r.db.Model(newItem).Insert()
	if err != nil {
		return todoitem.ToDoItem{}, err
	}
	return newItem, nil
}

func (r *PostgresStore) AddNew(item todoitem.ToDoItem) (int, error) {
	_, err := r.db.Model(item).Insert()
	if err != nil {
		return -1, err
	}
	return item.Id, nil
}
func (r *PostgresStore) GetById(itemId int) (todoitem.ToDoItem, error) {
	var item todoitem.ToDoItem
	err := r.db.Model(item).Where("id = ?", itemId).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return todoitem.ToDoItem{}, errors.New("cannot find item by provided id")
		}
		return todoitem.ToDoItem{}, err
	}
	return item, nil
}

func (r *PostgresStore) GetAll() ([]todoitem.ToDoItem, error) {
	var items []todoitem.ToDoItem
	err := r.db.Model(items).Select()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *PostgresStore) UpdateItemTitleById(newTitle string, itemId int) error {
	_, err := r.db.Model(todoitem.ToDoItem{}).Set("title = ?", newTitle).Where("id = ?", itemId).Update()
	return err
}

func (r *PostgresStore) UpdateItemCompletionStatusById(completionStatus bool, itemId int) error {
	_, err := r.db.Model(todoitem.ToDoItem{}).Set("is_complete = ?", completionStatus).Where("id = ?", itemId).Update()
	return err
}

func (r *PostgresStore) DeleteItemById(itemId int) error {
	_, err := r.db.Model(&todoitem.ToDoItem{}).Where("id = ?", itemId).Delete()
	return err
}
