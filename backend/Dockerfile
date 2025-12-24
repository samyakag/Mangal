
# Stage 1: Build the Go application
FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build with CGO disabled for better compatibility
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Run the application
# Use Debian slim for better TLS/SSL compatibility with MongoDB Atlas
FROM debian:bookworm-slim

# Install ca-certificates for MongoDB Atlas TLS connections
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8001

CMD ["./main"]

