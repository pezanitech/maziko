#!/bin/bash
set -euo pipefail

NODE_VERSION="v20.11.1"
INSTALL_DIR="./build/nodejs"

# Detect OS
OS="$(uname -s)"
case "$OS" in
    Linux) PLATFORM="linux" ;;
    Darwin) PLATFORM="darwin" ;;
    MINGW*|MSYS*|CYGWIN*) PLATFORM="win" ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

# Detect ARCH
ARCH="$(uname -m)"
case "$ARCH" in
    x86_64) ARCH="x64" ;;
    aarch64) ARCH="arm64" ;;
    armv7l) ARCH="armv7l" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Final target name
NODE_TARGET="${PLATFORM}-${ARCH}"

# File extension based on platform
EXT="tar.xz"
[ "$PLATFORM" = "win" ] && EXT="zip"

# Build URL
BASE_URL="https://nodejs.org/dist/${NODE_VERSION}"
FILENAME="node-${NODE_VERSION}-${NODE_TARGET}.${EXT}"
URL="${BASE_URL}/${FILENAME}"

echo "Downloading Node.js ${NODE_VERSION} for ${NODE_TARGET}..."

# Prepare install dir
rm -rf "$INSTALL_DIR"
mkdir -p "$INSTALL_DIR"

# Download and extract
if [ "$EXT" = "tar.xz" ]; then
    curl -fsSL "$URL" | tar -xJ --strip-components=1 -C "$INSTALL_DIR"
else
    TMPDIR=$(mktemp -d)
    curl -fsSL -o "$TMPDIR/$FILENAME" "$URL"
    unzip -q "$TMPDIR/$FILENAME" -d "$TMPDIR"
    mv "$TMPDIR/node-${NODE_VERSION}-${NODE_TARGET}"/* "$INSTALL_DIR"
    rm -rf "$TMPDIR"
fi

echo "Done. Node.js is available at: $INSTALL_DIR/bin/node"
echo "Add to PATH with: export PATH=\"$INSTALL_DIR/bin:\$PATH\""
