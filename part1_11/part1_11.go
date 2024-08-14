package part1_11

import (
	"encoding/json"
	"fmt"
)

type ToDoItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func ToDoItemsAsJson() {
	item1 := ToDoItem{1, "Washing up"}
	item2 := ToDoItem{2, "Ironing"}
	item3 := ToDoItem{3, "Food shop"}
	item4 := ToDoItem{4, "Wash car"}
	item5 := ToDoItem{5, "Take out recycling"}
	item6 := ToDoItem{6, "Dust the mantle"}
	item7 := ToDoItem{7, "Whites wash"}
	item8 := ToDoItem{8, "Lights wash"}
	item9 := ToDoItem{9, "Dark wash"}
	item10 := ToDoItem{10, "Return parcel"}

	toDoItems := []ToDoItem{item1, item2, item3, item4, item5, item6, item7, item8, item9, item10}
	marshalledItems, _ := json.Marshal(toDoItems)
	fmt.Println(string(marshalledItems))
}
