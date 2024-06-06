FROM golang:1.22.0 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY /server-sync /app

# Build the Go application
RUN go build -o syncit

# Start a new stag
FROM ubuntu:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app /app

# Expose the port the application runs on
EXPOSE 50000

# Command to run the application
CMD ["/app/syncit"]