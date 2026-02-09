package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// GetProducts returns all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Query all products from database
	query := `
		SELECT id, name, slug, description, price, capacity, features, 
		       image_url, is_featured, stock_quantity, created_at, updated_at
		FROM products
		ORDER BY id ASC
	`

	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Collect all products
	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(
			&p.ID, &p.Name, &p.Slug, &p.Description, &p.Price,
			&p.Capacity, &p.Features, &p.ImageURL, &p.IsFeatured,
			&p.StockQuantity, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			http.Error(w, "Failed to scan product", http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	// Return as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProductBySlug returns a single product by slug
func GetProductBySlug(w http.ResponseWriter, r *http.Request) {
	// Extract slug from URL path
	// URL format: /api/products/hexa-classic
	slug := strings.TrimPrefix(r.URL.Path, "/api/products/")

	if slug == "" {
		http.Error(w, "Product slug required", http.StatusBadRequest)
		return
	}

	// Query single product
	query := `
		SELECT id, name, slug, description, price, capacity, features, 
		       image_url, is_featured, stock_quantity, created_at, updated_at
		FROM products
		WHERE slug = $1
	`

	var p Product
	err := DB.QueryRow(context.Background(), query, slug).Scan(
		&p.ID, &p.Name, &p.Slug, &p.Description, &p.Price,
		&p.Capacity, &p.Features, &p.ImageURL, &p.IsFeatured,
		&p.StockQuantity, &p.CreatedAt, &p.UpdatedAt,
	)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Return as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
