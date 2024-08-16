package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHomePage)

	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}
}

func handleHomePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("GET", "Root '/' ")

	activeTemplate, err := template.ParseFiles("home.gohtml")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	err = activeTemplate.Execute(writer, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
