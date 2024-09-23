# Build stage
FROM golang:1.21-alpine AS build

# Install necessary packages
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Now copy the rest of the files
COPY . .

# Build the Go app
RUN go build -o server ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the build stage
COPY --from=build /app/server .

# Expose port 8080 (can be adjusted if needed)
EXPOSE 8080

# Start the server
CMD ["./server"]
