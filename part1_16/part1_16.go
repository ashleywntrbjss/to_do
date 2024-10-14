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

func printToDoTitle(item ToDoItem, waitGroup *sync.WaitGroup, nextPairChannel chan struct{}, statusFollowTitleChannel chan struct{}) {
	defer waitGroup.Done()

	<-nextPairChannel
	fmt.Println("Item name:", item.Title)
	statusFollowTitleChannel <- struct{}{}
}

func printToDoStatus(item ToDoItem, waitGroup *sync.WaitGroup, nextPairChannel chan struct{}, statusFollowTitleChannel chan struct{}) {
	defer waitGroup.Done()
	defer close(statusFollowTitleChannel)

	<-statusFollowTitleChannel
	fmt.Println("Completion status:", item.IsComplete)
	nextPairChannel <- struct{}{}
}

func PrintList() {
	toDoItems := []ToDoItem{{1, "Washing up", true}, {2, "Ironing", false}, {3, "Food shop", true}, {4, "Wash car", false}, {5, "Buy fish and chips", false}, {6, "Renew car insurance", true}, {7, "Complete Go tasks", true}, {8, "return parcel", false}, {9, "Organise cutlery drawer", true}, {10, "Dust the mantle", true}}

	var waitGroup = new(sync.WaitGroup)

	var nextPairChannel = make(chan struct{}, 1)

	nextPairChannel <- struct{}{}
	defer close(nextPairChannel)

	for _, item := range toDoItems {
		waitGroup.Add(2)
		var statusFollowTitleChannel = make(chan struct{})
		go printToDoTitle(item, waitGroup, nextPairChannel, statusFollowTitleChannel)
		go printToDoStatus(item, waitGroup, nextPairChannel, statusFollowTitleChannel)
	}
	waitGroup.Wait()
}

func main() {
	PrintList()
}
