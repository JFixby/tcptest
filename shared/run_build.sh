#!/bin/bash

set -e  # выходим при ошибке

echo "📦 Очистка зависимостей..."
go mod tidy

go test -v ./...
