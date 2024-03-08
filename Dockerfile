# Use a more recent official Go image for updated features and security patches.
FROM golang:1.18-alpine as builder

# Set environment variables for building the Go application.
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create and set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum for dependency management and leverage Docker layer caching.
COPY go.* ./
RUN go mod download

# Copy the rest of the application code.
COPY . .

# Build the binary with the output name cbamapi.
RUN go build -v -o cbamapi

# Start a new stage from scratch to create a small final image.
FROM alpine:latest

# Set the working directory in the container.
WORKDIR /app

# Copy the built binary from the builder stage to the production image.
COPY --from=builder /app/cbamapi /app/

# Create a non-root user 'jndunlap' and switch to it for security reasons.
RUN adduser -D jndunlap
USER jndunlap

# Health check (adjust the path to your health check endpoint if necessary).
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:1323/health || exit 1

# Inform Docker that the container listens on port 1323.
EXPOSE 1323

# Command to run the application.
CMD ["./cbamapi"]
