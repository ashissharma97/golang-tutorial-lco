package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files in Golang")

	contents := "Some Dummy Things"

	file, err := os.Create("./text.txt")

	// if err != nil {
	// 	panic(err)
	// }

	checkNilError(err)

	length, err := io.WriteString(file, contents)

	checkNilError(err)

	fmt.Println("File length is:", length)

	data := readFile("./text.txt")

	fmt.Println("File contents are:", data)

}

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	checkNilError(err)
	return string(data)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
