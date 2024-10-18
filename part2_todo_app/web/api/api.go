package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ListenAndServe() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/get/{itemId}", handleGETToDoItem)
	mux.HandleFunc("GET /api/get-all", handleGETAllToDoItemsPage)
	mux.HandleFunc("POST /create", handlePOSTAddNewToDoItemPage)
	mux.HandleFunc("PATCH /edit", handlePATCHEditToDoItem)

	fmt.Println("Starting api server at http://localhost:8085")
	err := http.ListenAndServe("localhost:8085", mux)
	if err != nil {
		log.Fatalln("there's an error with the server:", err)
	}
}

func encodeJson(writer http.ResponseWriter, data any) {
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		fmt.Println("error encoding json:", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
