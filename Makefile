
all: fmt vet test build

build:
	go build -o bin/query query.go
	go build -o bin/log log.go

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./... -v

