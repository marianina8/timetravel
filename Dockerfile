# Use an official Golang runtime as a parent image
FROM golang:1.19

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o timetravel main.go

# Command to run the executable
ENTRYPOINT ["./timetravel"]
