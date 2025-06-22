package handlers

import (
	"io"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

// returns a pointer to new Hello struct by setting the log field as passed logger
// This is similar to a constructor in other languages
// In Go, we typically use a function to create and return a new instance of a struct
func NewHello(l *log.Logger) *Hello {
	return &Hello{log:l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Log the request
	//instead of log.Println, we use h.log.Println
	// This allows us to use the logger passed to the newHello function
	h.log.Println("Hello, World!\n")

	data, error := io.ReadAll(r.Body)
	if error != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	w.Write([]byte("hello " + string(data)))
}
