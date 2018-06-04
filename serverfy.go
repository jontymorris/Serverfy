package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Couldn not get the current directory")
	}

	// Make the file server
	fs := http.FileServer(http.Dir(dir))

	// Output pretty message
	fmt.Println(">>> Listening at localhost:8080")

	// Start the server
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
