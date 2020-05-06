MODULE = $(shell go list -m)
SHELL := /bin/bash
LINT_TOOL=$(shell go env GOPATH)/bin/golangci-lint
GO_PKGS=$(foreach pkg, $(shell go list ./...), $(if $(findstring /vendor/, $(pkg)), , $(pkg)))
GO_FILES=$(shell find . -type f -name '*.go' -not -path './vendor/*')

ENV := local
ifdef $$ENV
ENV := $$ENV
endif

export PROJECT = github.com/qreasio/go-starter-kit

build:
	env GOOS=linux GOARCH=amd64 go build -o bin/server $(PROJECT)/cmd
	chmod +x bin/server

build-mac:
	env GOOS=darwin GOARCH=amd64 go build -o bin/server $(PROJECT)/cmd
	chmod +x bin/server

run:
	go run .c/cmd/server/main.go

migrate-up:
	go run ./cmd/admin/main.go migrate config/${ENV}.yaml up

migrate-down:
	go run ./cmd/admin/main.go migrate config/${ENV}.yaml down

seed:
	go run ./cmd/admin/main.go seed config/${ENV}.yaml test/testdata/seed.sql

test:
	go test ./... -count=1

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -t -d -v ./...

deps-cleancache:
	go clean -modcache

fmt:
	@go fmt $(GO_PKGS)
	@goimports -w -l $(GO_FILES)

$(LINT_TOOL):
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.26.0

qc: $(LINT_TOOL)
	$(LINT_TOOL) run --config=.golangci.yaml ./...