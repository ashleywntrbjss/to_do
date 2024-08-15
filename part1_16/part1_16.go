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

func printToDoTitle(item ToDoItem, waitGroup *sync.WaitGroup, mutex *sync.Mutex, nextPairChannel chan struct{}, statusFollowTitleChannel chan struct{}) {
	defer waitGroup.Done()

	<-nextPairChannel
	fmt.Println("Item name:", item.Title)
	statusFollowTitleChannel <- struct{}{}
}

func printToDoStatus(item ToDoItem, waitGroup *sync.WaitGroup, mutex *sync.Mutex, channel chan struct{}) {
	defer waitGroup.Done()
	<-channel
	fmt.Println("Completion status:", item.IsComplete)
	close(channel)
}

func PrintList() {
	toDoItems := []ToDoItem{{1, "Washing up", true}, {2, "Ironing", false}, {3, "Food shop", true}, {4, "Wash car", false}, {5, "Buy fish and chips", false}, {6, "Renew car insurance", true}, {7, "Complete Go tasks", true}, {8, "return parcel", false}, {9, "Organise cutlery drawer", true}, {10, "Dust the mantle", true}}

	var waitGroup = new(sync.WaitGroup)
	var mutex = new(sync.Mutex)

	var nextPairChannel = make(chan struct{}, 1)

	nextPairChannel <- struct{}{}

	for _, item := range toDoItems {
		waitGroup.Add(2)
		var statusFollowTitleChannel = make(chan struct{})
		go printToDoTitle(item, waitGroup, mutex, nextPairChannel, statusFollowTitleChannel)
		go printToDoStatus(item, waitGroup, mutex, statusFollowTitleChannel)
	}
	waitGroup.Wait()
	close(nextPairChannel)
}

func main() {
	PrintList()
}
