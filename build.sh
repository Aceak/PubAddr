#!/usr/bin/env bash

set -e

APP_NAME="pubaddr"
BUILD_DIR="build"

echo "[>] Starting cross compilation..."

rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR"

build() {
    local GOOS=$1
    local GOARCH=$2
    local SUFFIX=$3
    local OUT="$BUILD_DIR/${APP_NAME}-${GOOS}-${GOARCH}${SUFFIX}"

    # 捕获错误，不让 go build 的输出直接破坏脚本格式
    if ! GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -o "$OUT" ./cmd/main.go 2>build_error.log; then
        echo "[!] build pubaddr_$GOOS_$GOARCH$SUFFIX failed"
        echo "    $(cat build_error.log)"
        rm -f build_error.log
        exit 1
    fi

    echo "[+] build $GOOS/$GOARCH$SUFFIX success"
}

### Linux
build linux amd64 ""
build linux arm ""
build linux arm64 ""
build linux riscv64 ""

### Windows
build windows amd64 ".exe"
build windows arm64 ".exe"

### macOS
build darwin amd64 ""
build darwin arm64 ""

echo "[=] All targets compiled successfully"
