#!/bin/bash
set -e

echo "🚀 Setting up all Go microservices..."

# Daftar service
services=("user-service" "auth-service" "gateway")

# Prefix nama module
MODULE_PREFIX="github.com/bijoaja/GoIQ-microservices.v2"

for service in "${services[@]}"; do
  echo ""
  echo "🔧 Setting up $service..."
  cd "$service"

  # 1️⃣ Hapus file go.mod jika ada
  if [ -f "go.mod" ]; then
    echo "🧹 Removing old go.mod..."
    rm go.mod
  fi

  # 2️⃣ Hapus file go.sum jika ada
  if [ -f "go.sum" ]; then
    echo "🧹 Removing old go.sum..."
    rm go.sum
  fi

  # 3️⃣ Jalankan go mod init dan go mod tidy
  echo "🆕 Initializing new Go module for $service..."
  go mod init "$MODULE_PREFIX/$service"
  
  echo "📦 Running go mod tidy..."
  go mod tidy

  echo "✅ $service module setup complete."
  cd ..
done

echo ""
echo "🎉 All microservices have been successfully initialized and dependencies installed!"
