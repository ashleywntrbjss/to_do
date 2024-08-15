package main

import (
	"fmt"
	"sync"
)

type ToDoItem struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"is_complete"`
}

func printToDoTitle(item ToDoItem, waitGroup *sync.WaitGroup, mutex *sync.Mutex, channel chan bool) {
	defer waitGroup.Done()
	fmt.Println("Item name:", item.Title)
	channel <- true
}

func printToDoStatus(item ToDoItem, waitGroup *sync.WaitGroup, mutex *sync.Mutex, channel chan bool) {
	defer waitGroup.Done()
	<-channel
	fmt.Println("Completion status:", item.IsComplete)
}

func main() {
	PrintList()

}

func PrintList() {
	toDoItems := []ToDoItem{{1, "Washing up", true}, {2, "Ironing", false}, {3, "Food shop", true}, {4, "Wash car", false}, {5, "Buy fish and chips", false}, {6, "Renew car insurance", true}, {7, "Complete Go tasks", true}, {8, "return parcel", false}, {9, "Organise cutlery drawer", true}, {10, "Dust the mantle", true}}

	var waitGroup = new(sync.WaitGroup)
	var mutex = new(sync.Mutex)

	for _, item := range toDoItems {

		waitGroup.Add(2)
		var titlePrintedChannel = make(chan bool)

		go printToDoTitle(item, waitGroup, mutex, titlePrintedChannel)
		go printToDoStatus(item, waitGroup, mutex, titlePrintedChannel)
	}
	waitGroup.Wait()
}
