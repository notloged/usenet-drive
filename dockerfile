FROM golang:1.21-bullseye AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY go.mod go.sum ./
RUN go mod download

# Build the Go application inside the container
COPY . ./
RUN make build

# Expose the port that the application will listen on
FROM golang:1.21-bullseye

COPY --from=builder /app/bin/ /
EXPOSE 8080
VOLUME ["/config", "/nzbs"]

