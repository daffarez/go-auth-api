# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o go-auth-api ./cmd/server

# Stage 2: Runtime
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/go-auth-api .
COPY .env .

EXPOSE 8088
CMD ["./go-auth-api"]
