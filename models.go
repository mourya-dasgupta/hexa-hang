package main

import "time"

// Product represents a medal holder product
type Product struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   string    `json:"description"`
	Price         int       `json:"price"`
	Capacity      string    `json:"capacity"`
	Features      []string  `json:"features"`
	ImageURL      string    `json:"image_url"`
	IsFeatured    bool      `json:"is_featured"`
	StockQuantity int       `json:"stock_quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
