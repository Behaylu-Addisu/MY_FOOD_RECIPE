# Use the official Golang image as a base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Modules files and download dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod tidy

# Copy the backend source code into the container
COPY backend/ ./

# Build the Go application
RUN go build -o /app/main .

# Expose the port the app will run on
EXPOSE 8080

# Command to run the application
CMD ["/app/main"]
