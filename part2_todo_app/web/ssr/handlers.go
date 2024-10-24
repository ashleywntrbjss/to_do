package ssr

import (
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"
)

func handleGETHomePage(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	getTemplateAndExecute(ctx, "home.gohtml", writer, nil)
}

func handleGETViewToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	activeId := request.PathValue("itemId")
	activeIdAsInt, err := strconv.Atoi(activeId)

	if err != nil {
		slog.ErrorContext(ctx, "invalid item id format in request: %v", err.Error())
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	activeItem, err := activeRepo.GetById(ctx, activeIdAsInt)
	if err != nil {
		slog.WarnContext(ctx, "unable to find requested item: %v", err.Error())
		http.Error(writer, "itemId not found", http.StatusNotFound)
		return
	}

	getTemplateAndExecute(ctx, "view.gohtml", writer, activeItem)
}

func handleGETViewAllToDoItemsPage(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	items, err := activeRepo.GetAll(ctx)
	if err != nil {
		slog.WarnContext(ctx, "items not found: %v", err.Error())
		http.Error(writer, "items not found", http.StatusNotFound)
		return
	}
	getTemplateAndExecute(ctx, "viewAll.gohtml", writer, items)
}

func handleGETCreateToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	getTemplateAndExecute(ctx, "create.gohtml", writer, nil)
}

func handleGETEditToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	activeId := request.PathValue("itemId")
	activeIdAsInt, err := strconv.Atoi(activeId)

	if err != nil {
		slog.ErrorContext(ctx, "invalid item id format in request: %v", err.Error())
		http.Error(writer, "invalid itemId format", http.StatusBadRequest)
		return
	}

	activeItem, err := activeRepo.GetById(ctx, activeIdAsInt)
	if err != nil {
		slog.WarnContext(ctx, "item id not found: %v", err.Error())
		http.Error(writer, "itemId not found", http.StatusNotFound)
		return
	}

	getTemplateAndExecute(ctx, "edit.gohtml", writer, activeItem)
}

func handleGETFavicon(writer http.ResponseWriter, request *http.Request) {
	faviconFilepath := filepath.Join("part2_todo_app", "web", "ssr", "templates", "layout", "toDoFavicon.ico")
	http.ServeFile(writer, request, faviconFilepath)
}
