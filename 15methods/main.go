package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	fmt.Println("Methods in Golang")

	Ashis := User{"Ashis", 23, "a@ashissharma.com"}

	fmt.Printf("User is %+v \n", Ashis)

	fmt.Printf("User Name is %v and email is %v \n", Ashis.Name, Ashis.Email)

	Ashis.GetName()

}

func (u User) GetName() {
	fmt.Printf("Name of User is %v", u.Name)
}
