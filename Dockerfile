# Use the official Golang image
FROM golang:1.23.2-alpine3.20

# Set the working directory
WORKDIR /backend_app

# Install Air
RUN go install github.com/air-verse/air@latest

# Copy all files
COPY ./ ./

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Run the application with Air for live reloading
CMD ["air"]