package main

import (
	"context"
	"file-server/config"
	"file-server/files"
	"file-server/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

func main() {
	// Load configuration from environment variables
	cfg := config.LoadConfig()

	// Get config values
	var bindAddress = cfg.GetBindAddress()
	var logLevel = cfg.GetLogLevel()
	var basePath = cfg.GetBasePath()

	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images",
			Level: hclog.LevelFromString(logLevel),
		},
	)

	// create a logger for the server from the default logger
	sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	// create the storage class, use local storage
	// max filesize 5MB
	stor, err := files.NewLocalStorage(basePath, 1024*1000*5)
	if err != nil {
		l.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}

	// create the handlers
	fh := handlers.NewFiles(stor, l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// filename regex: {filename:[a-zA-Z]+\\.[a-z]{3}}
	// problem with FileServer is that it is dumb.
	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename}", fh.ServeHTTP)

	// get files
	//use this curl to upload file-  curl http://localhost:9095/images/1/file-copied.txt -d @file-to-copy.txt
	gh := sm.Methods(http.MethodGet).Subrouter()

	//using FileServer hander here, because that is already available. Trimming prefix /images because
	//we have given the base path and inside base path we have id and then file name, there is no /images in base folder
	//use curl http://localhost:9095/images/1/file-postman.txt to get the file content
	gh.Handle(
		"/images/{id:[0-9]+}/{filename}",
		http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))),
	)

	// Health check endpoint
	sm.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "service": "file-server"}`))
	}).Methods(http.MethodGet)

	// List all files endpoint
	sm.HandleFunc("/files", fh.ListFiles).Methods(http.MethodGet)

	// Delete file endpoint
	dh := sm.Methods(http.MethodDelete).Subrouter()
	dh.HandleFunc("/images/{id:[0-9]+}/{filename}", fh.DeleteFile)

	//CORS middleware to allow cross-origin requests from browsers
	ch := gohandlers.CORS(
		// allow all origins, methods, and headers
		gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		//allowed 2 headers for requests
		gohandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// create a new server
	s := http.Server{
		Addr:         bindAddress,       // configure the bind address
		Handler:      ch(sm),            // set the default handler
		ErrorLog:     sl,                // the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Info("Starting server", "bind_address", bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Unable to start server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	l.Info("Shutting down server with", "signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
