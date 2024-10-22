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

# Use a separate stage to create a smaller image
FROM debian:bullseye

# Install required packages
RUN apt-get update && \
    apt-get install -y wget && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Install Neo4j
RUN wget -O neo4j.deb https://neo4j.com/artifact/binary/neo4j-community-5.24.2-unix.tar.gz && \
    tar -xvf neo4j-community-5.24.2-unix.tar.gz && \
    mv neo4j-community-5.24.2 /usr/local/neo4j && \
    rm neo4j.deb && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Set environment variables for Neo4j
ENV NEO4J_HOME=/usr/local/neo4j
ENV PATH="$NEO4J_HOME/bin:$PATH"

# Copy the built binary from the builder stage
COPY --from=builder /app/myapp /usr/local/bin/myapp

# Expose the ports
EXPOSE 4040
EXPOSE 7474 7687 7473

# Command to run the application
CMD ["myapp"]