# Hexa Hang - Premium Medal Holders

**Show What You've Earned**

A minimal, premium D2C brand for medal holders. Built with Go backend and clean, Apple-inspired frontend.

---

## 🎯 What's Included

✅ **Premium Frontend** - Apple-inspired minimal design with:
- Elegant serif + sans-serif typography (Crimson Pro + DM Sans)
- Smooth scroll animations
- Responsive mobile design
- Gold accent theme (matches medal/achievement brand)
- Product showcase grid
- Editorial-style about section

✅ **Go Backend** with:
- PostgreSQL database integration
- Connection pooling with pgx
- RESTful API structure
- Static file serving
- Health check endpoint

---

## 🚀 Local Development Setup

### Prerequisites

- Go 1.22+ installed
- PostgreSQL 14+ installed locally
- Git installed

### Step 1: Clone/Setup Project

```bash
# Navigate to your project directory
cd hexa-hang

# Install Go dependencies
go mod download
```

### Step 2: Setup Local Database

```bash
# Create database
createdb hexahang

# Create user
psql postgres -c "CREATE USER hexauser WITH PASSWORD 'hexapass123';"

# Grant privileges
psql postgres -c "GRANT ALL PRIVILEGES ON DATABASE hexahang TO hexauser;"
psql -d hexahang -c "GRANT ALL ON SCHEMA public TO hexauser;"

# Verify
psql -d hexahang -c "SELECT current_database();"
```

### Step 3: Run the Application

```bash
# Start the server
go run *.go

# Or build and run
go build -o hexa-hang
./hexa-hang
```

**Visit:** http://localhost:8080

---

## 📁 Project Structure

```
hexa-hang/
├── main.go              # HTTP server & routing
├── db.go                # Database connection & functions
├── go.mod               # Go module dependencies
├── .gitignore           # Git ignore rules
└── static/
    └── index.html       # Frontend (HTML + CSS + JS)
```

---

## 🔧 Understanding the Code

### **db.go - Database Connection**

```go
// Connection string for local PostgreSQL
connString := "postgres://hexauser:hexapass123@localhost:5432/hexahang?sslmode=disable"

// Creates a connection pool (reuses connections)
DB, err = pgxpool.New(context.Background(), connString)

// Tests the connection
err = DB.Ping(context.Background())
```

**Why connection pooling?**
- Reuses database connections instead of creating new ones
- Much faster for concurrent requests
- Industry standard for production apps

### **main.go - HTTP Server**

```go
// Initialize database on startup
err := InitDB()
defer CloseDB()  // Close when program exits

// Serve static files
http.Handle("/static/", http.StripPrefix("/static/", fs))

// API endpoint
http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"status": "healthy"}`)
})
```

---

## ✅ Test Your API

```bash
# Health check
curl http://localhost:8080/api/health

# Expected response:
{"status": "healthy", "service": "Hexa Hang", "database": "connected"}
```

---

## 🎨 Frontend Features

- **Scroll animations**: Elements fade in as you scroll
- **Hover effects**: Product cards lift on hover with shimmer
- **Smooth scrolling**: Navigation links scroll smoothly
- **Mobile responsive**: Perfect on all devices
- **Premium aesthetic**: Gold accents, elegant typography

---

## 📝 Development Workflow

### Daily Development

1. **Make changes** in Cursor/VSCode
2. **Test locally:**
   ```bash
   go run *.go
   ```
3. **Commit changes:**
   ```bash
   git add .
   git commit -m "Description of changes"
   ```

### Before Committing

```bash
# Check for errors
go vet ./...

# Format code
go fmt ./...

# Test (when we add tests)
go test ./...
```

---

## 🎓 Go Concepts Learned So Far

1. ✅ **Modules**: `go.mod` defines your project
2. ✅ **HTTP Server**: `http.ListenAndServe()`
3. ✅ **Routing**: `http.HandleFunc()` maps URLs to functions
4. ✅ **Static Files**: `http.FileServer()` serves HTML/CSS/JS
5. ✅ **Database**: `pgxpool` for PostgreSQL connections
6. ✅ **Context**: `context.Background()` for database operations
7. ✅ **Error Handling**: Checking `err != nil`
8. ✅ **Defer**: `defer CloseDB()` runs when function exits
9. ✅ **Package Structure**: Multiple `.go` files in same package

---

## 📚 Next Steps (What We'll Build)

### Phase 3: First Database Table
- Create `products` table
- Write SQL migrations
- Learn database schema design

### Phase 4: Product API
- `GET /api/products` - List all products
- `GET /api/products/:id` - Get single product
- `POST /api/products` - Add new product (admin)
- Learn JSON marshaling/unmarshaling

### Phase 5: Order System
- Create `orders` table
- Create `order_items` table
- Relationships between tables
- Transaction handling

### Phase 6: Razorpay Integration
- Payment gateway setup
- Order + payment flow
- Webhook handling
- Payment verification

### Phase 7: Authentication
- User registration/login
- JWT tokens
- Protected routes
- Admin panel

### Phase 8: Deployment
- Deploy to VPS
- Environment variables
- Systemd service
- SSL certificate
- Go live! 🚀

---

## 🛠️ Customization Tips

### Change Product Names/Prices
Edit `static/index.html`, find the `.product-card` sections.

### Change Brand Colors
Edit the `:root` CSS variables in `index.html`:
```css
:root {
    --color-gold: #d4af37;  /* Change to your brand color */
}
```

### Add Your Logo
Replace the `.logo` text in the nav with an `<img>` tag.

---

## 💡 Pro Tips

1. **Always use `go run *.go`** (not just `go run main.go`)
2. **Hard refresh browser**: `Cmd+Shift+R` to clear cache
3. **Check terminal for errors**: Server logs appear there
4. **Use `go fmt`**: Auto-formats your code beautifully
5. **Read error messages**: Go errors are very descriptive

---

## 🐛 Troubleshooting

### Database connection refused
```bash
# Check if PostgreSQL is running
brew services list | grep postgresql

# Start if not running
brew services start postgresql@14
```

### Port 8080 already in use
```bash
# Find what's using the port
lsof -i :8080

# Kill the process or change port in main.go
```

### Import errors
```bash
# Download dependencies
go mod download

# Tidy up (removes unused dependencies)
go mod tidy
```

---

## 📖 Learning Resources

- [Go by Example](https://gobyexample.com/) - Quick Go syntax reference
- [pgx Documentation](https://pkg.go.dev/github.com/jackc/pgx/v5) - PostgreSQL driver docs
- [Effective Go](https://go.dev/doc/effective_go) - Go best practices

---

**Ready for Phase 3?** Reply "Ready" and we'll create your first database table! 🔥
