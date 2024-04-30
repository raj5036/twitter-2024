package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raj5036/twitter-2024/router"
)

func main() {
	fmt.Println("Twitter Server is starting....")

	// Setup routers
	router := router.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
	fmt.Println("Listening on PORT: 4000...........")
}
