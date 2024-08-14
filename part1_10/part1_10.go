package part1_10

import "fmt"

func VariadicToDoList(arg ...string) {
	startOfList := 1

	fmt.Println("To do:")
	for index, item := range arg {
		fmt.Printf("%v: %v\n", index+startOfList, item)
	}

}
