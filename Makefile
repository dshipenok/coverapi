.PHONY: build run test check show-coverage build-coverage run-coverage

BINARY_NAME:=calcapi

build:
	go build -o ../../bin/$(BINARY_NAME) ./cmd/calcapi

run:
	 ../../bin/$(BINARY_NAME)

test:
	go test -cover ./... -covermode=count  -coverprofile=coverage.out

check:
	curl -X GET "localhost:8080/divide?a=1&b=10"

show-coverage:
	go tool cover -html=integration.out

build-coverage:
	go test -tags=coverage -c -o ../../bin/$(BINARY_NAME).coverage -covermode=count -coverpkg ./... ./cmd/calcapi

run-coverage:
	 ../../bin/$(BINARY_NAME).coverage -test.coverprofile=integration.out