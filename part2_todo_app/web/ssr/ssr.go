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
)

func ListenAndServe() {
	mux := http.NewServeMux()

	mux.HandleFunc("/view-all", handleViewAllToDoItemsPage)
	mux.HandleFunc("/add-new", handleAddNewToDoItemPage)
	mux.HandleFunc("/edit", handleEditToDoItemPage)

	mux.HandleFunc("/favicon.ico", handleFavicon)
	mux.HandleFunc("/", handleHomePage)

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

func handleHomePage(writer http.ResponseWriter, request *http.Request) {
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

func handleViewAllToDoItemsPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/view-all'")
	if request.Method == "GET" {

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
}

func handleAddNewToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/view-all'")
	if request.Method == "GET" {
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

	if request.Method == "POST" {
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
}

func handleEditToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/edit'")
	if request.Method == "GET" {

	}
	if request.Method == "PUT" {

	}
}

func handleFavicon(writer http.ResponseWriter, request *http.Request) {
	return
}
