# Step 1: Build the application in a Golang container
# Step 1: Build the application in a Golang container
FROM golang:1.23.2-alpine AS builder

# Install build dependencies for CGO
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container
COPY . .

ENV CGO_ENABLED=1

# Build the application
RUN go build -o /LateMateBot ./main.go

# Step 2: Create a minimal container with only the built binary
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /LateMateBot /LateMateBot


# Start the application
CMD ["/LateMateBot"]
