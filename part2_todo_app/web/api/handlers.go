package api

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo/inMemory"
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"log/slog"
	"net/http"
	"strconv"
)

func handleGETToDoItem(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	activeIdAsInt, err := strconv.Atoi(request.PathValue("itemId"))

	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error converting activeId to int: %s", err.Error()))

		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	responseItem, err := activeRepo.GetById(ctx, activeIdAsInt)

	if err != nil {

		fmt.Println("error getting item:", err.Error())

		if errors.Is(err, pg.ErrNoRows) || errors.Is(err, inMemory.NotFoundError) {
			http.NotFound(writer, request)
			return
		}

		http.Error(writer, "internal server error", http.StatusInternalServerError)

		return
	}

	encodeJson(ctx, writer, responseItem)
}

func handleGETAllToDoItems(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	items, err := activeRepo.GetAll(ctx)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("failed to get all to do items: %s", err.Error()))

		if errors.Is(err, pg.ErrNoRows) || errors.Is(err, inMemory.NotFoundError) {
			http.Error(writer, "items not found", http.StatusNotFound)
			return
		}

		http.Error(writer, "internal server error", http.StatusInternalServerError)
	}

	encodeJson(ctx, writer, items)
}

func handlePOSTCreateToDoItem(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var toDo todoitem.ToDoItem

	err := decodeJSONBody(writer, request, &toDo)

	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error decoding request body: %s", err.Error()))
		http.Error(writer, "error decoding request content", http.StatusBadRequest)
		return
	}

	if toDo.Title == "" {
		fmt.Println("Validation failed: item must have title")
		http.Error(writer, "validation failed: item must have title", http.StatusBadRequest)
		return
	}

	newItemIndex, err := activeRepo.AddNew(ctx, toDo)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error adding new item: %s", err.Error()))
		http.Error(writer, "failed to save new to do item", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Location", "/item/"+strconv.Itoa(newItemIndex))
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write([]byte("Item added with index: " + strconv.Itoa(newItemIndex)))
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error writing response: %v", err.Error()))

		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlePUTEditToDoItem(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	slog.InfoContext(ctx, "Update to do item")

	var toDo todoitem.ToDoItem

	err := decodeJSONBody(writer, request, &toDo)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error decoding request body: %v", err.Error()))
		http.Error(writer, "error decoding request content", http.StatusBadRequest)
		return
	}

	if toDo.Title == "" {
		slog.ErrorContext(ctx, "Validation failed: item must have title")
		http.Error(writer, "validation failed: item must have title", http.StatusBadRequest)
		return
	}

	_, err = activeRepo.GetById(ctx, toDo.Id)

	if err != nil {
		slog.ErrorContext(ctx, "Validation failed: failed to retrieve existing to do item")
		http.Error(writer, "validation failed: failed to retrieve existing to do item", http.StatusBadRequest)
		return
	}

	err = activeRepo.UpdateItemTitleById(ctx, toDo.Title, toDo.Id)

	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("Failed to update item: %v ", err.Error()))
		http.Error(writer, "failed to update to do item title", http.StatusBadRequest)
		return
	}

	err = activeRepo.UpdateItemCompletionStatusById(ctx, toDo.IsComplete, toDo.Id)

	if err != nil {

		slog.ErrorContext(ctx, fmt.Sprintf("Failed to update item: %v", err.Error()))
		http.Error(writer, "failed to update to do item title", http.StatusInternalServerError)
		return
	}
}

func handlePATCHToggleComplete(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	requestIdAsInt, err := strconv.Atoi(request.PathValue("itemId"))

	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error converting activeId to int: %v", err.Error()))
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	requestItem, err := activeRepo.GetById(ctx, requestIdAsInt)

	if err != nil {
		slog.ErrorContext(ctx, "Validation failed: failed to retrieve existing to do item")
		http.Error(writer, "validation failed: item must have title", http.StatusBadRequest)
		return
	}

	err = activeRepo.UpdateItemCompletionStatusById(ctx, !requestItem.IsComplete, requestItem.Id)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("failed to update item completion status: %v", err.Error()))
		http.Error(writer, "failed to update item completion status", http.StatusInternalServerError)
		return
	}

}

func handleDELETEToDoItem(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	activeIdAsInt, err := strconv.Atoi(request.PathValue("itemId"))

	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error converting activeId to int: %s", err.Error()))
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	err = activeRepo.DeleteItemById(ctx, activeIdAsInt)

	if err != nil {

		slog.ErrorContext(ctx, fmt.Sprintf("error deleting item: %s", err.Error()))

		if errors.Is(err, pg.ErrNoRows) || errors.Is(err, inMemory.NotFoundError) {
			http.Error(writer, "item not found", http.StatusNotFound)
			return
		}

		http.Error(writer, "failed to delete item", http.StatusInternalServerError)
		return
	}
}
