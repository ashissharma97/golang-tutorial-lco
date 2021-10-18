package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("User Input in Golang")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter any number:")

	input, _ := reader.ReadString('\n')

	fmt.Println("You entered", input)
}
