FROM golang:1.21-bullseye AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
# Build the Go application inside the container
RUN make build

FROM golang:1.21-bullseye AS builder

COPY --from=builder /app/bin/ /
# Expose the port that the application will listen on
EXPOSE 8080
VOLUME ["/config", "/nzbs"]

