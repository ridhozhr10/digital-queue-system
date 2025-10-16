# Stage 1: Build the application
FROM golang:1.25.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY ../go.mod ../go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application for the http-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/http-service ./cmd/http-service

# Stage 2: Create the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/http-service .
COPY --from=builder /app/third_party/swagger-ui /app/third_party/swagger-ui
COPY --from=builder /app/api /app/api

# Expose the port the application runs on
EXPOSE 80

# Command to run the application
CMD ["./http-service"]
