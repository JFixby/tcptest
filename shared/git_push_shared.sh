#!/bin/bash

set -e

VERSION="v0.1.1"
MODULE="shared"

echo "📦 Релиз модуля $MODULE с версией $VERSION"

# Коммитим изменения (если нужно)
git add .
git commit -m "release($MODULE): prepare $VERSION"

# Ставим тег в формате <module>/vX.Y.Z
git tag "$MODULE/$VERSION"

# Пушим
git push origin main
git push origin "$MODULE/$VERSION"

echo "✅ Готово! Тег $MODULE/$VERSION опубликован."
