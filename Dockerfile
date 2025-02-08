# Stage 1: Build Golang Application
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main

# Stage 2: Create Final Image
FROM alpine:latest
WORKDIR /root/

# Copy the compiled Go binary
COPY --from=builder /app/main .

# Ensure the database directory exists in the container
RUN mkdir -p /root/database

# Copy the SQLite database from the project directory
COPY database/notifications.db /root/database/notifications.db

# Expose the necessary port
EXPOSE 5002

# Start the application
CMD ["./main"]
