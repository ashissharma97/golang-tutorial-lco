package main

import "fmt"

func main() {
	fmt.Println("Defer in Golang")
	fmt.Println("Hello")

	defer fmt.Println("Two")
	defer fmt.Println("One")
	count()
	defer fmt.Println("World")
}

func count() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
