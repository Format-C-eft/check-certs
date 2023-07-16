GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.20","$(shell printf "$(GO_VERSION_SHORT)\n1.20" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.20. Found: $(GO_VERSION_SHORT))
endif

# определение текущий ос
ifndef HOST_OS
HOST_OS:=$(shell go env GOHOSTOS)
endif
# определение текущий архитектуры
ifndef HOST_ARCH
HOST_ARCH:=$(shell go env GOHOSTARCH)
endif

ifndef BIN_DIR
BIN_DIR=./bin
endif

##################### PROJECT RELATED VARIABLES #####################
GIT_HASH:=$(shell git log --format="%h" -n 1 2> /dev/null)
GIT_BRANCH:=$(shell git branch 2> /dev/null | grep '*' | cut -f2 -d' ')
BUILD_TS:=$(shell date +%FT%T%z)

LDFLAGS = -X 'gitlab.com/fake_services/kraken-go/internal/config.branch=$(GIT_BRANCH)'\
          -X 'gitlab.com/fake_services/kraken-go/internal/config.commitHash=$(GIT_HASH)'\
          -X 'gitlab.com/fake_services/kraken-go/internal/config.timeBuild=$(BUILD_TS)'


BUILD_ENV_PARAMS:=CGO_ENABLED=0

.PHONY: build
build:
	$(info Building...)
	
	$(BUILD_ENV_PARAMS) $(GOX_BIN) -output="$(BIN_DIR)/{{.Dir}}" -osarch="$(HOST_OS)/$(HOST_ARCH)" -ldflags "$(LDFLAGS)" ./cmd/check-certs

