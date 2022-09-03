FROM golang:latest

WORKDIR /usr/src/goApp

COPY . ./

RUN apt-get update

RUN go mod download


CMD ["go", "run", "main.go"]