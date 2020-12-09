lint:
	golangci-lint run

test-coverage:
	go test ./... -test.v -coverprofile cp.out

test:
	go test ./... -test.v

start:
	go run ./cmd/main.go

build:
	go build ./cmd/main.go