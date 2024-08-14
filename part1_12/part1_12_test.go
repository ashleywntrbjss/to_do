package part1_12

func ExampleToDoItemsAsJson(){
	ToDoItemsAsJson(ToDoItem{1, "Washing up"}, ToDoItem{2, "Ironing"}, ToDoItem{3, "Food shop"}, ToDoItem{4,  "Wash car"})
	// Output: [{"id":1,"title":"Washing up"},{"id":2,"title":"Ironing"},{"id":3,"title":"Food shop"},{"id":4,"title":"Wash car"}]
}