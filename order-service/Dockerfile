# Stage 1: Build the Go binary
FROM golang:1.18-alpine AS builder

RUN mkdir /app
COPY . /app
WORKDIR /app

# Build the binary
RUN go build -o orderapp main.go
RUN chmod +x /app/orderapp

# Stage 2: Use lightweight image for runtime
FROM alpine:latest

RUN mkdir /app
COPY --from=builder /app/orderapp /app

# Start the binary when the container runs
CMD ["/app/orderapp"]
