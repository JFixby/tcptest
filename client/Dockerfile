# Используем минимальный официальный образ Go
FROM golang:1.21-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарник
RUN go build -o client-app ./main.go

# Аргументы командной строки по умолчанию
ENTRYPOINT ["./client-app"]
