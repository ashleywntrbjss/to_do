package main

import "testing"

func TestAddItem(t *testing.T) {
	toDoItemRepo = []ToDoItem{}
	item := ToDoItem{Id: 1, Title: "Task 1", IsComplete: false}
	AddItem(item)
	if len(toDoItemRepo) != 1 {
		t.Errorf("Expected 1 item, got %d", len(toDoItemRepo))
	}
	if toDoItemRepo[0] != item {
		t.Errorf("Expected item %v, got %v", item, toDoItemRepo[0])
	}
}

func TestRemoveItemById(t *testing.T) {
	toDoItemRepo = []ToDoItem{
		{Id: 1, Title: "Task 1", IsComplete: false},
		{Id: 2, Title: "Task 2", IsComplete: false},
	}
	RemoveItemById(2)
	if len(toDoItemRepo) != 1 {
		t.Errorf("Expected 1 item, got %d", len(toDoItemRepo))
	}
	if toDoItemRepo[0].Id != 1 {
		t.Errorf("Expected item with ID 1, got %d", toDoItemRepo[0].Id)
	}
}

func TestUpdateItemTitleById(t *testing.T) {
	toDoItemRepo = []ToDoItem{
		{Id: 1, Title: "Task 1", IsComplete: false},
		{Id: 2, Title: "Task 2", IsComplete: false},
	}
	UpdateItemTitleById(2, "Updated Task 2")
	if toDoItemRepo[1].Title != "Updated Task 2" {
		t.Errorf("Expected title 'Updated Task 2', got %s", toDoItemRepo[1].Title)
	}
}

func TestToggleCompletionById(t *testing.T) {
	toDoItemRepo = []ToDoItem{
		{Id: 1, Title: "Task 1", IsComplete: false},
		{Id: 2, Title: "Task 2", IsComplete: false},
	}
	ToggleItemCompletionStatusById(2)
	if !toDoItemRepo[1].IsComplete {
		t.Errorf("Expected IsComplete to be true, got %v", toDoItemRepo[1].IsComplete)
	}
	ToggleItemCompletionStatusById(2)
	if toDoItemRepo[1].IsComplete {
		t.Errorf("Expected IsComplete to be false, got %v", toDoItemRepo[1].IsComplete)
	}
}

func TestGetById(t *testing.T) {
	toDoItemRepo = []ToDoItem{
		{Id: 1, Title: "Task 1", IsComplete: false},
		{Id: 2, Title: "Task 2", IsComplete: false},
	}
	item := GetById(2)
	if item == nil || item.Id != 2 {
		t.Errorf("Expected item with ID 2, got %v", item)
	}
}

func TestGetAll(t *testing.T) {
	toDoItemRepo = []ToDoItem{
		{Id: 1, Title: "Task 1", IsComplete: false},
		{Id: 2, Title: "Task 2", IsComplete: false},
	}
	items := GetAll()
	if len(*items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(*items))
	}
}

func TestFindIndexById(t *testing.T) {
	toDoItemRepo = []ToDoItem{
		{Id: 1, Title: "Task 1", IsComplete: false},
		{Id: 2, Title: "Task 2", IsComplete: false},
	}
	index, found := findIndexById(2)
	if !found || index != 1 {
		t.Errorf("Expected index 1, got %d", index)
	}
	index, found = findIndexById(3)
	if found || index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}
