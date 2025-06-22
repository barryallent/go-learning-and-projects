package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"web-server/handlers"
	"time"
)

func main() {
	// Create a logger that writes to stdout with a prefix and timestamp
	l := log.New(os.Stdout, "hello-api ", log.LstdFlags)
	
	// Initialize handler instances with the logger
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	// Create a new router (ServeMux) to map URLs to handlers
	// ServeHTTP function of each handler will be called when the URL matches
	// For example, hh.ServeHTTP will be called for requests to "/"
	// and gh.ServeHTTP will be called for requests to "/goodbye"
	sm := http.NewServeMux()
	sm.Handle("/", hh)          // Maps "/" to Hello handler
	sm.Handle("/goodbye", gh)   // Maps "/goodbye" to Goodbye handler

	// Define the custom HTTP server configuration
	s := &http.Server{
		Addr:         ":9080",           // Server will listen on port 9080
		Handler:      sm,                // Use our custom router
		IdleTimeout:  120 * time.Second, // Max idle time before disconnect
		ReadTimeout:  1 * time.Second,   // Max time to read a request
		WriteTimeout: 1 * time.Second,   // Max time to write a response
	}

	// Start the HTTP server in a new goroutine
	// Goroutine is a lightweight thread in Go — it allows the server to run in the background
	// so the main thread can wait for OS signals (like Ctrl+C)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal("Error starting server: ", err)
		}
	}()

	// Create a channel to receive OS signals (like Interrupt or Kill)
	sigChain := make(chan os.Signal)

	// signal.Notify will send OS interrupt/kill signals into sigChain
	// This is how we detect Ctrl+C or program termination
	signal.Notify(sigChain, os.Interrupt)
	signal.Notify(sigChain, os.Kill)

	// Wait here until we receive a shutdown signal
	// This blocks the main thread until a signal is received
	sig := <-sigChain
	l.Println("Received terminate, Graceful shutdown", sig)

	// Create a context with timeout of 30 seconds to allow graceful shutdown
	// This gives running requests a chance to complete before server exits
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// Shutdown the server gracefully using the context timeout
	s.Shutdown(tc)

	log.Println("Server stopped gracefully")
}
