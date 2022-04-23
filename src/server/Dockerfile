FROM golang:latest

RUN mkdir -p /go/src

COPY . /go/src 

ENTRYPOINT cd /go/src && go run main.go