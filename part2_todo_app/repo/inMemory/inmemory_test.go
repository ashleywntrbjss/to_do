package inMemory

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"sync"
	"testing"
)

func TestAddItemFromTitle(t *testing.T) {
	testRepo := InMemory{
		store: []todoitem.ToDoItem{},
		lock:  sync.RWMutex{},
	}

	title := "Task 1"
	item := testRepo.CreateItemFromTitle(title)

	if len(testRepo.store) != 1 {
		t.Errorf("Expected 1 item, got %d", len(testRepo.store))
	}
	if testRepo.store[0].Title != title {
		t.Errorf("Expected item title %v, got %v", title, testRepo.store[0].Title)
	}
	if testRepo.store[0].Id != item.Id {
		t.Errorf("Expected item ID %v, got %v", item.Id, testRepo.store[0].Id)
	}
}

func TestRemoveItemById(t *testing.T) {
	testRepo := InMemory{
		store: []todoitem.ToDoItem{
			{Id: 1, Title: "Task 1", IsComplete: false},
			{Id: 2, Title: "Task 2", IsComplete: false}},
		lock: sync.RWMutex{},
	}

	err := testRepo.DeleteItemById(2)
	if err != nil {
		t.Fatalf("Error deleting item: %v", err)
	}
	if len(testRepo.store) != 1 {
		t.Errorf("Expected 1 item, got %d", len(testRepo.store))
	}
	if testRepo.store[0].Id != 1 {
		t.Errorf("Expected item with ID 1, got %d", testRepo.store[0].Id)
	}
}

func TestUpdateItemTitleById(t *testing.T) {
	testRepo := InMemory{
		store: []todoitem.ToDoItem{
			{Id: 1, Title: "Task 1", IsComplete: false},
			{Id: 2, Title: "Task 2", IsComplete: false}},
		lock: sync.RWMutex{},
	}

	err := testRepo.UpdateItemTitleById("Updated Task 2", 2)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if testRepo.store[1].Title != "Updated Task 2" {
		t.Errorf("Expected title 'Updated Task 2', got %s", testRepo.store[1].Title)
	}
}

func TestToggleCompletionById(t *testing.T) {
	testRepo := InMemory{
		store: []todoitem.ToDoItem{
			{Id: 1, Title: "Task 1", IsComplete: false},
			{Id: 2, Title: "Task 2", IsComplete: false}},
		lock: sync.RWMutex{},
	}

	err := testRepo.UpdateItemCompletionStatusById(true, 2)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !testRepo.store[1].IsComplete {
		t.Errorf("Expected IsComplete to be true, got %v", testRepo.store[1].IsComplete)
	}
	err = testRepo.UpdateItemCompletionStatusById(false, 2)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if testRepo.store[1].IsComplete {
		t.Errorf("Expected IsComplete to be false, got %v", testRepo.store[1].IsComplete)
	}
}

func TestGetById(t *testing.T) {
	testRepo := InMemory{
		store: []todoitem.ToDoItem{
			{Id: 1, Title: "Task 1", IsComplete: false},
			{Id: 2, Title: "Task 2", IsComplete: false}},
		lock: sync.RWMutex{},
	}

	item, err := testRepo.GetById(2)

	if err != nil {
		t.Fatalf("Should not receive error")
	}

	if item.Id != 2 && item.Title != "Task 2" && item.IsComplete {
		t.Errorf("Expected item with ID 2, got %v", item)
	}
}

func TestGetAll(t *testing.T) {
	testRepo := InMemory{
		store: []todoitem.ToDoItem{
			{Id: 1, Title: "Task 1", IsComplete: false},
			{Id: 2, Title: "Task 2", IsComplete: false}},
		lock: sync.RWMutex{},
	}

	items := testRepo.GetAll()
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}
}

func TestFindIndexById(t *testing.T) {
	testRepo := InMemory{
		store: []todoitem.ToDoItem{
			{Id: 1, Title: "Task 1", IsComplete: false},
			{Id: 2, Title: "Task 2", IsComplete: false}},
		lock: sync.RWMutex{},
	}

	index, found := testRepo.findIndexById(2)
	if !found || index != 1 {
		t.Errorf("Expected index 1, got %d", index)
	}
	index, found = testRepo.findIndexById(3)
	if found || index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}
