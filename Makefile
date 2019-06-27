SHORT_NAME ?= relayr-app

include versioning.mk

PKGS := $(shell go list ./...)

# # It's necessary to set this because some environments don't link sh -> bash.
SHELL := /bin/bash
# Common flags passed into Go's linker.
GOTEST := go test -v

LDFLAGS := "-s -w \
-X github.com/cloud104/tks-controller/pkg/version.version=${VERSION} \
-X github.com/cloud104/tks-controller/pkg/version.gitCommit=${GITCOMMIT} \
-X github.com/cloud104/tks-controller/pkg/version.buildDate=${DATE}"

BINARY_DEST_DIR := rootfs/usr/local/bin

GOOS ?= linux
GOARCH ?= amd64

test: lint
	$(GOTEST) $(PKGS)

$(GOLINT):
	go get -u github.com/golang/lint/golint

.PHONY: lint
lint: $(GOLINT)
	go list ./... | xargs golint

build: clean
	env GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 go build -ldflags ${LDFLAGS} -o ${BINARY_DEST_DIR}/relayr-app cmd/relayr/main.go

docker-release:
	upx rootfs/usr/local/bin/* || true
	docker build -t ${REGISTRY}/${IMAGE_PREFIX}/${SHORT_NAME}:${VERSION} ./rootfs
	docker push ${REGISTRY}/${IMAGE_PREFIX}/${SHORT_NAME}:${VERSION}
	make clean

clean:
	rm -f ${BINARY_DEST_DIR}/*

.PHONY: test build clean
