package ssr

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"context"
	"errors"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var activeRepo repo.Repo

func ListenAndServe(ctx context.Context, repo repo.Repo) {
	mux := http.NewServeMux()

	activeRepo = repo

	mux.HandleFunc("GET /view/{itemId}", handleGETViewToDoItemPage)
	mux.HandleFunc("GET /view-all", handleGETViewAllToDoItemsPage)
	mux.HandleFunc("GET /create", handleGETCreateToDoItemPage)
	mux.HandleFunc("GET /edit/{itemId}", handleGETEditToDoItemPage)

	mux.HandleFunc("GET /favicon.ico", handleGETFavicon)
	mux.HandleFunc("GET /", handleGETHomePage)

	slog.InfoContext(ctx, "Starting template server at http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", middleware(ctx, mux))
	if err != nil {
		log.Fatalln("there's an error with the server:", err)
	}
}

func middleware(ctx context.Context, existingHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handlerCtx := context.WithValue(request.Context(), "logger", ctx.Value("logger"))
		handlerCtx, cancel := context.WithTimeout(handlerCtx, 5*time.Second)
		defer cancel()

		request = request.WithContext(handlerCtx)

		slog.InfoContext(ctx, fmt.Sprintf("%v - %v", request.Method, request.URL.Path))
		existingHandler.ServeHTTP(writer, request)
	})
}

func getTemplateAndExecute(ctx context.Context, filename string, writer http.ResponseWriter, data any) {
	activeTemplate, err := templateBuilder(filename)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error getting template: %v", err))
		http.Error(writer, "internal Server Error, see logs for details", http.StatusInternalServerError)
		return
	}

	executeTemplate(ctx, activeTemplate, writer, data)
}

func templateBuilder(filename string) (template.Template, error) {
	baseFilepath := filepath.Join("part2_todo_app", "web", "ssr", "templates")

	layoutFilepath := filepath.Join(baseFilepath, "layout")

	baseTemplatePath := filepath.Join(layoutFilepath, "base.gohtml")
	footerTemplatePath := filepath.Join(layoutFilepath, "footer.gohtml")
	navbarTemplatePath := filepath.Join(layoutFilepath, "navbar.gohtml")

	currentPagePath := filepath.Join(baseFilepath, filename)

	activeTemplate, err := template.ParseFiles(currentPagePath, navbarTemplatePath, footerTemplatePath, baseTemplatePath)
	if err != nil {
		// include the current working directory to provide context
		cwd, getWdErr := os.Getwd()
		if getWdErr != nil {
			panic(getWdErr)
		}

		return template.Template{}, errors.New("error parsing template " + err.Error() + ". current working directory:" + cwd)
	}

	return *activeTemplate, nil
}

func executeTemplate(ctx context.Context, template template.Template, writer http.ResponseWriter, data any) {
	err := template.ExecuteTemplate(writer, "base", data)

	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error executing template: %v", err))
		http.Error(writer, "internal server error, see logs for details", http.StatusInternalServerError)
		return
	}
}
