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
echo "[quick-open] Success!"

if [ ! -d "$HOME/.config/quick-open" ]; then
    echo "[quick-open] Creating config directory '$HOME/.config/quick-open'"
    mkdir -p "$HOME/.config/quick-open"
fi

read -p "Do you want to install completions? [y/N] " response
if [[ "$response" =~ ^[Yy]$ ]]; then
    if [[ "$SHELL" == *"bash"* ]]; then
        sudo cp ./autocomplete/bash_autocomplete /etc/bash_completion.d/qo
    elif [[ "$SHELL" == *"zsh"* ]]; then
        if [ ! -f ~/.config/quick-open/zsh_autocomplete ]; then
            cp ./autocomplete/zsh_autocomplete ~/.config/quick-open/zsh_autocomplete
            if [ -f ~/.zshrc ]; then
                echo "export PROG=qo; source ~/.config/quick-open/zsh_autocomplete" >> ~/.zshrc
            else
                echo "Error: ~/.zshrc not found. Please add 'export PROG=qo; source ~/.config/quick-open/zsh_autocomplete' to your zsh configuration manually."
            fi
        fi
    else
        echo "Unsupported shell: $SHELL"
        exit 1
    fi
    echo "[quick-open] Shell completion installed successfully"
fi

echo "[quick-open] Installation complete! Run 'qo' to get started."
echo "[quick-open] example: '\$ qo -h'"
