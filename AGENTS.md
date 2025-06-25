# Agent Guidelines for quick-open

## Build & Run Commands
- Build: `go build -o qo`
- Install: `./install.sh` or `make install`
- Run: `./qo [command] [options]`
- Test: `go test ./...`
- Test single package: `go test ./internal/bookmarks`
- Test single function: `go test -run TestFunctionName ./path/to/package`

## Code Style Guidelines
- Follow standard Go code style with `gofmt`
- Error handling: Wrap errors with context using `fmt.Errorf("context: %w", err)`
- Imports: Group standard library, then external, then internal imports
- Types: Define structs with field tags where needed (`json:"field_name"`)
- Function naming: Use CamelCase (e.g., `OpenBookmark`, not `open_bookmark`)
- Package naming: Use lowercase without underscores
- Keep functions small and focused on a single responsibility
- Document exported functions and types with comments
- Use meaningful variable names that describe their purpose