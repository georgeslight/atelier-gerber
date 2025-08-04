# Fetch dependencies
FROM golang:latest AS fetch-stage
COPY go.mod go.sum ./
RUN go mod download

# Build CSS assets
FROM oven/bun:1.1-slim AS css-stage
WORKDIR /app

# Copy package files and install dependencies
COPY package.json bun.lockb* ./
RUN bun install

# Copy source files for Tailwind to scan
COPY views/ ./views/
COPY internal/ ./internal/
COPY static/ ./static/
#COPY tailwind.config.js ./

# Build CSS with Tailwind scanning all template files
RUN bunx tailwindcss --input ./static/css/input.css --output ./static/css/styles.css --minify

# Generate templ files
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# Build Go application
FROM golang:latest AS build-stage
COPY --from=fetch-stage /go/pkg /go/pkg
COPY --from=generate-stage /app /app
COPY --from=css-stage /app/static/css/styles.css /app/static/css/styles.css
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app ./cmd

# Test
FROM build-stage AS test-stage
RUN go test -v ./...

# Deploy
FROM gcr.io/distroless/base-debian12 AS deploy-stage
WORKDIR /
COPY --from=build-stage /app/app /app
COPY --from=build-stage /app/static /static
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/app"]