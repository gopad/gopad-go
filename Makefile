OPENAPI_VERSION ?= 1.0.0-alpha1
OPENAPI_URL ?= https://dl.gopad.eu/openapi/$(OPENAPI_VERSION).yml

SHELL := bash
NAME := gopad-go
IMPORT := github.com/gopad/$(NAME)

GOBUILD ?= CGO_ENABLED=0 go build
PACKAGES ?= $(shell go list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f -not -path */.devenv/* -not -path */.direnv/*)
GENERATE ?= $(PACKAGES)
TAGS ?= netgo

LDFLAGS += -s -w -extldflags "-static"

.PHONY: all
all: build

.PHONY: sync
sync:
	go mod download

.PHONY: clean
clean:
	go clean -i ./...

.PHONY: fmt
fmt:
	gofmt -s -w $(SOURCES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: golangci
golangci:
	go tool github.com/golangci/golangci-lint/cmd/golangci-lint run ./...

.PHONY: lint
lint:
	for PKG in $(PACKAGES); do go tool github.com/mgechev/revive -config revive.toml -set_exit_status $$PKG || exit 1; done;

.PHONY: generate
generate:
	go generate $(GENERATE)

.PHONY: openapi
openapi:
	go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -templates hack/templates/ -generate types,client -package gopad -o gopad/gen.go $(OPENAPI_URL)

.PHONY: test
test:
	go test -coverprofile coverage.out $(PACKAGES)

.PHONY: build
build: $(SOURCES)
	go build -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o /dev/null ./...
