.PHONY: build
TARGET_DIR ?= build/app
build:
	[[ -d ${TARGET_DIR} ]] || mkdir -p ${TARGET_DIR}
	CGO_ENABLED=0 go build -o ${TARGET_DIR} ./cmd/api

start:
	go run ./cmd/api
