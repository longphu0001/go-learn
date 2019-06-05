#!/bin/bash

source "sh/project.sh"

export GOROOT="$GOROOT"
export GOPATH="$GOPACKAGE"
export GOBIN="$GOPACKAGE/bin"
export PATH="$GOROOT/bin:$PATH:$GOPACKAGE/bin"

cd "$GOPACKAGE_SRC/$PROJECT_NAME"

echo "========================================"
go env
echo "========================================"
