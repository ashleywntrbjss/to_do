package testing

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/todoitem"
	"bjss.com/ashley.winter/to_do/part2_todo_app/web/api"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"testing"
)

func viewItemRequest(baseUrl string, itemId int, waitGroup *sync.WaitGroup, t *testing.T) {
	defer waitGroup.Done()
	viewItemUrl := baseUrl + "/api/get/" + strconv.Itoa(itemId)

	viewItemResponse, err := getRequest(viewItemUrl)

	if err != nil {
		t.Error("Error with get request:", err)
	}

	fmt.Println(viewItemResponse)
}

func editItemRequest(baseUrl string, item todoitem.ToDoItem, waitGroup *sync.WaitGroup, t *testing.T) {
	defer waitGroup.Done()

	editItemUrl := baseUrl + "/api/edit"

	err := makePutRequest(editItemUrl, item)

	if err != nil {
		t.Error("Error with put request:", err)
	}
}

func makePutRequest(url string, data any) error {
	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}

	// Create a new PUT request
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing body", err)
		}
	}(resp.Body)

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response status: %v", resp.Status)
	}

	return nil
}

func TestStressApi(t *testing.T) {
	go api.ListenAndServe()

	serverAddress := "http://" + api.ServerAddress

	var wait = new(sync.WaitGroup)

	for i := range 500 {
		wait.Add(2)
		go viewItemRequest(serverAddress, 1, wait, t)

		if i%2 == 0 {
			editItem := todoitem.ToDoItem{
				Id:         1,
				Title:      "Washing up",
				IsComplete: true,
			}
			go editItemRequest(serverAddress, editItem, wait, t)
		} else if i%3 == 0 {
			editItem := todoitem.ToDoItem{
				Id:         1,
				Title:      "walk the dog",
				IsComplete: true,
			}
			go editItemRequest(serverAddress, editItem, wait, t)
		} else if i%5 == 0 {
			editItem := todoitem.ToDoItem{
				Id:         1,
				Title:      "load the dishwasher",
				IsComplete: false,
			}
			go editItemRequest(serverAddress, editItem, wait, t)
		} else {
			editItem := todoitem.ToDoItem{
				Id:         1,
				Title:      "Washing up",
				IsComplete: false,
			}
			go editItemRequest(serverAddress, editItem, wait, t)
		}

	}

	wait.Wait()
}

func getRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("failed to close request body: %v", err)
		}
	}(resp.Body)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}
	return string(body), nil
}
