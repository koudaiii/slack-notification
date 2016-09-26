NAME := slack-notification
ORGANIZER := wantedly
VERSION := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)

LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

DIST_DIRS := find * -type d -exec

.DEFAULT_GOAL := bin/$(NAME)

.PHONY: ci-build
ci-build: deps
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: deploy
deploy:
	@./script/deploy

.PHONY: build
build: deps
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: install
install: deps
	go install $(LDFLAGS)

.PHONY: deps
deps: glide
	glide install

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: test
test: deps
	go test -v `glide novendor`
