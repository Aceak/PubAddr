#!/usr/bin/env bash

set -e

APP_NAME="pubaddr"
BUILD_DIR="build"

rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR"

build() {
    local GOOS=$1
    local GOARCH=$2
    local SUFFIX=$3
    local OUT="$BUILD_DIR/${APP_NAME}-${GOOS}-${GOARCH}${SUFFIX}"

    echo "[+] Building ${GOOS}/${GOARCH} ..."
    GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -o "$OUT" ./cmd/main.go
}

### Linux
build linux amd64 ""
build linux arm ""           # armv7
build linux arm64 ""
build linux riscv64 ""

### Windows
build windows amd64 ".exe"
build windows arm64 ".exe"

### macOS（不支持 riscv64）
build darwin amd64 ""
build darwin arm64 ""

echo "[OK] All binaries are in $BUILD_DIR/"
