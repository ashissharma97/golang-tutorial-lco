package main

import "fmt"

/*
	Global variable cannot be declared with := operator
	can use const or var
*/

const User = "Ashis"

func main() {
	fmt.Println("Variables:-")

	// String
	var username string = "Ashis"
	fmt.Println(username)
	fmt.Printf("Variable of type: %T \n", username)

	// Boolean
	var isUser bool = false
	fmt.Println(isUser)
	fmt.Printf("Variable of type: %T \n", isUser)

	// Uint
	// https://golang.org/ref/spec#Numeric_types
	var valueUint uint8 = 255
	fmt.Println(valueUint)
	fmt.Printf("Variable of type: %T \n", valueUint)

	// Int
	var valueInt int = 10000000
	fmt.Println(valueInt)
	fmt.Printf("Variable of type: %T \n", valueInt)

	// Another way to declare variable without type

	var valueNoType = "100"
	fmt.Println(valueNoType)
	fmt.Printf("Variable of type: %T \n", valueNoType)

	// Another way to declare variable without type

	valueAutoType := 100
	fmt.Println(valueAutoType)
	fmt.Printf("Variable of type: %T \n", valueAutoType)

	fmt.Println(User)
}
