# Use the official Golang image as the base image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of your code to the working directory
COPY . .

# Set the working directory for the build
WORKDIR /app/cmd/app

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]
