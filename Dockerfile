# Use the official Go image as a base
FROM golang:1.23.4

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files (go.mod and go.sum)
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Expose the port your app listens on
EXPOSE 5000

# Define the command to run when the container starts
CMD ["./main"]