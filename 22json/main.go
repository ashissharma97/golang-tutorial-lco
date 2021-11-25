package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json in Golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React JS", 200, "Udemy", "123", []string{"JS", "FB"}},
		{"Angular", 200, "LCO", "123", []string{"TS", "Google"}},
		{"Node JS", 200, "Cousera", "123", nil},
	}

	jsonData, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React JS",
		"Price": 200,
		"website": "Udemy",
		"tags": [
				"JS",
				"FB"
		]
	}`)

	var aCourse course

	checkValidJson := json.Valid(jsonDataFromWeb)

	if checkValidJson {
		fmt.Println("This is a valid json")
		json.Unmarshal(jsonDataFromWeb, &aCourse)
		fmt.Printf("%#v\n", aCourse)
	} else {
		fmt.Println("JSON was invalid")
	}

	// storing json in key value from json

	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key is %v Value is %v and type is %t \n", k, v, v)
	}
}
