FROM golang:1.25.0 AS builder

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код бэкенда
COPY backend/ ./backend/

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./backend/cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder stage
COPY --from=builder /app/server .

# Копируем фронтенд
COPY frontend/ ./frontend

EXPOSE 8082

CMD ["./server","--port","8082"]