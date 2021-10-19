package main

import (
	"fmt"
	"net/url"
)

const uri string = "https://ashissharma.com:3000/payment?paymentid=123432423&name=ashis"

func main() {
	fmt.Println("URL in golang")
	parsedUrl, _ := url.Parse(uri)
	fmt.Println(parsedUrl.RawQuery)
	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Port())

	query := parsedUrl.Query()

	fmt.Printf("Type of Query %T\n", query)

	for i, v := range query {
		fmt.Println(i, v)
	}

	newURL := &url.URL{
		Scheme: "https",
		Host:   "ashissharma.com",
		Path:   "/hello",
	}

	fmt.Println(newURL)
}
