all: fmt vet build install

init: download

download:
	go mod download

run:
	go run main.go

build:
	go build ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test --count=1 ./...

install:
	go install ./...

clean:
	go clean

.PHONY: all init build fmt vet test install
