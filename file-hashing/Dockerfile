# Use the official Golang image as the base image
FROM golang:1.17-alpine3.14 AS build
# Set the working directory to /app
WORKDIR /app
# Copy the Go module files
COPY go.mod go.sum ./
# Download and cache Go modules
RUN go mod download
# Copy the rest of the application code
COPY . .
# Build the application
RUN go build -o file-hashing
# Expose port 8080 for the service
EXPOSE 8080
# Run the service when the container starts
CMD ["/app/file-hashing"]
