GIT_TAG ?= $(or ${TRAVIS_COMMIT},${TRAVIS_TAG},latest)
MUTABLE_VERSION ?= latest
VERSION ?= ${GIT_TAG}
GITCOMMIT ?= $(shell git rev-parse HEAD)
DATE ?= $(shell date -u "+%Y-%m-%dT%H:%M:%SZ")

REGISTRY ?= quay.io
IMAGE_PREFIX ?= andrleite

IMAGE := ${REGISTRY}/${IMAGE_PREFIX}/${SHORT_NAME}:${VERSION}
MUTABLE_IMAGE := ${REGISTRY}/${IMAGE_PREFIX}/${SHORT_NAME}:${MUTABLE_VERSION}

info:
	@echo "Build tag:       ${VERSION}"
	@echo "Registry:        ${REGISTRY}"
	@echo "Immutable tag:   ${IMAGE}"
	@echo "Mutable tag:     ${MUTABLE_IMAGE}"