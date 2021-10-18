package main

import "fmt"

func main() {
	fmt.Println("If Else in Golang")

	num := 9

	if num%2 == 0 {
		fmt.Println("Num is even")
	} else if num%2 != 0 {
		fmt.Println("Num is odd")
	} else {
		fmt.Println("Invalid Value")
	}
}
