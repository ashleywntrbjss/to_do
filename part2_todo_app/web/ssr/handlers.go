package ssr

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"net/http"
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

	activeItem, err := repo.GetById(activeIdAsInt)
	if err != nil {
		http.Error(writer, "itemId not found", http.StatusNotFound)
		return
	}

	getTemplateAndExecute("view.gohtml", writer, activeItem)
}

func handleGETViewAllToDoItemsPage(writer http.ResponseWriter, request *http.Request) {
	getTemplateAndExecute("viewAll.gohtml", writer, repo.GetAll())
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

	activeItem, err := repo.GetById(activeIdAsInt)
	if err != nil {
		http.Error(writer, "itemId not found", http.StatusNotFound)
		return
	}

	getTemplateAndExecute("edit.gohtml", writer, activeItem)
}

func handleGETFavicon(writer http.ResponseWriter, request *http.Request) {
	return
}
