# Use official Golang image as a base
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o /app/main ./cmd

# Verify that the binary was created
RUN ls -l /app

# Expose application port
EXPOSE 8080

# Start the application
CMD ["/app/main"]

