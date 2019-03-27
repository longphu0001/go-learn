#!/bin/bash

initArch() {
    ARCH=$(uname -m)
    if [ -n "$DEP_ARCH" ]; then
        echo "Using DEP_ARCH"
        ARCH="$DEP_ARCH"
    fi
    case $ARCH in
        amd64) ARCH="amd64";;
        x86_64) ARCH="amd64";;
        i386) ARCH="386";;
        ppc64) ARCH="ppc64";;
        ppc64le) ARCH="ppc64le";;
        s390x) ARCH="s390x";;
        armv6*) ARCH="arm";;
        armv7*) ARCH="arm";;
        aarch64) ARCH="arm64";;
        *) echo "Architecture ${ARCH} is not supported by this installation script"; exit 1;;
    esac
    echo "ARCH = $ARCH"
}

initOS() {
    OS=$(uname | tr '[:upper:]' '[:lower:]')
    OS_CYGWIN=0
    if [ -n "$DEP_OS" ]; then
        echo "Using DEP_OS"
        OS="$DEP_OS"
    fi
    case "$OS" in
        darwin) OS='darwin';;
        linux) OS='linux';;
        freebsd) OS='freebsd';;
        mingw*) OS='windows';;
        msys*) OS='windows';;
        cygwin*)
            OS='windows'
            OS_CYGWIN=1
        ;;
        *) echo "OS ${OS} is not supported by this installation script"; exit 1;;
    esac
    echo "OS = $OS"
}

# identify platform based on uname output
initArch
initOS

GO_VERSION="1.12.1"
FILE_TAR="go$GO_VERSION.$OS-$ARCH.tar.gz"
LINK_FILE="https://dl.google.com/go/$FILE_TAR"
GOROOT="$HOME/go"
GOPACKAGE="$HOME/go_package"
GOPACKAGE_SRC="$GOPACKAGE/src"
PROJECT_NAME="$(basename "$PWD")"

if [ ! -d "$GOROOT" ]; then
    if [ -f "$FILE_TAR" ]; then
        rm "$FILE_TAR"
    fi
    
    wget "$LINK_FILE"
    tar zvxf "$FILE_TAR" -C "$HOME/"
    rm "$FILE_TAR"
fi

if [ ! -d "$GOPACKAGE_SRC" ]; then
    mkdir -p "$GOPACKAGE_SRC"
fi

if [ ! -e "$GOPACKAGE_SRC/$PROJECT_NAME" ]; then
    ln -s $PWD "$GOPACKAGE_SRC/$PROJECT_NAME"
fi

export GOROOT="$GOROOT"
export GOPATH="$GOPACKAGE"
export GOBIN="$GOPACKAGE/bin"
export PATH="$GOROOT/bin:$PATH:$GOPACKAGE/bin"

cd "$GOPACKAGE_SRC/$PROJECT_NAME"

echo "========================================"
go env
echo "========================================"


