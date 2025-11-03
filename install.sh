#!/bin/bash

set -e

GITHUB_REPO="turzzzin/keep-it-secret"
INSTALL_DIR="/usr/local/bin"
CMD_NAME="kis"

echo_info() {
    echo "[INFO] $1"
}

echo_error() {
    echo "[ERROR] $1" >&2
    exit 1
}

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case $ARCH in
    x86_64) ARCH="amd64" ;;
    arm64) ARCH="arm64" ;;
    aarch64) ARCH="arm64" ;;
    *) echo_error "Unsupported architecture: $ARCH" ;;
esac

LATEST_RELEASE_URL="https://api.github.com/repos/$GITHUB_REPO/releases/latest"
DOWNLOAD_URL=$(curl -s $LATEST_RELEASE_URL | grep "browser_download_url.*$OS-$ARCH" | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
    echo_error "Could not find a binary for your OS/architecture ($OS-$ARCH)."
fi

TEMP_DIR=$(mktemp -d)
BINARY_PATH="$TEMP_DIR/$CMD_NAME"

echo_info "Downloading $CMD_NAME from $DOWNLOAD_URL..."
curl -L --progress-bar -o "$BINARY_PATH" "$DOWNLOAD_URL"

chmod +x "$BINARY_PATH"

echo_info "Installing $CMD_NAME to $INSTALL_DIR..."
sudo mv "$BINARY_PATH" "$INSTALL_DIR/$CMD_NAME"

rm -rf "$TEMP_DIR"

echo_info "Installation complete! You can now use the '$CMD_NAME' command."
