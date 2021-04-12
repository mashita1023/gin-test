FROM golang:1.16.3

WORKDIR /go/src/github.com/NUTFes/shift-app

RUN go get -u github.com/cosmtrek/air

CMD air