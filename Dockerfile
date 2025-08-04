# syntax = docker/dockerfile:1

# Build stage for Go application
FROM golang:1.21-alpine AS go-builder

WORKDIR /app

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate templ files and build Go binary
RUN templ generate && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Build stage for frontend assets
FROM oven/bun:1.1-slim AS frontend-builder

WORKDIR /app

# Copy package files
COPY package.json bun.lock* ./
RUN bun install

# Copy static files and build CSS
COPY static/ ./static/
RUN bunx @tailwindcss/cli -i ./static/css/input.css -o ./static/css/styles.css --minify

# Final stage
FROM alpine:latest

LABEL fly_launch_runtime="Go"

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy Go binary
COPY --from=go-builder /app/main .

# Copy static assets
COPY --from=frontend-builder /app/static ./static

EXPOSE 3000

CMD ["./main"]