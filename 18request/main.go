package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const Example = "https://example.com"

func main() {
	fmt.Println("Web Request in Golang")

	resp, err := http.Get(Example)

	if err != nil {
		panic(err)
	}

	response, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(response))
}
