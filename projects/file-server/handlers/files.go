package handlers

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"file-server/files"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

// Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

// ServeHTTP implements the http.Handler interface
func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	// no need to check for invalid id or filename as the mux router will not send requests
	// here unless they have the correct parameters

	f.saveFile(id, fn, rw, r)
}

// ListFiles returns a list of all files in the storage
func (f *Files) ListFiles(rw http.ResponseWriter, r *http.Request) {
	f.log.Info("Handle GET /files - listing all files")

	files, err := f.store.ListFiles()
	if err != nil {
		f.log.Error("Unable to list files", "error", err)
		http.Error(rw, "Unable to list files", http.StatusInternalServerError)
		return
	}

	// Set content type to JSON
	rw.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	err = json.NewEncoder(rw).Encode(files)
	if err != nil {
		f.log.Error("Unable to encode response", "error", err)
		http.Error(rw, "Unable to encode response", http.StatusInternalServerError)
		return
	}
}

// DeleteFile deletes a file from storage
func (f *Files) DeleteFile(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle DELETE", "id", id, "filename", fn)

	// Construct the file path
	filePath := filepath.Join(id, fn)

	// Delete the file
	err := f.store.DeleteFile(filePath)
	if err != nil {
		f.log.Error("Unable to delete file", "error", err)
		http.Error(rw, "Unable to delete file", http.StatusInternalServerError)
		return
	}

	// Return success response
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]string{
		"message":  "File deleted successfully",
		"id":       id,
		"filename": fn,
	})
}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
