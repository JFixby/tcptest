# Этап сборки
FROM golang:1.20 AS builder

WORKDIR /app

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка сервера
RUN go build -o server-app ./cmd/server

# Финальный образ
FROM alpine:latest

# Добавим сертификаты (если планируется TLS)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем скомпилированное приложение
COPY --from=builder /app/server-app .
COPY server/wisdoms.json ./wisdoms.json

# Открываем порт (по желанию, для docker-compose)
EXPOSE 1337

# Запуск с флагами
ENTRYPOINT ["./server-app"]
