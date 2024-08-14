package main

import (
	"fmt"
	"sync"
)

var number int

var waitGroup = new(sync.WaitGroup)

func updateNumberOdds() {
	defer waitGroup.Done()
	for i := 1; i < 150; i += 2 {
		number = i
		fmt.Println("odds", number)
	}

}

func updateNumberEvens() {
	defer waitGroup.Done()
	for i := 2; i < 150; i += 2 {
		number = i
		fmt.Println("evens", number)
	}

}

func main() {

	waitGroup.Add(2)
	go updateNumberOdds()
	go updateNumberEvens()
	waitGroup.Wait()
}
