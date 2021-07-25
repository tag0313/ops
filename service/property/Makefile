.PHONY: build proto init pack_code clean docker test linux_build
DATETIME = `date '+%Y%m%d_%H%M'`

GOPATH:=$(shell go env GOPATH)
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3

proto_cp: proto
	cp -r proto/* ../contract/proto/
	cp -r proto/* ../../web/proto/

build:
	go build -o property *.go

test:
	go test -v ./... -cover


pack_code: clean
	tar -czvf property-${DATETIME}.tar.gz *

clean:
	@rm -f user_info property-*.tar.gz


linux_build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -a -v -o property

docker: linux_build
	./push_image.sh property
