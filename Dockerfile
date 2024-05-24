# Use an official Golang runtime as a parent image
FROM golang:1.19

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o timetravel main.go

# Command to run the executable
ENTRYPOINT ["./timetravel"]
