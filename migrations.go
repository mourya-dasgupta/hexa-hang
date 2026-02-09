package main

import (
	"context"
	"log"
)

// RunMigrations creates all database tables
func RunMigrations() error {
	log.Println("🔄 Running database migrations...")

	// Create products table
	createProductsTable := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		slug VARCHAR(255) UNIQUE NOT NULL,
		description TEXT NOT NULL,
		price INTEGER NOT NULL,
		capacity VARCHAR(100),
		features TEXT[],
		image_url VARCHAR(500),
		is_featured BOOLEAN DEFAULT false,
		stock_quantity INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(context.Background(), createProductsTable)
	if err != nil {
		return err
	}

	log.Println("✅ Products table created")
	return nil
}

// SeedProducts inserts sample products if table is empty
func SeedProducts() error {
	log.Println("🌱 Checking if products need to be seeded...")

	// Check if products already exist
	var count int
	err := DB.QueryRow(context.Background(), "SELECT COUNT(*) FROM products").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		log.Printf("ℹ️  Products already exist (%d products), skipping seed\n", count)
		return nil
	}

	// Insert sample products
	insertQuery := `
	INSERT INTO products (name, slug, description, price, capacity, features, image_url, is_featured, stock_quantity)
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9),
		($10, $11, $12, $13, $14, $15, $16, $17, $18),
		($19, $20, $21, $22, $23, $24, $25, $26, $27)
	`

	_, err = DB.Exec(context.Background(), insertQuery,
		// Product 1: Hexa Classic
		"Hexa Classic",
		"hexa-classic",
		"Our signature hexagonal design. Perfect for 10-15 medals. Wall-mounted elegance that transforms your achievements into art.",
		249900, // ₹2,499 in paise
		"10-15 medals",
		[]string{"Premium wood finish", "Easy wall mount", "Minimal design", "Durable construction"},
		"/static/images/hexa-classic.jpg",
		true,
		50,

		// Product 2: Hexa Pro
		"Hexa Pro",
		"hexa-pro",
		"Extended capacity for serious athletes. Holds 20+ medals with premium finish. Perfect for those with an extensive collection.",
		399900, // ₹3,999 in paise
		"20+ medals",
		[]string{"Extended capacity", "Premium metal hooks", "Reinforced backing", "Matte black finish"},
		"/static/images/hexa-pro.jpg",
		true,
		30,

		// Product 3: Hexa Elite
		"Hexa Elite",
		"hexa-elite",
		"The ultimate display. LED backlighting included. A statement piece for champions who want their achievements to shine.",
		599900, // ₹5,999 in paise
		"25+ medals",
		[]string{"LED backlighting", "Remote controlled", "Premium glass front", "Exhibition quality", "Wireless charging"},
		"/static/images/hexa-elite.jpg",
		true,
		20,
	)

	if err != nil {
		return err
	}

	log.Println("✅ Sample products seeded successfully")
	return nil
}
