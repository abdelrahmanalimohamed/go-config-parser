# Start from official Golang base image
FROM golang:1.20-alpine

# Set working directory inside the container
WORKDIR /app

# Copy Go files
COPY main.go .
COPY config.ini .

# Build the Go app
RUN go build -o myapp main.go

# Run the compiled binary
CMD ["./myapp"]
