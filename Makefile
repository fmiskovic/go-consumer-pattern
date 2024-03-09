ARTIFACT_NAME := go-consumer-pattern

build:
	@go build -o bin/${ARTIFACT_NAME} cmd/main.go

run:
	@go run cmd/main.go

test:
	@go test -v $(shell go list ./... | grep -v /test/)

test-cover:
	@go test -coverprofile cover.out -v $(shell go list ./... | grep -v /test/)
	@go tool cover -html=cover.out

gen-mocks:
	@mockery --all --with-expecter --keeptree
