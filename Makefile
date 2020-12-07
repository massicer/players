lint:
	golangci-lint run

test:
	go test ./... -test.v -coverprofile cp.out

start:
	go run ./cmd/main.go

build:
	go build ./cmd/main.go