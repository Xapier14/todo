# syntax=docker/dockerfile:1
FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/todo/
COPY . .

# RUN go get -d -v

RUN go build -o /go/bin/server cmd/server/server.go

# Runner
FROM scratch

ARG SERVICE_PORT=8080

COPY --from=builder /go/bin/server /go/bin/server

# Expose port to the outside world
EXPOSE ${SERVICE_PORT}
CMD [ "/go/bin/server" ]
