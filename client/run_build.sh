#!/bin/bash

set -e  # выходим при ошибке

echo "🧹 Удаление старых бинарников..."
mkdir -p bin
rm -f bin/*

echo "🎨 Форматирование кода..."
go fmt ./...

echo "📦 Очистка зависимостей..."
go mod tidy

echo "🔨 Сборка..."
go build -o bin/wisdom_client main.go

echo "✅ Готово! Бинарник: ./wisdom_client"
