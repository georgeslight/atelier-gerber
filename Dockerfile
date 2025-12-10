# FROM nixos/nix:latest AS builder

# RUN nix-channel --update

# WORKDIR /app

# COPY . .

# # Build
# RUN nix --extra-experimental-features "nix-command flakes" --option filter-syscalls false build .#default

# Runtime
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy pre-built binary from artifacts
COPY . /app/

# Copy built binary
# COPY --from=builder /app/result/ /app/
# Copy static files (images) from source
# COPY ./static /app/static/

# Render sets PORT environment variable
EXPOSE 3000
# EXPOSE $PORT

# Run the binary
CMD ["/app/bin/cmd"]
