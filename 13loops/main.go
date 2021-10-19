package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// Simple Loop

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for i, v := range days {
		fmt.Printf("Index is: %v and value is: %v \n", i, v)
	}

	for i, v := range days {
		if i == 0 || i == 6 {
			fmt.Printf("Weekend today %v \n", v)
		} else {
			fmt.Printf("Working today %v \n", v)
		}
	}
}
