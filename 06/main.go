package main

import "fmt"

func main() {
	fmt.Println("Pointer in Golang")

	myValue := 22

	myValuePtr := &myValue

	fmt.Println("Pointer Value is ", myValuePtr)
	fmt.Println("Pointer Value is ", *myValuePtr)

	*myValuePtr = *myValuePtr * 2

	fmt.Println("Pointer Value is ", *myValuePtr)

}
