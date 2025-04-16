FROM golang:1.23.4-alpine AS builder

# Install git and required tools for Go modules
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy go mod files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the Go source code and .env file to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Build the migration binary
RUN go build -o migrate ./internal/migration/migration.go

# Create a new stage for the final image
FROM alpine:latest

# Install SSL certificates (needed for HTTPS or DB connections)
RUN apk add --no-cache ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /app/main /

COPY --from=builder /app/migrate /migrate

# Copy the .env file into the root directory of the final image
COPY --from=builder /app/.env /

# Set working directory
# WORKDIR /app

# Set the command to run the binary
CMD ["/main"]