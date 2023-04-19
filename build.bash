#!/usr/bin/env bash

set -e

GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o "${HOME}/Desktop/doraemon-bin/doraemon-linux-x86_64"  .
GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o "${HOME}/Desktop/doraemon-bin/doraemon-linux-aarch64" .
