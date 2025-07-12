package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type DB struct {
	*sql.DB
}

// Config holds database configuration
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewConnection creates a new database connection
func NewConnection(config Config) (*DB, error) {
	// Create connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	// Open database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	return &DB{db}, nil
}

// CreateTables creates the necessary tables if they don't exist
func (db *DB) CreateTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
		sku VARCHAR(50) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL
	);
	
	-- Create index on SKU for faster lookups
	CREATE INDEX IF NOT EXISTS idx_products_sku ON products(sku);
	
	-- Insert sample data if table is empty
	INSERT INTO products (name, description, price, sku) 
	SELECT 'Latte', 'frothy coffee with steamed milk', 2.50, 'SKU-001'
	WHERE NOT EXISTS (SELECT 1 FROM products WHERE sku = 'SKU-001');
	
	INSERT INTO products (name, description, price, sku) 
	SELECT 'Espresso', 'strong coffee shot', 1.50, 'SKU-002'
	WHERE NOT EXISTS (SELECT 1 FROM products WHERE sku = 'SKU-002');
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	log.Println("Database tables created successfully")
	return nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.DB.Close()
}
