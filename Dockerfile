# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

RUN apk update
RUN apk add git

ENV PKG_NAME=project-fiber
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH

COPY . $PKG_PATH
RUN go mod tidy
RUN go build main.go

WORKDIR $PKG_PATH
CMD ["go", "run", "main.go"]