# Base stage for common setup
FROM golang:1.22 as base

WORKDIR /app

# Copy only the necessary files for dependency resolution
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the container
COPY . .

# Development stage
FROM base AS development

# Install air for live reloading
RUN go install github.com/cosmtrek/air@v1.27.3

# Set the working directory
WORKDIR /app

# Command to run during development
CMD ["air"]

# Production stage
FROM base AS production

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# Final stage for production
FROM alpine:3.14 AS final

# Define an argument for the application user
ARG APP_USER

# Install CA certificates
RUN apk --no-cache add ca-certificates && \
    # Create non-root user for security
    adduser -D "${APP_USER}"

# Set the working directory
WORKDIR /app

# Copy the binary from the production stage
COPY --from=production /app/main .

# Switch to a non-root user for security
USER "${APP_USER}"

# Command to run the application
CMD ["./main"]
