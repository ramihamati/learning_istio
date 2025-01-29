FROM golang:1.23.2-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the API
RUN go build -o api ./

# Use a minimal base image for production
FROM alpine:latest
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/api .

# Expose port (change based on API requirement)
EXPOSE 3001

# Command to run the API
CMD ["./api"]