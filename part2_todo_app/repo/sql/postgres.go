package sql

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type PostgresStore struct {
	db *pg.DB
}

func (r *PostgresStore) InitDB(connectionString string) error {
	opt, err := pg.ParseURL(connectionString)
	if err != nil {
		return err
	}

	r.db = pg.Connect(opt)

	err = createSchema(r.db)
	if err != nil {
		return err
	}
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

func (r *PostgresStore) CreateItemFromTitle(title string) (todoitem.ToDoItem, error) {
	newItem := todoitem.NewToDoItem(title)
	result, err := r.db.Model(&newItem).Insert()

	if err != nil {
		return todoitem.ToDoItem{}, err
	}
	fmt.Println("Added item, rows affected:", result.RowsAffected())
	return newItem, nil
}

func (r *PostgresStore) AddNew(item todoitem.ToDoItem) (int, error) {
	result, err := r.db.Model(&item).Insert()
	if err != nil {
		return -1, err
	}
	fmt.Println("Added item, rows affected:", result.RowsAffected())
	return item.Id, nil
}
func (r *PostgresStore) GetById(itemId int) (todoitem.ToDoItem, error) {
	var item todoitem.ToDoItem
	err := r.db.Model(&item).Where("id = ?", itemId).Select()
	if err != nil {
		return todoitem.ToDoItem{}, err
	}
	fmt.Println("Found item by id:", item)
	return item, nil
}

func (r *PostgresStore) GetAll() ([]todoitem.ToDoItem, error) {
	var items []todoitem.ToDoItem
	err := r.db.Model(&items).Order("id DESC").Select()
	if err != nil {
		return nil, err
	}
	fmt.Println("Finding all items")
	return items, nil
}

func (r *PostgresStore) UpdateItemTitleById(newTitle string, itemId int) error {
	result, err := r.db.Model(&todoitem.ToDoItem{}).Set("title = ?", newTitle).Where("id = ?", itemId).Update()
	if err != nil {
		return err
	}
	fmt.Println("Updating item title, rows affected:", result.RowsAffected())
	return nil
}

func (r *PostgresStore) UpdateItemCompletionStatusById(completionStatus bool, itemId int) error {
	result, err := r.db.Model(&todoitem.ToDoItem{}).Set("is_complete = ?", completionStatus).Where("id = ?", itemId).Update()
	if err != nil {
		return err
	}

	fmt.Println("Updating item completion status, rows affected:", result.RowsAffected())
	return err
}

func (r *PostgresStore) DeleteItemById(itemId int) error {
	result, err := r.db.Model(&todoitem.ToDoItem{}).Where("id = ?", itemId).Delete()
	if err != nil {
		return err
	}
	fmt.Println("Deleting item, rows affected:", result.RowsAffected())
	return err
}
