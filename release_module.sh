#!/bin/bash
set -e

MODULE="$1"
VERSION="$2"
TAG="$MODULE/$VERSION"

if [ -z "$MODULE" ] || [ -z "$VERSION" ]; then
  echo "❌ Использование: ./release_module.sh <module> <version>"
  exit 1
fi

echo "📦 Релиз модуля '$MODULE' с версией '$VERSION'"

git add "$MODULE"

if git diff --cached --quiet; then
  echo "ℹ️ Нет изменений в $MODULE. Пропускаем коммит."
else
  git commit -m "release($MODULE): prepare $VERSION"
fi

if git rev-parse "$TAG" >/dev/null 2>&1; then
  echo "⚠️ Тег $TAG уже существует"
else
  git tag "$TAG"
  git push origin "$TAG"
  echo "✅ Тег $TAG создан и запушен"
fi
