package ssr

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
)

func handleGETHomePage(writer http.ResponseWriter, request *http.Request) {
	getTemplateAndExecute("home.gohtml", writer, nil)
}

func handleGETViewToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	activeId := request.PathValue("itemId")
	activeIdAsInt, err := strconv.Atoi(activeId)

	if err != nil {
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	activeItem, err := activeRepo.GetById(activeIdAsInt)
	if err != nil {
		http.Error(writer, "itemId not found", http.StatusNotFound)
		return
	}

	getTemplateAndExecute("view.gohtml", writer, activeItem)
}

func handleGETViewAllToDoItemsPage(writer http.ResponseWriter, request *http.Request) {
	items, err := activeRepo.GetAll()
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "items not found", http.StatusNotFound)
		return
	}
	getTemplateAndExecute("viewAll.gohtml", writer, items)
}

func handleGETCreateToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	getTemplateAndExecute("create.gohtml", writer, nil)
}

func handleGETEditToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	activeId := request.PathValue("itemId")
	activeIdAsInt, err := strconv.Atoi(activeId)

	if err != nil {
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	activeItem, err := activeRepo.GetById(activeIdAsInt)
	if err != nil {
		http.Error(writer, "itemId not found", http.StatusNotFound)
		return
	}

	getTemplateAndExecute("edit.gohtml", writer, activeItem)
}

func handleGETFavicon(writer http.ResponseWriter, request *http.Request) {
	faviconFilepath := filepath.Join("part2_todo_app", "web", "ssr", "templates", "layout", "toDoFavicon.ico")
	http.ServeFile(writer, request, faviconFilepath)
}
