FROM golang:alpine

RUN mkdir /app
COPY go.mod go.sum /app/

ADD client/client.go /app

WORKDIR /app

RUN go build -o client .
