package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ashissharma97/mongoapi/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server Starting on port 4000")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Server Listening on port 4000")
}
