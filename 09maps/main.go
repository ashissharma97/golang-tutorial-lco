package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	students := make(map[string]string)

	students["name"] = "Ashis"
	students["course"] = "React"

	fmt.Println(students)

	fmt.Println(students["course"])

	// Looping

	for key, v := range students {
		fmt.Printf("Key %v Value %v \n", key, v)
	}
}
