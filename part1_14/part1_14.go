package main

import (
	"fmt"
	"sync"
)

func updateNumberOdds(number *int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for i := 1; i < 300; i += 2 {
		*number = i
		fmt.Println("odds", *number)
	}

}

func updateNumberEvens(number *int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for i := 2; i < 300; i += 2 {
		*number = i
		fmt.Println("evens", *number)
	}

}

func main() {

	var number *int
	number = new(int)

	var waitGroup = new(sync.WaitGroup)

	waitGroup.Add(2)
	go updateNumberOdds(number, waitGroup)
	go updateNumberEvens(number, waitGroup)
	waitGroup.Wait()
}
