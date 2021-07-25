
.PHONY: init proto build docker test deploy_build

GOPATH:=$(shell go env GOPATH)

init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3

proto:
	protoc --proto_path=. --micro_out=:. --go_out=:. proto/follower/*.proto

build:
	go build -o follower *.go

test:
	go test -v ./... -cover

docker: deploy_build
	./push_image.sh follower

deploy_build:
	GOOS=linux GOARCH=amd64 go build -o follower main.go
