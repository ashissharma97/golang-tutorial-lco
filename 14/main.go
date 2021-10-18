package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")

	result := adder(1, 3)

	fmt.Printf("Results is %v", result)

	proAdder, message := proAdder(1, 2, 5, 3, 7)

	fmt.Printf("Sum %v, Message %v \n", proAdder, message)

}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, v := range values {
		total += v
	}

	return total, "Success"
}
