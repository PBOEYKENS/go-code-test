# Use the official Golang image as a base image
FROM golang:1.20-alpine

# Expose the port on which your backend listens
EXPOSE 8080

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files into the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code into the container
COPY . .

# Build the application
RUN cd cmd && go build -o main && cd ..

# Set the start command to run the binary
CMD ["./cmd/main"]
