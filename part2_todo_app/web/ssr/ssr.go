package ssr

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func ListenAndServe() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /view-all", handleGETViewAllToDoItemsPage)
	mux.HandleFunc("GET /add-new", handleGETAddNewToDoItemPage)
	mux.HandleFunc("POST /add-new", handlePOSTAddNewToDoItemPage)
	mux.HandleFunc("GET /edit/{itemId}", handleGETEditToDoItemPage)

	mux.HandleFunc("GET /favicon.ico", handleGETFavicon)
	mux.HandleFunc("GET /", handleGETHomePage)

	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}
}

func getTemplateByFilename(filename string) (template.Template, error) {
	baseFilepath := filepath.Join("part2_todo_app", "web", "ssr", "templates")

	templatePath := filepath.Join(baseFilepath, filename)

	activeTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		// include the current working directory to provide context
		cwd, getWdErr := os.Getwd()
		if getWdErr != nil {
			panic(getWdErr)
		}

		return template.Template{}, errors.New("Error parsing template " + err.Error() + ". Current working directory:" + cwd)
	}

	return *activeTemplate, nil
}

func handleGETHomePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "Root '/' ")

	if request.Method == "GET" {
		activeTemplate, err := getTemplateByFilename("home.gohtml")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = activeTemplate.Execute(writer, nil)
		if err != nil {
			fmt.Println("Error executing template:", err)
			return
		}
	}
}

func handleGETViewAllToDoItemsPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/view-all'")

	activeTemplate, err := getTemplateByFilename("viewAll.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = activeTemplate.Execute(writer, repo.GetAll())
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

}

func handleGETAddNewToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/view-all'")
	activeTemplate, err := getTemplateByFilename("addNew.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = activeTemplate.Execute(writer, repo.GetAll())
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}

func handlePOSTAddNewToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/view-all'")

	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Unable to parse form", http.StatusBadRequest)
		return
	}

	title := request.FormValue("title")
	if title == "" {
		http.Error(writer, "Title is required", http.StatusBadRequest)
		return
	}

	_ = repo.CreateItemFromTitle(title)

	http.Redirect(writer, request, "/view-all", http.StatusSeeOther)
}

func handleGETEditToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/edit/'")
	activeId := request.PathValue("itemId")
	activeIdAsInt, err := strconv.Atoi(activeId)

	if err != nil {
		http.Error(writer, "Invalid itemId format", http.StatusBadRequest)
	}

	activeItem, err := repo.GetById(activeIdAsInt)
	if err != nil {
		http.Error(writer, "Itemid not found", http.StatusNotFound)
	}

	fmt.Println(activeItem)

}

func handleGETFavicon(writer http.ResponseWriter, request *http.Request) {
	return
}
