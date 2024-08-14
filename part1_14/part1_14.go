package main

import "fmt"

var numbers []int32

func updateData(even bool) {
	if even {
		numbers = append(numbers, numbers[len(numbers)-1]+2)
	} else {
		numbers = append(numbers, numbers[len(numbers)-1]+1)
	}
}

func main() {

	numbers = append(numbers, 1)

	for _ = range 50 {
		go updateData(false)
		go updateData(true)

		fmt.Println(numbers)
	}
}
