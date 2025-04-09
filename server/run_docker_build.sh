#!/bin/bash

set -e  # остановка при ошибке

if ! docker info > /dev/null 2>&1; then
  echo "❌ Docker не запущен."
  exit 1
fi

echo "🐳 Сборка Docker-образа сервера..."
docker build -t wisdom-server .

echo "✅ Готово: образ wisdom-server собран!"
