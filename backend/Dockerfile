# Use the official Golang image
FROM golang:1.23.1 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY backend/ .

# Build the Go app
RUN go build -o main .

# Use a smaller base image to run the app
FROM alpine:latest

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=build /app/main .

# Expose port 7777 to the outside world
EXPOSE 7777

# Command to run the executable
CMD ["./main"]
