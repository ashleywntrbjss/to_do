package main

import "fmt"

func main() {
	VariadicToDoList("Washing up", "Ironing", "Food shop", "Wash car", "Take out recycling", "Dust the mantle", "Whites wash", "Lights wash", "Dark wash", "Return parcel")
}

func VariadicToDoList(arg ...string) {
	startOfList := 1

	fmt.Println("To do:")
	for index, item := range arg {
		fmt.Printf("%v: %v\n", index+startOfList, item)
	}

}
