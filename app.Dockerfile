FROM golang:latest

WORKDIR /go/src/github.com/lakshanwd/muve-go/go-crud

COPY go-crud/auth ./auth
COPY go-crud/dao ./dao
COPY go-crud/db ./db
COPY go-crud/handler ./handler
COPY go-crud/repo ./repo
COPY go-crud/server.go ./server.go

RUN go get -d -v ./... && go install -v ./...

COPY go-crud/config.docker.json ./config.json

CMD ["go-crud"]

EXPOSE 8080