package ssr

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func ListenAndServe() {
	mux := http.NewServeMux()

	mux.HandleFunc("/view-all", handleViewAllPage)
	mux.HandleFunc("/add-new", handleAddNewPage)

	mux.HandleFunc("/", handleHomePage)

	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}
}

func handleHomePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("GET", "Root '/' ")

	templatePath := filepath.Join("part2_18_todo_app_web", "home.gohtml")

	activeTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		cwd, err := os.Getwd()
		fmt.Println("Error parsing template:", err)
		fmt.Println("Current working directory:", cwd)
		return
	}

	err = activeTemplate.Execute(writer, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}

func handleViewAllPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("GET", "'/view-all'")

	templatePath := filepath.Join("part2_18_todo_app_web", "viewAll.gohtml")

	activeTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		cwd, err := os.Getwd()
		fmt.Println("Error parsing template:", err)
		fmt.Println("Current working directory:", cwd)
		return
	}

	err = activeTemplate.Execute(writer, repo.GetAll())
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}

func handleAddNewPage(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		fmt.Println("GET", "'/view-all'")

		templatePath := filepath.Join("part2_18_todo_app_web", "addNew.gohtml")

		activeTemplate, err := template.ParseFiles(templatePath)
		if err != nil {
			cwd, err := os.Getwd()
			fmt.Println("Error parsing template:", err)
			fmt.Println("Current working directory:", cwd)
			return
		}

		err = activeTemplate.Execute(writer, repo.GetAll())
		if err != nil {
			fmt.Println("Error executing template:", err)
			return
		}
	}

	if request.Method == "POST" {
		fmt.Println("POST", "'/add-new'")

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
