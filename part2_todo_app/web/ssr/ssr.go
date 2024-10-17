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
	mux.HandleFunc("PATCH /edit", handlePATCHEditToDoItem)

	mux.HandleFunc("GET /favicon.ico", handleGETFavicon)
	mux.HandleFunc("GET /", handleGETHomePage)

	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}
}

func getTemplateAndExecute(filename string, writer http.ResponseWriter, data any) {

	activeTemplate, err := getTemplateByFilename(filename)
	if err != nil {
		http.Error(writer, "Internal Server Error, see logs for details", http.StatusInternalServerError)
		return
	}

	executeTemplate(activeTemplate, writer, data)
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

func executeTemplate(template template.Template, writer http.ResponseWriter, data any) {
	err := template.Execute(writer, data)

	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(writer, "Internal Server Error, see logs for details", http.StatusInternalServerError)
		return
	}
}

func handleGETHomePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "Root '/' ")

	getTemplateAndExecute("home.gohtml", writer, nil)
}

func handleGETViewAllToDoItemsPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/view-all'")

	getTemplateAndExecute("viewAll.gohtml", writer, repo.GetAll())
}

func handleGETAddNewToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/add-new'")

	getTemplateAndExecute("addNew.gohtml", writer, nil)
}

func handlePOSTAddNewToDoItemPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/add-new'")

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
		return
	}

	activeItem, err := repo.GetById(activeIdAsInt)
	if err != nil {
		http.Error(writer, "ItemId not found", http.StatusNotFound)
		return
	}

	getTemplateAndExecute("edit.gohtml", writer, activeItem)

}

func handlePATCHEditToDoItem(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, "'/edit/'")
	err := request.ParseForm()
	if err != nil {
		fmt.Println("unable to parse form", err)
		http.Error(writer, "Unable to parse form", http.StatusBadRequest)
		return
	}

	itemId := request.FormValue("id")
	itemIdAsInt, err := strconv.Atoi(itemId)
	if err != nil {
		fmt.Println("invalid item id format", err)
		http.Error(writer, "Invalid itemId format", http.StatusBadRequest)
		return
	}

	item, err := repo.GetById(itemIdAsInt)
	if err != nil {
		fmt.Println("item id not found", err)
		http.Error(writer, "ItemId not found", http.StatusNotFound)
		return
	}

	title := request.FormValue("title")

	if title != "" {
		if title != item.Title {
			err := repo.UpdateItemTitleById(title, itemIdAsInt)
			if err != nil {
				fmt.Println("unable to update item title", err)
				http.Error(writer, "unable to update item title", http.StatusInternalServerError)
				return
			}
			successMessage(writer, "Updated item title")
		}
	}

	isComplete := request.FormValue("isComplete")

	if isComplete != "" {
		if isComplete == "true" {
			repo.UpdateItemCompletionStatusById(true, itemIdAsInt)
		} else if isComplete == "false" {
			repo.UpdateItemCompletionStatusById(false, itemIdAsInt)
		} else {
			fmt.Println("unable to update item completion status", err)
			http.Error(writer, "Invalid IsComplete value", http.StatusBadRequest)
			return
		}
		successMessage(writer, "Updated item completion status")
	}
}

func successMessage(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println("Error writing success message:", err)
		http.Error(writer, "Internal Server Error, see logs for details", http.StatusInternalServerError)
		return
	}
}

func handleGETFavicon(writer http.ResponseWriter, request *http.Request) {
	return
}
