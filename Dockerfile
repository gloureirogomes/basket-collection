FROM golang:1.21.1-alpine as setup
RUN apk add git make

WORKDIR /app

EXPOSE 8888

COPY go.mod .
COPY go.sum .
COPY Makefile .
RUN make setup

COPY . /app
RUN make build

ENTRYPOINT ["/app/basket-collection"]