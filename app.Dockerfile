FROM golang:latest

WORKDIR /go/src/app

COPY go-crud/auth ./auth
COPY go-crud/dao ./dao
COPY go-crud/db ./db
COPY go-crud/handler ./handler
COPY go-crud/repo ./repo
COPY go-crud/server.go ./server.go

RUN go get -d -v ./...

RUN go install -v ./...

COPY go-crud/config.docker.json ./config.json

CMD ["app"]

EXPOSE 8080