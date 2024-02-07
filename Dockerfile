FROM golang:1.22 AS builder
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy app source code
COPY . .
# Build app on /app/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# From empty image, as we will run the built binary file
FROM scratch
# Set /root working directory
WORKDIR /root
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
# Command to run the executable
CMD ["/root/main"]