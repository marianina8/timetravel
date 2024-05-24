# Use an official Golang runtime as a parent image
FROM golang:1.19

# Copy go.mod and go.sum files
COPY . /app

# Set the current working directory inside the container
WORKDIR /app

# Build the Go app
RUN go build -o timetravel main.go

# Command to run the executable
ENTRYPOINT ["./timetravel"]
