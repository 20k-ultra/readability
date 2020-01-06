FROM golang:latest

RUN mkdir -p $GOPATH/src/github.com/20k-ultra/readability

WORKDIR $GOPATH/src/github.com/20k-ultra/readability

COPY . .