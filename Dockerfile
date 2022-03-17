FROM golang:alpine

RUN apk add git

WORKDIR /kumparan

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./kumparan
