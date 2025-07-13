# File Server

A simple HTTP file server for uploading and downloading files.

## Features

- Upload files via HTTP POST
- Download files via HTTP GET
- List all files on the server
- Delete files from the server
- Configurable storage location
- Structured logging
- CORS enabled for web frontends

## Configuration

Set these environment variables:

```bash
bindAddress=:9095          # Server bind address
logLevel=debug             # Log level (debug, info, warn, error)
basePath=./imagestore      # Directory to store files
```

## Running

```bash
go run main.go
```

Server starts on `:9095` by default.

## Endpoints

### Upload File
```bash
curl http://localhost:9095/images/1/filename.txt -d @local-file.txt
```

### Download File
```bash
curl http://localhost:9095/images/1/filename.txt
```

### List All Files
```bash
curl http://localhost:9095/files
```

### Delete File
```bash
curl -X DELETE http://localhost:9095/images/1/filename.txt
```

### Health Check
```bash
curl http://localhost:9095/health
```

## File Structure

```
./imagestore/
├── 1/
│   └── filename.txt
└── 2/
    └── another-file.jpg
```

Files are organized by ID in subdirectories under the base path.

## Response Format

### List Files Response
```json
[
  {
    "id": "1",
    "filename": "test.txt",
    "size": 1024,
    "path": "1/test.txt"
  }
]
```

### Delete Response
```json
{
  "message": "File deleted successfully",
  "id": "1",
  "filename": "test.txt"
}
``` 