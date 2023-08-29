FROM golang:1.21.0-alpine3.18 as setup
RUN apk add git make

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY Makefile .
RUN make setup

FROM setup as builder

COPY . /app
RUN make build

CMD ["./basket-collection"]