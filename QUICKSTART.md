# Quick Start Guide

Follow these steps to get Hexa Hang running locally in 5 minutes.

---

## ⚡ Quick Setup (5 Minutes)

### 1. Download & Extract Files

Download all files to your local machine and extract them to a folder called `hexa-hang`.

### 2. Open in Cursor

```bash
cd hexa-hang
cursor .
```

Or: Open Cursor → File → Open Folder → Select `hexa-hang`

### 3. Install Dependencies

In Cursor's terminal (`Cmd+J` or `Ctrl+J`):

```bash
# Install Go dependencies
go mod download
```

### 4. Setup Database (Automatic)

```bash
# Run the setup script
./setup.sh
```

**Or manually:**

```bash
# Create database
createdb hexahang

# Create user
psql postgres -c "CREATE USER hexauser WITH PASSWORD 'hexapass123';"

# Grant privileges
psql postgres -c "GRANT ALL PRIVILEGES ON DATABASE hexahang TO hexauser;"
psql -d hexahang -c "GRANT ALL ON SCHEMA public TO hexauser;"
```

### 5. Run the Application

```bash
go run *.go
```

**Expected output:**
```
2024/02/09 12:00:00 ✅ Database connected successfully!
🚀 Hexa Hang server running on http://localhost:8080
📊 Database: Connected
```

### 6. Visit Your Website

Open browser: **http://localhost:8080**

You should see your premium Hexa Hang landing page! 🎉

---

## ✅ Verify Everything Works

### Test the API:

```bash
curl http://localhost:8080/api/health
```

**Expected response:**
```json
{"status": "healthy", "service": "Hexa Hang", "database": "connected"}
```

---

## 🎓 Next Steps

### Initialize Git:

```bash
git init
git add .
git commit -m "Initial commit: Hexa Hang MVP"
```

### Optional - Push to GitHub:

1. Create a new repo on GitHub
2. Run:
   ```bash
   git remote add origin https://github.com/YOUR_USERNAME/hexa-hang.git
   git branch -M main
   git push -u origin main
   ```

---

## 🐛 Troubleshooting

### PostgreSQL not installed?

```bash
# Install PostgreSQL
brew install postgresql@14

# Start service
brew services start postgresql@14
```

### Port 8080 already in use?

```bash
# Find what's using it
lsof -i :8080

# Kill the process or change port in main.go
```

### Database connection error?

```bash
# Check if PostgreSQL is running
brew services list | grep postgresql

# Restart if needed
brew services restart postgresql@14
```

---

## 📚 File Overview

- **main.go** - HTTP server & routing
- **db.go** - Database connection logic
- **go.mod** - Dependencies
- **static/index.html** - Frontend
- **setup.sh** - Database setup script
- **.gitignore** - Git ignore rules

---

**All set?** Reply "Ready for Phase 3" and we'll create your first database table! 🚀
