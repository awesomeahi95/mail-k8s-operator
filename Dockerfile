# Use the official Golang image as the base image
FROM golang:1.19 as builder

# Set the working directory inside the container
WORKDIR /workspace

# Copy go.mod and go.sum to the workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code to the workspace
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

# Use a minimal base image for the final image
FROM alpine:3.14
RUN apk add --no-cache ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /workspace/manager .

# Command to run the binary
ENTRYPOINT ["/root/manager"]
