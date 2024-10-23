package api

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"fmt"
	"net/http"
	"strconv"
)

func handleGETToDoItem(writer http.ResponseWriter, request *http.Request) {
	activeIdAsInt, err := strconv.Atoi(request.PathValue("itemId"))

	if err != nil {
		fmt.Println("error converting activeId to int:", err)
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	responseItem, err := activeRepo.GetById(activeIdAsInt)

	if err != nil {
		fmt.Println("error getting item:", err)
		http.NotFound(writer, request)
		return
	}

	encodeJson(writer, responseItem)
}

func handleGETAllToDoItems(writer http.ResponseWriter, request *http.Request) {
	items, err := activeRepo.GetAll()
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "items not found", http.StatusNotFound)
		return
	}

	encodeJson(writer, items)
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

	newItemIndex, err := activeRepo.AddNew(toDo)
	if err != nil {
		fmt.Println("error adding new item:", err)
		http.Error(writer, "failed to save new to do item", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Location", "/item/"+strconv.Itoa(newItemIndex))
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write([]byte("Item added with index: " + strconv.Itoa(newItemIndex)))
	if err != nil {
		fmt.Println("error writing response:", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
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

	_, err = activeRepo.GetById(toDo.Id)

	if err != nil {
		fmt.Println("Validation failed: failed to retrieve existing to do item")
		http.Error(writer, "validation failed: item must have title", http.StatusBadRequest)
		return
	}

	err = activeRepo.UpdateItemTitleById(toDo.Title, toDo.Id)

	if err != nil {
		fmt.Println("Failed to update item:", err)
		http.Error(writer, "failed to update to do item title", http.StatusBadRequest)
		return
	}

	err = activeRepo.UpdateItemCompletionStatusById(toDo.IsComplete, toDo.Id)

	if err != nil {
		fmt.Println("Failed to update item:", err)
		http.Error(writer, "failed to update to do item title", http.StatusInternalServerError)
		return
	}
}

func handlePATCHToggleComplete(writer http.ResponseWriter, request *http.Request) {
	requestIdAsInt, err := strconv.Atoi(request.PathValue("itemId"))

	if err != nil {
		fmt.Println("error converting activeId to int:", err)
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	requestItem, err := activeRepo.GetById(requestIdAsInt)

	if err != nil {
		fmt.Println("Validation failed: failed to retrieve existing to do item")
		http.Error(writer, "validation failed: item must have title", http.StatusBadRequest)
		return
	}

	err = activeRepo.UpdateItemCompletionStatusById(!requestItem.IsComplete, requestItem.Id)
	if err != nil {
		fmt.Println("failed to update item completion status")
		http.Error(writer, "failed to update item completion status", http.StatusInternalServerError)
		return
	}

}

func handleDELETEToDoItem(writer http.ResponseWriter, request *http.Request) {
	activeIdAsInt, err := strconv.Atoi(request.PathValue("itemId"))

	if err != nil {
		fmt.Println("error converting activeId to int:", err)
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	err = activeRepo.DeleteItemById(activeIdAsInt)

	if err != nil {
		fmt.Println("error deleting item:", err)
		http.Error(writer, "failed to delete item", http.StatusInternalServerError)
		return
	}
}
