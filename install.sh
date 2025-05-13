#!/usr/bin/env bash
set -e

echo "[quick-open] Installing 'qo' command into '$HOME/.local/bin'"

# check for go
if ! [ -x "$(command -v go)" ]; then
  echo 'Error: go is not installed.' >&2
  exit 1
fi

if [ ! -d "$HOME/.local/bin" ]; then
  echo "Error: $HOME/.local/bin does not exist. Please create this directory before running the install script" >&2
  exit 1
fi

go mod download
go build -o $HOME/.local/bin/qo

echo "[quick-open] Installation complete! Run 'qo' to get started."
echo "[quick-open] example: '\$ qo -h'"


