package ssr

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func handleGETHomePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "Root '/' ")

	getTemplateAndExecute("home.gohtml", writer, nil)
}

func handleGETViewToDoItem(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/view-all'")

	acceptHeader := request.Header.Get("Accept")
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
		http.Error(writer, "unable to retrieve item", http.StatusBadRequest)
		return
	}

	if acceptHeader == "application/json" {
		writer.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(writer).Encode(responseItem); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(writer, "unsupported Accept header: "+acceptHeader, http.StatusNotAcceptable)
		// no fallback html page so only allow json content type
	}
}

func handleGETViewAllToDoItemsPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/view-all'")

	acceptHeader := request.Header.Get("Accept")

	if acceptHeader == "application/json" {
		writer.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(writer).Encode(repo.GetAll()); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	} else {
		getTemplateAndExecute("viewAll.gohtml", writer, repo.GetAll())
	}

}

func handleGETAddNewToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/add-new'")

	getTemplateAndExecute("addNew.gohtml", writer, nil)
}

func handleGETEditToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/edit/'")
	activeId := request.PathValue("itemId")
	activeIdAsInt, err := strconv.Atoi(activeId)

	if err != nil {
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	activeItem, err := repo.GetById(activeIdAsInt)
	if err != nil {
		http.Error(writer, "itemId not found", http.StatusNotFound)
		return
	}

	getTemplateAndExecute("edit.gohtml", writer, activeItem)

}

func handlePOSTAddNewToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/add-new'")

	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "unable to parse form", http.StatusBadRequest)
		return
	}

	title := request.FormValue("title")
	if title == "" {
		http.Error(writer, "title is required", http.StatusBadRequest)
		return
	}

	_ = repo.CreateItemFromTitle(title)

	http.Redirect(writer, request, "/view-all", http.StatusSeeOther)
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

func handleGETFavicon(writer http.ResponseWriter, request *http.Request) {
	return
}
