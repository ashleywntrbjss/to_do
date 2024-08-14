package main

func ExampleToDoItemsAsJson() {
	marshallToDoItems(ToDoItem{1, "Washing up"}, ToDoItem{2, "Ironing"}, ToDoItem{3, "Food shop"}, ToDoItem{4, "Wash car"}, ToDoItem{5, "Buy fish and chips"}, ToDoItem{6, "Renew car insurance"}, ToDoItem{7, "Complete Go tasks"}, ToDoItem{8, "return parcel"}, ToDoItem{9, "Organise cutlery drawer"}, ToDoItem{10, "Dust the mantle"})
	// Output: [{"id":1,"title":"Washing up"},{"id":2,"title":"Ironing"},{"id":3,"title":"Food shop"},{"id":4,"title":"Wash car"},{"id":5,"title":"Buy fish and chips"},{"id":6,"title":"Renew car insurance"},{"id":7,"title":"Complete Go tasks"},{"id":8,"title":"return parcel"},{"id":9,"title":"Organise cutlery drawer"},{"id":10,"title":"Dust the mantle"}]
}
