
all: fmt vet test build

build:
	go build -o bin/gg main.go

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./... -v

