package config

import (
	"os"
	"strconv"
)

// AppConfig holds application configuration
type AppConfig struct {
	DatabaseConfig DatabaseConfig
	ServerConfig   ServerConfig
}

// DatabaseConfig holds database connection parameters
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port int
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() *AppConfig {
	return &AppConfig{
		DatabaseConfig: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5436),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "admin"),
			DBName:   getEnv("DB_NAME", "product_api"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		ServerConfig: ServerConfig{
			Port: getEnvAsInt("SERVER_PORT", 9080),
		},
	}
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
