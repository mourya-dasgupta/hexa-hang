package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Global database connection pool
var DB *pgxpool.Pool

// InitDB connects to PostgreSQL and returns a connection pool
func InitDB() error {
	// Database connection string
	// Format: postgres://username:password@host:port/database
	connString := "postgres://hexauser:hexapass123@localhost:5432/hexahang?sslmode=disable"

	// Create connection pool
	var err error
	DB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Test the connection
	err = DB.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("unable to ping database: %w", err)
	}

	log.Println("✅ Database connected successfully!")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
