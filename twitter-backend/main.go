package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raj5036/twitter-2024/router"
)

func main() {
	fmt.Println("Twitter Server is running successfully")

	// Setup routers
	router := router.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
	fmt.Println("Listening on PORT: 4000...........")

	// Setup MongoDB
	// password := "P@ssw0rd"
	// bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	// fmt.Println(string(bytes))
	// hashedPassword := string(bytes)

	// err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	// fmt.Println(err == nil)
}
