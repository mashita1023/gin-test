FROM golang:1.16.3

ENV GO111MODULE=on

WORKDIR /go/src/github.com/NUTFes/shift-app

RUN go get -u github.com/cosmtrek/air
# RUN go mod tidy
CMD air
