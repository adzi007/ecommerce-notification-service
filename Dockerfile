# Stage 1: Build the Go application
FROM buildpack-deps:bookworm AS builder

WORKDIR /app
COPY . .

# Enable CGO for SQLite and install necessary packages
RUN apt-get update && apt-get install -y gcc libc-dev
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Create final minimal image
FROM debian:bookworm-slim

WORKDIR /root/
COPY --from=builder /app/main .
COPY .env .

# Ensure SQLite database is copied
RUN mkdir -p /root/database
COPY database/notifications.db /root/database/notifications.db

EXPOSE 5002
CMD ["./main"]
