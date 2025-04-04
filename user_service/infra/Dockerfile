# Stage 1: This is the build stage
FROM golang:1.23.4 as builder
# Set the working directory inside the container
WORKDIR /app
# Copy go.mod and go.sum files and install dependencies
COPY ../go.mod ./ 
COPY ../go.sum ./
RUN go mod download
# Copy the source code
COPY ../. .
RUN go build -o user_service_api .

# Stage 2: Runtime stage
FROM gcr.io/distroless/base-debian11
# Set the working directory inside the container
WORKDIR /app
# Copy the binary from the builder stage
COPY --from=builder /app/user_service_api .
# Command to run the application
CMD ["./user_service_api"]
