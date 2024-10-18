package api

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func handleGETToDoItem(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/view'")

	activeId := request.PathValue("itemId")
	activeIdAsInt, err := strconv.Atoi(activeId)

	if err != nil {
		fmt.Println("error converting activeId to int:", err)
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	responseItem, err := repo.GetById(activeIdAsInt)

	if err != nil {
		fmt.Println("error getting item:", err)
		http.NotFound(writer, request)
		return
	}

	encodeJson(writer, responseItem)
}

func handleGETAllToDoItemsPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/view-all'")

	encodeJson(writer, repo.GetAll())
}

func handlePOSTAddNewToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/add-new'")

	acceptHeader := request.Header.Get("Accept")

	switch acceptHeader {
	case "application/json":
		var toDo todoitem.ToDoItem

		if err := json.NewDecoder(request.Body).Decode(&toDo); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		newItemIndex := repo.AddNew(toDo)

		writer.Header().Set("Location", "/item/"+strconv.Itoa(newItemIndex))
		_, err := writer.Write([]byte("Item added with index: " + strconv.Itoa(newItemIndex)))
		if err != nil {
			return
		}
		writer.WriteHeader(http.StatusCreated)

		http.Redirect(writer, request, "/view-all", http.StatusSeeOther)
	default:
		http.Error(writer, "unsupported Accept header: "+acceptHeader, http.StatusNotAcceptable)
	}
}

func handlePATCHEditToDoItem(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/edit/'")
	err := request.ParseForm()
	if err != nil {
		fmt.Println("unable to parse form", err)
		http.Error(writer, "unable to parse form", http.StatusBadRequest)
		return
	}

	itemId := request.FormValue("id")
	itemIdAsInt, err := strconv.Atoi(itemId)
	if err != nil {
		fmt.Println("invalid item id format", err)
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	item, err := repo.GetById(itemIdAsInt)
	if err != nil {
		fmt.Println("item id not found", err)
		http.Error(writer, "itemId not found", http.StatusNotFound)
		return
	}

	title := request.FormValue("title")

	if title != "" {
		if title != item.Title {
			err := repo.UpdateItemTitleById(title, itemIdAsInt)
			if err != nil {
				fmt.Println("unable to update item title", err)
				http.Error(writer, "unable to update item title", http.StatusInternalServerError)
				return
			}
			successMessage(writer, "updated item title")
		}
	}

	isComplete := request.FormValue("isComplete")

	if isComplete != "" {
		var err error
		if isComplete == "true" {
			err = repo.UpdateItemCompletionStatusById(true, itemIdAsInt)
		} else if isComplete == "false" {
			err = repo.UpdateItemCompletionStatusById(false, itemIdAsInt)
		} else {
			err = errors.New("invalid value for isComplete")
		}

		if err != nil {
			fmt.Println("unable to update item completion status", err)
			http.Error(writer, "unable to update item completion status", http.StatusBadRequest)
			return
		}

		successMessage(writer, "updated item completion status")
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
