package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Time in Golang")
	currentTime := time.Now()
	fmt.Println(currentTime)

	// Format Doc https://pkg.go.dev/time#example-Time.Format

	formatedTime := currentTime.Format("02-01-2006, 15:04:05")

	fmt.Println(formatedTime)

	createdDate := time.Date(1997, time.November, 22, 12, 30, 0, 0, time.Local)

	fmt.Println(createdDate)
}
