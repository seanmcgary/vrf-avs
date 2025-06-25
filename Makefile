# -----------------------------------------------------------------------------
# This Makefile is used for building your AVS application.
#
# It contains basic targets for building the application, installing dependencies,
# and building a Docker container.
#
# Modify each target as needed to suit your application's requirements.
# -----------------------------------------------------------------------------

GO = $(shell which go)
OUT = ./bin

# Detect platform and set GMP paths accordingly
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
    # macOS with Homebrew
    GMP_LIB_PATH := /opt/homebrew/lib
    GMP_INCLUDE_PATH := /opt/homebrew/include
else
    # Linux (Docker/CI)
    GMP_LIB_PATH := /usr/lib/x86_64-linux-gnu:/usr/lib
    GMP_INCLUDE_PATH := /usr/include
endif

build: deps build-vdf
	@mkdir -p $(OUT) || true
	@echo "Building binaries..."
	./scripts/compile-bindings.sh
	go build -o $(OUT)/performer ./cmd/performer/main.go
	go build -o $(OUT)/vrf-client ./cmd/client/main.go

deps:
	GOPRIVATE=github.com/Layr-Labs/* go mod tidy

build/container:
	./.hourglass/scripts/buildContainer.sh

test:
	@export LIBRARY_PATH="$(GMP_LIB_PATH):$$LIBRARY_PATH" && \
	export CPATH="$(GMP_INCLUDE_PATH):$$CPATH" && \
	go test ./... -v -p 1

build-vdf:
	@echo "Building VDF shared library..."
	@export LIBRARY_PATH="$(GMP_LIB_PATH):$$LIBRARY_PATH" && \
	export CPATH="$(GMP_INCLUDE_PATH):$$CPATH" && \
	cargo build --release --manifest-path pkg/vdf-ffi/Cargo.toml

test-vdf:
	@export LIBRARY_PATH="$(GMP_LIB_PATH):$$LIBRARY_PATH" && \
	export CPATH="$(GMP_INCLUDE_PATH):$$CPATH" && \
	go test -v ./pkg/vdf

test-contracts:
	cd .devkit/contracts && forge test
