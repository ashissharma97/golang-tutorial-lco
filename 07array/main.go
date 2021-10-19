package main

import "fmt"

func main() {
	fmt.Println("Arrays in Golang")

	var fruits [4]string

	fruits[0] = "Banana"
	fruits[1] = "Apple"
	fruits[2] = "Orange"
	fruits[3] = "Apricot"

	fmt.Println("Fruits", fruits)

	var veg = [2]string{"beans", "cabbage"}

	fmt.Println("Veg", veg)

	fmt.Println("Veg Length", len(veg))

}
