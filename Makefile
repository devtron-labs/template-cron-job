#!make

all: build

TAG?=$(shell bash -c 'git log --pretty=format:'%h' -n 1')
FLAGS=
ENVVAR=
GOOS?=darwin
REGISTRY?=686244538589.dkr.ecr.us-east-2.amazonaws.com
BASEIMAGE?=alpine:3.9
#BUILD_NUMBER=$$(date +'%Y%m%d-%H%M%S')
BUILD_NUMBER := $(shell bash -c 'echo $$(date +'%Y%m%d-%H%M%S')')
ENV_FILE?=scripts/dev-conf/envfile.env
GIT_COMMIT =$(shell sh -c 'git log --pretty=format:'%h' -n 1')
BUILD_TIME= $(shell sh -c 'date -u '+%Y-%m-%dT%H:%M:%SZ'')
SERVER_MODE_FULL= FULL
SERVER_MODE_EA_ONLY=EA_ONLY
include $(ENV_FILE)
export

build: clean wire
	$(ENVVAR) GOOS=$(GOOS) go build -o template-cron-job \
			-ldflags="-X 'github.com/devtron-labs/template-cron-job/util.GitCommit=${GIT_COMMIT}' \
			-X 'github.com/devtron-labs/template-cron-job/util.BuildTime=${BUILD_TIME}' \
			-X 'github.com/devtron-labs/template-cron-job/util.ServerMode=${SERVER_MODE_FULL}'"

wire:
	wire

clean:
	rm -f template-cron-job

run: build
	./template-cron-job

.PHONY: build
docker-build-image:  build
	 docker build -t template-cron-job:$(TAG) .

.PHONY: build, all, wire, clean, run, set-docker-build-env, docker-build-push, template-cron-job,
docker-build-push: docker-build-image
	docker tag template-cron-job:${TAG}  ${REGISTRY}/template-cron-job:${TAG}
	docker push ${REGISTRY}/template-cron-job:${TAG}

#############################################################################

build-all: build
	make --directory ./cmd/external-app build