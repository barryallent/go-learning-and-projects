package config

import (
	"os"
	"strconv"
)

// Fileconfig holds file server config parameters
type FILECONFIG struct {
	bindAddress string
	logLevel    string
	basePath    string
}

func LoadConfig() *FILECONFIG {
	return &FILECONFIG{
		bindAddress: getEnv("bindAddress", ":9095"),
		logLevel:    getEnv("logLevel", "debug"),
		basePath:    getEnv("basePath", "./imagestore"),
	}
}

// Getter methods for FILECONFIG struct
func (f *FILECONFIG) GetBindAddress() string {
	return f.bindAddress
}

func (f *FILECONFIG) GetLogLevel() string {
	return f.logLevel
}

func (f *FILECONFIG) GetBasePath() string {
	return f.basePath
}

// Setter methods for FILECONFIG struct
func (f *FILECONFIG) SetBindAddress(addr string) {
	f.bindAddress = addr
}

func (f *FILECONFIG) SetLogLevel(level string) {
	f.logLevel = level
}

func (f *FILECONFIG) SetBasePath(path string) {
	f.basePath = path
}

// getEnv gets environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets environment variable as integer or returns default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
