# Start from the official Golang image for building the app
FROM golang:1.25-alpine AS builder
# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new minimal image for running
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file from the build context to the runtime image
COPY --from=builder /app/.env .

# Copy any other necessary files (e.g., migrations, configs) if needed
# COPY --from=builder /app/migrations ./migrations

# Expose the port your app runs on (change if needed)
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
