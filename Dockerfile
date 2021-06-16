FROM golang:1.14.9-alpine

RUN mkdir /build
ADD go.mod go.sum main.go /build/
WORKDIR /build
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/gin-gonic/contrib/static
RUN go build
