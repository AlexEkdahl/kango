#!/bin/bash

go mod download
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/segmentio/golines@latest

cp scripts/pre-commit .git/hooks/pre-commit
echo "UID=$(id -u)" > .env
echo "GID=$(id -g)" >> .env
mkdir -p data
chmod +x .git/hooks/pre-commit
