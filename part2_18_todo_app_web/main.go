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

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}
}

func handleHomePage(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {

		http.NotFound(writer, request)
		return
	}
	if request.Method == "GET" {
		fmt.Println("GET", "Root '/' ")

		activeTemplate, _ := template.ParseFiles("./views/home.gohtml")

		err := activeTemplate.Execute(writer, nil)

		if err != nil {
			return
		}
	}
}
