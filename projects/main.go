package main

import (
	"log"
	"net/http"
	"os"
	"projects/handlers"
)

func main() {
	l := log.New(os.Stdout, "hello-api ", log.LstdFlags)
	hh := handlers.NewHello(l)
	// Create a new HTTP server and register the handler
	http.Handle("/", hh)
	// Start the server on port 9080
	log.Println("Starting server on :9080")
	// This will block until the server is stopped
	// The server will listen on port 9080 and handle requests using the handler we defined

	http.ListenAndServe(":9080", nil)
}
