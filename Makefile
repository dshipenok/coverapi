.PHONY: build run test check

BINARY_NAME:=calcapi

build:
	go build -o ../../bin/$(BINARY_NAME) ./cmd/calcapi

run:
	 ../../bin/$(BINARY_NAME)

test:
	go test -cover ./... -covermode=count  -coverprofile=coverage.out

check:
	curl -X GET "localhost:8080/divide?a=1&b=10"
