package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	fmt.Println("Structs in Golang")

	Ashis := User{"Ashis", 23, "a@ashissharma.com"}

	fmt.Printf("User is %+v \n", Ashis)

	fmt.Printf("User Name is %v and email is %v", Ashis.Name, Ashis.Email)

}
