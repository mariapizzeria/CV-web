FROM golang:1.21.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY backend/ ./backend/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/server ./

COPY frontend/ ./frontend

EXPOSE 8082

CMD ["./server","--port","8082"]