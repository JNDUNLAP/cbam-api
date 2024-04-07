# Start from the official Golang base image for the build stage
FROM golang:1.22 as base

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Development stage builds the application with live reloading using Air
FROM base AS development

# Install Air for live reloading
RUN go install github.com/cosmtrek/air@v1.27.3

# Expose port for live reloading
EXPOSE 8080

# Start Air for live reloading
CMD ["air"]

# Production stage builds the Go app as a static binary
FROM base AS production

# Build the Go app as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# Final stage uses the small Alpine Linux image
FROM alpine:3.14 AS final

# Accept the user to be created as an argument
ARG APP_USER

# Install ca-certificates in case the app makes outgoing HTTPS requests
RUN apk --no-cache add ca-certificates && \
    adduser -D "${APP_USER}"

# Set the current working directory inside the container
WORKDIR /app

# Copy the static binary from the production stage
COPY --from=production /app/main .

# Run the app as the specified user for security purposes
USER "${APP_USER}"

# Command to run the executable
CMD ["./main"]
