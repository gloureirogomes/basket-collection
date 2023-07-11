APP_NAME=basket-collection

## setup: run mod download and mod tidy
setup:
	GO111MODULE=on go mod download
	go mod tidy
	go mod verify

## lint: run all linters
lint:
	golangci-lint run ./...	

## build: build to create executable file
build:
	go build -o ${APP_NAME} .

## run-api: build and run api
run-api: build
	./${APP_NAME} api