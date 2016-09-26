NAME := slack-notification
ORGANIZER := wantedly
VERSION := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)

LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

DIST_DIRS := find * -type d -exec

DOCKER_REPOSITORY := quay.io
DOCKER_IMAGE_NAME := $(DOCKER_REPOSITORY)/$(ORGANIZER)/$(NAME)
DOCKER_IMAGE_TAG ?= latest
DOCKER_IMAGE := $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): deps
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: ci-docker-release
ci-docker-release: docker-build
	@docker login -e="$(DOCKER_EMAIL)" -u="$(DOCKER_USERNAME)" -p="$(DOCKER_PASSWORD)" $(DOCKER_REPOSITORY)
	docker push $(DOCKER_IMAGE)

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf dist/*
	rm -rf vendor/*

.PHONY: deps
deps: glide
	glide install

.PHONY: docker-build
docker-build:
ifeq ($(findstring ELF 64-bit LSB,$(shell file bin/$(NAME) 2> /dev/null)),)
	@echo "bin/$(NAME) is not a binary of Linux 64bit binary."
	@exit 1
endif
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_IMAGE)

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif

.PHONY: install
install: deps
	go install $(LDFLAGS)

.PHONY: test
test: deps
	go test -v `glide novendor`
