#!/bin/bash

echo "🚀 Hexa Hang - Database Setup Script"
echo "======================================"
echo ""

# Check if PostgreSQL is running
if ! brew services list | grep -q "postgresql@14.*started"; then
    echo "⚠️  PostgreSQL is not running. Starting it now..."
    brew services start postgresql@14
    sleep 2
fi

echo "✅ PostgreSQL is running"
echo ""

# Create database
echo "📊 Creating database 'hexahang'..."
createdb hexahang 2>/dev/null
if [ $? -eq 0 ]; then
    echo "✅ Database created successfully"
else
    echo "ℹ️  Database already exists"
fi

# Create user
echo "👤 Creating user 'hexauser'..."
psql postgres -c "CREATE USER hexauser WITH PASSWORD 'hexapass123';" 2>/dev/null
if [ $? -eq 0 ]; then
    echo "✅ User created successfully"
else
    echo "ℹ️  User already exists"
fi

# Grant privileges
echo "🔑 Granting privileges..."
psql postgres -c "GRANT ALL PRIVILEGES ON DATABASE hexahang TO hexauser;" >/dev/null 2>&1
psql -d hexahang -c "GRANT ALL ON SCHEMA public TO hexauser;" >/dev/null 2>&1
echo "✅ Privileges granted"

# Verify connection
echo ""
echo "🧪 Testing database connection..."
if psql -d hexahang -c "SELECT current_database();" >/dev/null 2>&1; then
    echo "✅ Database connection successful!"
    echo ""
    echo "======================================"
    echo "✨ Setup complete! You can now run:"
    echo "   go run *.go"
    echo "======================================"
else
    echo "❌ Connection test failed. Please check the errors above."
fi
