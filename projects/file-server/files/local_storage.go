package files

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"
)

// LocalStorage is an implementation of the Storage interface which works with the
// local disk on the current machine
type LocalStorage struct {
	maxFileSize int // maximum number of bytes for files
	basePath    string
}

// NewLocalStorage creates a new LocalStorage file-sytem with the given base path
// basePath is the base directory to save files to
// maxSize is the max number of bytes that a file can be
func NewLocalStorage(basePath string, maxSize int) (*LocalStorage, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &LocalStorage{basePath: p}, nil
}

// Save the contents of the Writer to the given path
// path is a relative path, basePath will be appended
func (l *LocalStorage) Save(path string, contents io.Reader) error {
	// get the full path for the file
	fp := l.fullPath(path)

	// get the directory and make sure it exists
	d := filepath.Dir(fp)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return xerrors.Errorf("Unable to create directory: %w", err)
	}

	// if the file exists delete it
	_, err = os.Stat(fp)
	if err == nil {
		err = os.Remove(fp)
		if err != nil {
			return xerrors.Errorf("Unable to delete file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		// if this is anything other than a not exists error
		return xerrors.Errorf("Unable to get file info: %w", err)
	}

	// create a new file at the path
	f, err := os.Create(fp)
	if err != nil {
		return xerrors.Errorf("Unable to create file: %w", err)
	}
	defer f.Close()

	// write the contents to the new file
	// ensure that we are not writing greater than max bytes
	_, err = io.Copy(f, contents)
	if err != nil {
		return xerrors.Errorf("Unable to write to file: %w", err)
	}

	return nil
}

// Get the file at the given path and return a Reader
// the calling function is responsible for closing the reader
func (l *LocalStorage) Get(path string) (*os.File, error) {
	// get the full path for the file
	fp := l.fullPath(path)

	// open the file
	f, err := os.Open(fp)
	if err != nil {
		return nil, xerrors.Errorf("Unable to open file: %w", err)
	}

	return f, nil
}

// ListFiles returns a list of all files in the storage
func (l *LocalStorage) ListFiles() ([]FileInfo, error) {
	var files []FileInfo

	// Walk through all directories and files
	err := filepath.Walk(l.basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories, only process files
		if !info.IsDir() {
			// Get relative path from basePath
			relativePath, err := filepath.Rel(l.basePath, path)
			if err != nil {
				return err
			}

			// Split the path to get ID and filename
			// Expected format: {id}/{filename}
			parts := strings.Split(relativePath, string(filepath.Separator))
			if len(parts) >= 2 {
				id := parts[0]
				filename := filepath.Base(relativePath)

				files = append(files, FileInfo{
					ID:       id,
					Filename: filename,
					Size:     info.Size(),
					Path:     relativePath,
				})
			}
		}

		return nil
	})

	if err != nil {
		return nil, xerrors.Errorf("Unable to list files: %w", err)
	}

	return files, nil
}

// DeleteFile deletes a file at the given path
func (l *LocalStorage) DeleteFile(path string) error {
	// get the full path for the file
	fp := l.fullPath(path)

	// check if file exists
	_, err := os.Stat(fp)
	if os.IsNotExist(err) {
		return xerrors.Errorf("File not found: %s", path)
	}
	if err != nil {
		return xerrors.Errorf("Unable to get file info: %w", err)
	}

	// delete the file
	err = os.Remove(fp)
	if err != nil {
		return xerrors.Errorf("Unable to delete file: %w", err)
	}

	return nil
}

// returns the absolute path
func (l *LocalStorage) fullPath(path string) string {
	// append the given path to the base path
	return filepath.Join(l.basePath, path)
}
