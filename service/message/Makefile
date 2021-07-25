DATETIME = `date '+%Y%m%d_%H%M'`

GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=:. --go_out=:. proto/*/*.proto

.PHONY: build
build:
	go build -o message *.go

.PHONY: test
test:
	go test -v ./... -cover

docker: deploy_build
	./push_image.sh message

deploy_build:
	GOOS=linux GOARCH=amd64 go build -o message main.go

pack_code: clean
	tar czvf message-${DATETIME}.tar.gz *

.PHONY:
clean:
	@rm -f message-*.tar.gz message


