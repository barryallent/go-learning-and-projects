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
	gh := handlers.NewGoodbye(l)


	// Create a new HTTP server and register the handler
	http.Handle("/", hh)
	http.Handle("/goodbye", gh)

	// Start the server on port 9080
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)  
	// This will block until the server is stopped
	// The server will listen on port 9080 and handle requests using the handler we defined

	http.ListenAndServe(":9080", sm)
}
