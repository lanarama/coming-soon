.PHONY: build

KO_DOCKER_REPO ?= ghcr.io/lanarama/coming-soon
TAG ?= $(shell git rev-parse --short HEAD)
PUSH ?= true

build:
	 KO_DOCKER_REPO=$(KO_DOCKER_REPO) ko build  -t $(TAG) --platform=linux/amd64,linux/arm64 --push=$(PUSH) .