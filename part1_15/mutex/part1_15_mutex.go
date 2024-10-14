package main

import (
	"fmt"
	"sync"
)

var numberOfAttempts = 3000000

var clashedOdds, clashedEvens int

func updateNumberOdds(number *int, waitGroup *sync.WaitGroup, mutex *sync.Mutex) {
	defer waitGroup.Done()
	for i := 1; i <= numberOfAttempts; i += 2 {
		mutex.Lock()
		*number = i
		fmt.Println("odds", *number)

		if *number%2 == 0 {
			fmt.Println("I found an even when it should be odd", *number)
			clashedOdds++
		}
		mutex.Unlock()
	}
}

func updateNumberEvens(number *int, waitGroup *sync.WaitGroup, mutex *sync.Mutex) {
	defer waitGroup.Done()
	for i := 2; i <= numberOfAttempts; i += 2 {
		mutex.Lock()
		*number = i
		fmt.Println("evens", *number)

		if *number%2 == 1 {
			fmt.Println("I found an odd when it should be even", *number)
			clashedEvens++
		}
		mutex.Unlock()
	}
}

func main() {

	var waitGroup = new(sync.WaitGroup)
	var mutex = new(sync.Mutex)

	waitGroup.Add(2)

	var number int

	go updateNumberOdds(&number, waitGroup, mutex)
	go updateNumberEvens(&number, waitGroup, mutex)

	waitGroup.Wait()
	fmt.Println("==============================")
	fmt.Println("Number of clashed Odds: ", clashedOdds)
	fmt.Println("Number of clashed Evens: ", clashedEvens)
}
