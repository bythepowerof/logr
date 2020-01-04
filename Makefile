
build:
	go mod download
	go build ./...

test: build
	go test ./... -coverprofile cover.out

.PHONY: build test
