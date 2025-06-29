package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"time"
)

func main() {
	// Create a logger that writes to stdout with a prefix and timestamp
	l := log.New(os.Stdout, "hello-api ", log.LstdFlags)

	// Initialize handler instances with the logger
	ph := handlers.NewProductsHandler(l)

	// using gorilla/mux for routing, its a powerful HTTP router and URL matcher for building Go web servers
	sm := mux.NewRouter()

	// using gorilla/mux, we can create subrouters for different HTTP methods
	getRouter := sm.Methods(http.MethodGet).Subrouter()

	// we are not passing any params when calling the GetProducts method
	// because handleFunc expects a function with the signature(w http.ResponseWriter, r *http.Request)
	// and GetProducts matches that signature
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/product/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/product", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	//sm.Handle("/", hh) // Maps "/" to Hello handler

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
