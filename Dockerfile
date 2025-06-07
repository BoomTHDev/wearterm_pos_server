FROM golang:1.24.4-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Ensure config directory exists and has the right permissions
RUN mkdir -p /app/config && \
    chmod -R 755 /app/config

# Build the application
RUN go build -o main .

# Expose port
EXPOSE 8080

# Copy the docker config file
COPY config.docker.yaml /app/config/config.yaml

# Command to run the application
CMD ["./main"]