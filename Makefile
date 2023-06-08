GOPATH:=$(shell go env GOPATH)

.PHONY: run
run:
	@go run challenge/cmd/server

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build ./cmd/...

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@DOCKER_BUILDKIT=1 docker build -t challenge:latest -f ./Dockerfile .
