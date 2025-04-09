# ===== Этап 1: Сборка приложения =====
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -o wisdom_client main.go

# ===== Этап 2: Финальный образ =====
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/wisdom_client .

ENTRYPOINT ["./wisdom_client"]
