package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct{
	log *log.Logger	
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{log:l} 
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goodbye, World!\n"))
} 