# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /bin/bash

.PHONY: all
all: fmt vet install

.PHONY: install
install:
	go install -mod=vendor  ./...

.PHONY: fmt
fmt:
	goimports -w . pkg
	gofmt -s -w . pkg

.PHONY: vet
vet:
	go vet . ./pkg/...

.PHONY: test
test:
	go test -mod=vendor -v ./...

.PHONE: dep
dep:
	go mod download
	go mod vendor
	go mod tidy
