package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Get request in Golang")

	GetRequest()

}

func GetRequest() {
	const url = "http://localhost:8000/"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(response))

	var responseString strings.Builder

	byteCount, _ := responseString.Write(response)

	fmt.Println(byteCount)

	fmt.Println(responseString.String())

}
