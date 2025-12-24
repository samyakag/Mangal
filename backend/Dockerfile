
# Stage 1: Build the Go application
FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Stage 2: Run the application
FROM alpine:latest

# Install ca-certificates for MongoDB Atlas TLS connections
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8001

CMD ["./main"]

