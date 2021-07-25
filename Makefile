.PHONY: proto build api contract follower message oop property report search userInfo userRegister \
		docker_push docker_build docker_clean clean up down swag

SHELL:=/bin/sh
DATETIME = $(shell date '+%Y%m%d_%H%M%S')
PROTOS = `ls proto`

# Path Related
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))
MKFILE_DIR := $(dir $(MKFILE_PATH))
RELEASE_DIR := ${MKFILE_DIR}bin
DOCKER_TAG := zeroone/opsnft

# Version
RELEASE?=1.0.1
ifndef GIT_COMMIT
  GIT_COMMIT := git-$(shell git rev-parse --short HEAD)
endif

GIT_REPO_INFO=$(shell git config --get remote.origin.url)
ifndef GIT_COMMIT
  GIT_COMMIT := git-$(shell git rev-parse --short HEAD)
endif

proto:
	for p in ${PROTOS}; do \
		echo "compiling $$p";\
		protoc --proto_path=. --micro_out=:. --go_out=:. proto/$$p/*.proto; \
	done

# Build Flags
GO_LD_FLAGS= "-s -w -X ops/pkg/version.RELEASE=${RELEASE} -X ops/pkg/version.COMMIT=${GIT_COMMIT} -X ops/pkg/version.REPO=${GIT_REPO_INFO} -X ops/pkg/version.BUILDTIME=${DATETIME} -X ops/pkg/version.SERVICENAME=$@"
CGO_SWITCH := 0
contract userRegister: CGO_SWITCH := 1

build: api contract follower message oop property report search userInfo userRegister

api contract userRegister follower message oop property report search userInfo:
	cd ${MKFILE_DIR} && \
	CGO_ENABLED=${CGO_SWITCH} go build -v -trimpath -ldflags ${GO_LD_FLAGS} \
	-o ${RELEASE_DIR}/$@ ${MKFILE_DIR}cmd/$@/

docker_build:
	docker build -t ${DOCKER_TAG}:${RELEASE} -f deployment/Dockerfile . --network=host

docker_push:
	docker tag ${DOCKER_TAG}:${RELEASE} registry.cn-hongkong.aliyuncs.com/${DOCKER_TAG}:${RELEASE} && \
	docker push registry.cn-hongkong.aliyuncs.com/${DOCKER_TAG}:${RELEASE}

docker_clean:
	 docker images | grep \<none\> | awk '{print  $$3}' |xargs docker rmi -f #clean <none>

clean:
	@rm -f ${MKFILE_DIR}bin/*

up:
	docker-compose -f deployment/docker-compose-local.yaml up -d

down:
	docker-compose -f deployment/docker-compose-local.yaml down


swag:  #It will too slow now, being patient, it works.
	@rm -f cmd/api/docs/swagger.*
	#swag init --dir cmd/api/ --output cmd/api/docs --parseDepth 2 --parseDependency
	swag init --dir cmd/api/ --output cmd/api/docs --parseDepth 2 --parseDependency
