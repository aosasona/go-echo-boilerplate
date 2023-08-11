.PHONY: build
TARGET_DIR ?= build/app
GENERATE ?= 0
build:
	@if [ "${GENERATE}" = 1 ]; then \
		go generate ./...; \
	fi
	CGO_ENABLED=0 go build -o ${TARGET_DIR} ./cmd/api

start:
	go run ./cmd/api

start-dev:
	air -c .air.toml

# Usage
#
# VERSION=0.x.x make create-docker-release

DOCKER_IMAGE ?= gopi
DOCKER_USER ?= your-username
VERSION ?= latest
create-docker-release:
	@echo "=====> Creating docker release ${VERSION}"
	docker build -t ${DOCKER_USER}/${DOCKER_IMAGE}:${VERSION} .
	docker push ${DOCKER_USER}/${DOCKER_IMAGE}:${VERSION}
	@echo "=====> Created docker release ${VERSION}"
	@if [ "${VERSION}" = "latest" ]; then \
		exit 0; \
	fi
	@echo "=====> Creating docker release latest"
	docker tag ${DOCKER_USER}/${DOCKER_IMAGE}:${VERSION} ${DOCKER_USER}/${DOCKER_IMAGE}:latest
	docker push ${DOCKER_USER}/${DOCKER_IMAGE}:latest
