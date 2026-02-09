package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// Initialize database connection
	err := InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer CloseDB()

	// Run migrations (create tables)
	err = RunMigrations()
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Seed products
	err = SeedProducts()
	if err != nil {
		log.Fatalf("Failed to seed products: %v", err)
	}

	// Serve static files (CSS, JS, images)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve the homepage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("static", "index.html"))
	})

	// API endpoints
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "healthy", "service": "Hexa Hang", "database": "connected"}`)
	})

	// Products API - List all products
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetProducts(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Single product by slug
	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetProductBySlug(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("🚀 Hexa Hang server running on http://localhost:8080")
	fmt.Println("📊 Database: Connected")
	fmt.Println("✨ API Endpoints:")
	fmt.Println("   GET /api/health")
	fmt.Println("   GET /api/products")
	fmt.Println("   GET /api/products/:slug")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
