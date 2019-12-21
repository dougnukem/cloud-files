GO_OS:=$(shell go env GOOS)
GO_ARCH:=$(shell go env GOARCH)
GO_BIN_DIR:=$(CURDIR)/bin/$(GO_OS)-$(GO_ARCH)
GO_TOOLS_BIN_DIR:=$(CURDIR)/bin/tools/$(GO_OS)-$(GO_ARCH)
GO_ENV=GOFLAGS="-mod=vendor"

all: test build

build:
	mkdir -p $(GO_BIN_DIR)
	$(GO_ENV) go build -o $(GO_BIN_DIR) ./cmd/...

test:
	$(GO_ENV) go test -v ./...

vendor:
	go mod vendor

clean:
	rm -rf bin

