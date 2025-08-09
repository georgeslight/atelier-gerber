# Fetch dependencies
FROM golang:latest AS fetch-stage
COPY go.mod go.sum ./
RUN go mod download

# Generate templ files FIRST
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# Build CSS assets AFTER templ generation
FROM oven/bun:1.1-slim AS css-stage
WORKDIR /app

# Copy package files and install dependencies
COPY package.json bun.lockb* ./
RUN bun install

# Copy ALL generated files from previous stage
COPY --from=generate-stage /app ./

# Now Tailwind scans the generated .go files, not just .templ files
RUN bunx tailwindcss --input ./static/css/input.css --output ./static/css/styles.css --minify

# Build Go application
FROM golang:latest AS build-stage
COPY --from=fetch-stage /go/pkg /go/pkg
COPY --from=css-stage /app /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o /app/app ./cmd

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