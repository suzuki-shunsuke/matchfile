#!/usr/bin/env bash

set -eu
set -o pipefail

GOLANGCI_LINT_VERSION=v1.64.7
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" "$GOLANGCI_LINT_VERSION"
