package ssr

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func ListenAndServe() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /view-all", handleGETViewAllToDoItemsPage)
	mux.HandleFunc("GET /create", handleGETCreateToDoItemPage)
	mux.HandleFunc("GET /edit/{itemId}", handleGETEditToDoItemPage)

	mux.HandleFunc("GET /favicon.ico", handleGETFavicon)
	mux.HandleFunc("GET /", handleGETHomePage)

	fmt.Println("Starting template server at http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatalln("there's an error with the server:", err)
	}
}

func getTemplateAndExecute(filename string, writer http.ResponseWriter, data any) {

	activeTemplate, err := templateBuilder(filename)
	if err != nil {
		fmt.Println("error getting template", err)
		http.Error(writer, "internal Server Error, see logs for details", http.StatusInternalServerError)
		return
	}

	executeTemplate(activeTemplate, writer, data)
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

func executeTemplate(template template.Template, writer http.ResponseWriter, data any) {
	err := template.ExecuteTemplate(writer, "base", data)

	if err != nil {
		fmt.Println("error executing template:", err)
		http.Error(writer, "internal server error, see logs for details", http.StatusInternalServerError)
		return
	}
}
