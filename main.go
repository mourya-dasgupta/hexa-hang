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

	// Serve static files (CSS, JS, images)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve the homepage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("static", "index.html"))
	})

	// API endpoint - health check
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "healthy", "service": "Hexa Hang", "database": "connected"}`)
	})

	fmt.Println("🚀 Hexa Hang server running on http://localhost:8080")
	fmt.Println("📊 Database: Connected")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
