# Use the official Golang image to build the application
# Use the official Golang image to build the application
FROM golang:1.22.5-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLE=1 go build -o myapp .

# Use a minimal image to run the application
FROM scratch

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Copy the config.yaml into the container
COPY internal/config/config.yaml ./internal/config/

# Expose the application port
EXPOSE 4040

# Start the application
CMD ["./myapp"]