#!/bin/sh

set -e

GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo "Installing Maziko CLI..."

if ! command -v go >/dev/null 2>&1; then
  echo "Error: Go is not installed. Install Go first."
  exit 1
fi

go install github.com/pezanitech/maziko/cli/maziko@v0.1.6

# Determine GOBIN
if [ -z "$GOBIN" ]; then
  GOBIN="$(go env GOPATH)/bin"
  echo "GOBIN is not set. Defaulting to: $GOBIN"
fi

# Export to current session
if ! echo "$PATH" | grep -q "$GOBIN"; then
  export PATH="$PATH:$GOBIN"
  echo "Temporarily added $GOBIN to PATH for this session."
fi

echo -e "${GREEN}Maziko CLI successfully installed!${NC}"
echo -e "${GREEN}Try running: maziko --help${NC}"

echo
echo "If maziko command cannot be found, add the following to your shell config (e.g., ~/.bashrc, ~/.zshrc):"
echo "  export GOBIN=\"$GOBIN\""
echo "  export PATH=\"\$PATH:\$GOBIN\""
