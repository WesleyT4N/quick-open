# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

quick-open (`qo`) is a terminal bookmark manager written in Go. It lets users open browser URLs, apps, and files from the command line using aliases.

## Build & Run

- **Build**: `go build -o qo`
- **Install**: `./install.sh` or `make install`
- **Run**: `./qo [command]` or `./qo <title|alias>` to open a bookmark directly
- **Test all**: `go test ./...`
- **Test single package**: `go test ./internal/bookmarks`
- **Test single function**: `go test -run TestFunctionName ./path/to/package`
- **Format**: `gofmt`

## Architecture

The CLI is built with `urfave/cli/v2`. The default action opens a bookmark by title or alias; the `bookmark` (`bm`) subcommand provides `add`, `list`/`ls`, `remove`/`rm`, and `open`/`o`.

- **`main.go`** — Entry point, CLI app and command registration
- **`cmd/`** — CLI command definitions and handlers (`bookmark.go`), config path constant (`constants.go`)
- **`internal/bookmarks/`** — Domain logic: `Bookmark` struct with `Open()`, and `BookmarkManager` for CRUD + JSON persistence to `~/.config/quick-open/bookmarks.json`
- **`internal/lib/`** — OS abstraction: `GetOpenCommand()` returns the platform-appropriate open command (`open`/`xdg-open`/`explorer.exe`)

## Code Style

- Standard Go style via `gofmt`
- Wrap errors with context: `fmt.Errorf("context: %w", err)`
- Import order: stdlib → external → internal
- Structs use `json:"field_name"` tags for serialization

## Notes

- No test files exist yet
- Bookmark data is stored as JSON at `~/.config/quick-open/bookmarks.json`
- Autocomplete scripts for bash/zsh are in `autocomplete/`
