# ===== Этап 1: Сборка приложения =====
FROM golang:1.21 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum и устанавливаем зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем остальной исходный код клиента и зависимые модули
COPY . .

# Сборка бинарника
RUN go build -o /build/wisdom_client main.go

# ===== Этап 2: Финальный образ =====
FROM alpine:latest

# Установка сертификатов (на случай TLS)
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копируем собранный бинарник из первого этапа
COPY --from=builder /build/wisdom_client .

# Точка входа — запуск клиента
ENTRYPOINT ["./wisdom_client"]
