# Use the official golang image as a base image
FROM golang:alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o ./bin/main ./cmd/http/main

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin/main .

# Ensure the binary is executable
RUN chmod +x ./main

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
