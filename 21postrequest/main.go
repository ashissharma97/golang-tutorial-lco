package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const myUrl = "http://localhost:8000/post"

func main() {
	fmt.Println("Post request in Golang")

	// PostJSONRequest()

	// postFormRequest()

	PostFromDataRequest()

}

func PostJSONRequest() {

	requestBody := strings.NewReader(`
		{
			"name":"Ashis",
			"designation":"Full Stack Dev"
		}
	`)
	resp, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	resposeData, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(resposeData))

}

func postFormRequest() {
	data := url.Values{}
	data.Add("firstName", "Ashis")
	data.Add("surName", "Sharma")
	resp, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	responseData, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(responseData))
}

func PostFromDataRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "ashis")
	data.Add("lastname", "sharma")
	data.Add("email", "ashis@gmail.com")

	resp, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
}
