package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Golang")

	var slice = []string{}

	fmt.Printf("Type of slice is %T \n", slice)

	slice = append(slice, "Apple", "Mango", "Orange")

	fmt.Println(slice)

	slice = append(slice[1:2])

	fmt.Println(slice)

	newSlice := make([]int, 4)

	newSlice[0] = 10
	newSlice[1] = 20
	newSlice[2] = 40
	newSlice[3] = 30

	fmt.Println(newSlice)

	fmt.Println(sort.IntsAreSorted(newSlice))

	sort.Ints(newSlice)

	// Removing element in slice based on index

	const index int = 2

	newSlice = append(newSlice[:index], newSlice[:index+1]...)

	fmt.Println(newSlice)
}
