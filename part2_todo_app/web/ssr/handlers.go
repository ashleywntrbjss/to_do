package ssr

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"fmt"
	"net/http"
	"strconv"
)

func handleGETHomePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "Root '/' ")

	getTemplateAndExecute("home.gohtml", writer, nil)
}

func handleGETViewAllToDoItemsPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/view-all'")

	getTemplateAndExecute("viewAll.gohtml", writer, repo.GetAll())
}

func handleGETCreateToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/add-new'")

	getTemplateAndExecute("create.gohtml", writer, nil)
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

func handleGETFavicon(writer http.ResponseWriter, request *http.Request) {
	return
}
