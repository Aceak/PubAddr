#!/usr/bin/env bash

set -e

VERSION="$1"

# 自动定位项目根目录（脚本所在目录的上一级）
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

CHANGELOG_FILE="$REPO_ROOT/CHANGELOG.md"

if [[ -z "$VERSION" ]]; then
    echo "Usage: $0 <version>"
    exit 1
fi

if [[ ! -f "$CHANGELOG_FILE" ]]; then
    echo "[!] CHANGELOG.md not found at $CHANGELOG_FILE"
    exit 1
fi

awk -v version="## $VERSION " '
    index($0, version) == 1 {found=1; next}
    found && /^## / {exit}
    found
' "$CHANGELOG_FILE"
