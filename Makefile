include build.mk

export GO111MODULE=on
export GOSUMDB=off

LOCAL_BIN:=$(CURDIR)/bin
GOX_BIN:=$(LOCAL_BIN)/gox

.PHONY: wire
wire: .install-wire
	$(LOCAL_BIN)/wire gen ./internal/bootstrap/...

.PHONY: .install-wire
.install-wire:
	$(BUILD_ENV_PARAMS) GOBIN=$(LOCAL_BIN) go install github.com/google/wire/cmd/wire@latest

.PHONY: test
test:
	go test ./... -count=1 -timeout=60s -v -short