package main

import (
	"fmt"
)

func updateNumberOdds(number *int) {

	for i := 1; true; i += 2 {
		*number = i
		fmt.Println("odds", *number)

		if *number%2 == 0 {
			fmt.Println("I found an even", *number)
		}

	}

}

func updateNumberEvens(number *int) {

	for i := 2; true; i += 2 {
		*number = i
		fmt.Println("evens", *number)
	}

	if *number%2 == 1 {
		fmt.Println("I found an odd", *number)
	}

}

func main() {

	number := 0

	go updateNumberOdds(&number)
	go updateNumberEvens(&number)

	select {}
}
