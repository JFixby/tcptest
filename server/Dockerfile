# ===== Этап 1: Сборка =====
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# ВАЖНО: отключаем CGO, чтобы собрать статически
ENV CGO_ENABLED=0
RUN go build -o wisdom_server main.go

# ===== Этап 2: Финальный образ =====
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/wisdom_server .
COPY wisdoms.json .

EXPOSE 1337

ENTRYPOINT ["./wisdom_server"]
