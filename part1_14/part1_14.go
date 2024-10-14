package main

import (
	"fmt"
)

func updateNumberOdds(number *int) {

	for i := 1; true; i += 2 {
		*number = i
		fmt.Print("\nOdds - ", *number)

		if *number%2 == 0 {
			fmt.Print(" - number is now even ", *number)
		}

	}

}

func updateNumberEvens(number *int) {

	for i := 2; true; i += 2 {
		*number = i
		fmt.Print("\nEvens - ", *number)

		if *number%2 == 1 {
			fmt.Print(" - number is now odd ", *number)
		}

	}

}

func main() {

	var number int

	go updateNumberOdds(&number)
	go updateNumberEvens(&number)

	select {}
}
