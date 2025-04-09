#!/bin/bash

set -e

# Проверка аргументов
if [ "$#" -ne 2 ]; then
  echo "❌ Использование: ./release_module.sh <module_name> <version>"
  echo "Пример: ./release_module.sh shared v0.1.0"
  exit 1
fi

MODULE="$1"
VERSION="$2"
TAG="$MODULE/$VERSION"

echo "📦 Релиз модуля '$MODULE' с версией '$VERSION'"

# Проверка существования директории модуля
if [ ! -d "$MODULE" ]; then
  echo "❌ Директория '$MODULE' не найдена"
  exit 1
fi

# Создание коммита
git add "$MODULE"
git commit -m "release($MODULE): prepare $VERSION"

# Создание и пуш тега
git tag "$TAG"
git push origin "$TAG"

echo "✅ Готово! Тег '$TAG' запушен в GitHub"
