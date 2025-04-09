# ===== Этап 1: Сборка =====
FROM golang:1.21 AS builder

# Рабочая директория
WORKDIR /app

# Копируем go.mod и go.sum — устанавливаем зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем всё, включая исходники и wisdoms.json
COPY . .

# cобираем
RUN go build -o /build/wisdom_server main.go

# ===== Этап 2: Финальный образ =====
FROM alpine:latest

# Добавим ca-certificates на случай TLS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копируем бинарник и JSON-файл с мудростями
COPY --from=builder /build/wisdom_server .
COPY wisdoms.json .

# Открываем порт (если нужно)
EXPOSE 1337

# Запуск
ENTRYPOINT ["./wisdom_server"]
