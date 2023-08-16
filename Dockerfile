# syntax=docker/dockerfile:1
FROM golang:1.20

# Arguments
ARG SERVICE_PORT=8080

# Set destination for COPY command
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY *.go ./

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-server

# Expose port to the outside world
EXPOSE ${SERVICE_PORT}
CMD [ "/todo-server" ]