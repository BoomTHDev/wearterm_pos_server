# Build stage
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-w -s' -o main .

# Final stage
FROM alpine:3.19

WORKDIR /app

# Install CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary and config from builder
COPY --from=builder /app/main .
COPY --from=builder /app/config.docker.yaml /app/config/config.yaml

# Create and set permissions for config directory
RUN mkdir -p /app/config && \
    chmod -R 755 /app/config

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./main"]