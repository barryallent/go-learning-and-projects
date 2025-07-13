package files

import "io"

// FileInfo represents basic file information
type FileInfo struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	Path     string `json:"path"`
}

// Storage defines the behavior for file operations
// Implementations may be of the time local disk, or cloud storage, etc
type Storage interface {
	Save(path string, file io.Reader) error
	ListFiles() ([]FileInfo, error)
	DeleteFile(path string) error
}
