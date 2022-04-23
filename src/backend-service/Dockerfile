FROM golang:latest

RUN mkdir -p /go/src

COPY src/backend-service /go/src 

ENTRYPOINT cd /go/src && go run main.go