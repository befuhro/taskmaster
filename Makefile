PROJECT_NAME := "taskmaster"
PKG := "./..."
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
TEST_LIST := $(shell go list ./tests/... | grep -v /vendor/)
GO_FILES := $(shell find . -name "*.go" | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep build clean test lint

all: build

lint:
	@golint -set_exit_status ${PKG_LIST}

test:
	@go test -short ${TEST_LIST}

race: dep
	@go test -race -short ${TEST_LIST}

msan: dep
	@go test -msan -short ${PKG_LIST}

dep:
	@go mod download

build: dep
	@go build -v $(PKG)

clean:
	@rm -f $(PROJECT_NAME)

install: dep
	@go install

uninstall:
	@rm $(GOPATH)/bin/$(PROJECT_NAME)
