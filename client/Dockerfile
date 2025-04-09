# Этап сборки
FROM golang:1.20 AS builder

WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники
COPY . .

# Собираем приложение
RUN go build -o client-app ./cmd/client

# Финальный минимальный образ (можно использовать scratch или alpine)
FROM alpine:latest

# Устанавливаем ca-certificates (если нужно TLS)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем скомпилированное приложение из билдера
COPY --from=builder /app/client-app .

# Указываем точку входа
ENTRYPOINT ["./client-app"]
