package main

import (
	"fmt"
)

var numberOfAttempts = 3000000

func updateNumberEvens(numberChan chan int) {
	defer close(numberChan)

	for i := 0; i <= numberOfAttempts; i += 2 {
		numberChan <- i
	}
}

func updateNumberOdds(numberChan chan int) {
	defer close(numberChan)

	for i := 1; i <= numberOfAttempts; i += 2 {
		numberChan <- i
	}
}

func main() {

	var oddChan = make(chan int)
	var evenChan = make(chan int)

	go updateNumberOdds(oddChan)
	go updateNumberEvens(evenChan)

	for {
		select {
		case number, ok := <-evenChan:
			if number%2 == 0 {
				fmt.Println("Even number", number)
			} else {
				fmt.Println("Received wrong type of int on the channel")
			}
			if !ok {
				fmt.Println("Even channel closed")
				evenChan = nil
			}

		case number, ok := <-oddChan:
			if number%2 == 1 {
				fmt.Println("Odd number", number)
			} else {
				fmt.Println("Received wrong type of int on the channel")
			}

			if !ok {
				fmt.Println("Odd channel closed")
				oddChan = nil
			}
		}

		if oddChan == nil && evenChan == nil {
			fmt.Println("Channels closed")
			break
		}

	}

}
