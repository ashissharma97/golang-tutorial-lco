package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Conversions in Go Lang")

	reader := bufio.NewReader(os.Stdin)

	value, _ := reader.ReadString('\n')

	fmt.Println("You entered", value)
	fmt.Printf("You entered a %T value \n", value)

	convertedValue, err := strconv.ParseInt(strings.TrimSpace(value), 10, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully converted input value", convertedValue)
		fmt.Printf("You entered a %T value \n", convertedValue)
	}

}
