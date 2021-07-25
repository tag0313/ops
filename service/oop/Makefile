
.PHONY: init build test docker

GOPATH:=$(shell go env GOPATH)
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=:. --go_out=:. proto/*/*.proto

.PHONY: build
build:
	go build -o oop *.go

.PHONY: test
test:
	go test -v ./... -cover

docker: deploy_build
	./push_image.sh oop

deploy_build:
	GOOS=linux GOARCH=amd64 go build -o oop main.go
