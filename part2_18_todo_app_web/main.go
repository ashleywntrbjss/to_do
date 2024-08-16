package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHomePage)

	mux.HandleFunc("/view-all", handleViewAllPage)

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

	err = activeTemplate.Execute(writer, GetAll())
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
