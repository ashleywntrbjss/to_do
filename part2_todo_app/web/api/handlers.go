package api

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo/inMemory"
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"fmt"
	"net/http"
	"strconv"
)

func handleGETToDoItem(writer http.ResponseWriter, request *http.Request, repo repo.repo) {
	activeId := request.PathValue("itemId")
	activeIdAsInt, err := strconv.Atoi(activeId)

	if err != nil {
		fmt.Println("error converting activeId to int:", err)
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	responseItem, err := inMemory.GetById(activeIdAsInt)

	if err != nil {
		fmt.Println("error getting item:", err)
		http.NotFound(writer, request)
		return
	}

	encodeJson(writer, responseItem)
}

func handleGETAllToDoItems(writer http.ResponseWriter, request *http.Request) {
	encodeJson(writer, inMemory.GetAll())
}

func handlePOSTCreateToDoItem(writer http.ResponseWriter, request *http.Request) {
	var toDo todoitem.ToDoItem

	err := decodeJSONBody(writer, request, &toDo)

	if err != nil {
		fmt.Println("error decoding request body:", err)
		http.Error(writer, "error decoding request content", http.StatusBadRequest)
		return
	}

	if toDo.Title == "" {
		fmt.Println("Validation failed: item must have title")
		http.Error(writer, "validation failed: item must have title", http.StatusBadRequest)
		return
	}

	newItemIndex, err := inMemory.AddNew(toDo)
	if err != nil {
		fmt.Println("error adding new item:", err)
		http.Error(writer, "error saving new to do item", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Location", "/item/"+strconv.Itoa(newItemIndex))
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write([]byte("Item added with index: " + strconv.Itoa(newItemIndex)))
	if err != nil {
		fmt.Println("error writing response:", err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func handlePUTEditToDoItem(writer http.ResponseWriter, request *http.Request) {
	var toDo todoitem.ToDoItem

	err := decodeJSONBody(writer, request, &toDo)

	if err != nil {
		fmt.Println("error decoding request body:", err)
		http.Error(writer, "error decoding request content", http.StatusBadRequest)
		return
	}

	if toDo.Title == "" {
		fmt.Println("Validation failed: item must have title")
		http.Error(writer, "validation failed: item must have title", http.StatusBadRequest)
		return
	}

	_, err = inMemory.GetById(toDo.Id)

	if err != nil {
		fmt.Println("Validation failed: failed to retrieve existing to do item")
		http.Error(writer, "validation failed: item must have title", http.StatusBadRequest)
		return
	}

	err = inMemory.UpdateItemTitleById(toDo.Title, toDo.Id)

	if err != nil {
		fmt.Println("Failed to update item:", err)
		http.Error(writer, "failed to update to do item title", http.StatusBadRequest)
		return
	}

	err = inMemory.UpdateItemCompletionStatusById(toDo.IsComplete, toDo.Id)

	if err != nil {
		fmt.Println("Failed to update item:", err)
		http.Error(writer, "failed to update to do item title", http.StatusBadRequest)
		return
	}
}

func successMessage(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println("error writing success message:", err)
		http.Error(writer, "internal server error, see logs for details", http.StatusInternalServerError)
		return
	}
}
