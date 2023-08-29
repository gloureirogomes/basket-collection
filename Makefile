APP_NAME=basket-collection

## setup: run mod download and mod tidy
setup:
	GO111MODULE=on go mod download
	go mod tidy
	go mod verify

## build: build to create executable file
build:
	go build -o ${APP_NAME} .

## run-api: build and run api
run-api: build
	./${APP_NAME}

## test: run all unit tests
test:
	go test -race $(ARGS) ./... -coverprofile=coverage.out -short -count=1

## integration test: run all integration tests
integration-test:
	go test -race  ./... -run Integration -count=1

## kill-process: kill process running on port 8888
kill-process:
	sudo kill -9 $(sudo lsof -t -i:8888)