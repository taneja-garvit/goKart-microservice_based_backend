# Step 1: Use official Go image as a builder
FROM golang:1.18-alpine AS builder

# Step 2: Create app directory inside container
RUN mkdir /app

# Step 3: Copy everything into container’s app directory
COPY . /app

# Step 4: Set working directory
WORKDIR /app

# Step 5: Download dependencies (optional in simple projects)
# RUN go mod download

# Step 6: Build the Go app
RUN go build -o userapp main.go

# Step 7: Give execute permission
RUN chmod +x /app/userapp

# Step 8: Use a minimal image for final build
FROM alpine:latest

# Step 9: Create app directory again
RUN mkdir /app

# Step 10: Copy compiled binary from builder stage
COPY --from=builder /app/userapp /app

# Step 11: Set the binary as the container entrypoint
CMD ["/app/userapp"]
