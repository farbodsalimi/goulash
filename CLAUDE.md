# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Goulash is a Go utility library providing functional programming utilities using Go generics. Each utility function lives in its own file with a corresponding `_test.go` file, all in the root `goulash` package.

## Common Commands

```bash
# Build
go build ./...

# Run all tests with verbose output and coverage
go test -v -coverprofile=coverage.out ./...

# Run a single test function
go test -v -run TestFunctionName ./...

# Task runner (https://taskfile.dev)
task          # build (default)
task test     # tests with coverage
task show-cover  # coverage report (requires dokimi)
```

## Code Architecture

### Key Design Patterns
- **Flat structure**: Every function is a top-level export in the `goulash` package — one function per file (`function.go` + `function_test.go`)
- **Generic functions**: All functions use Go generics with appropriate constraints (`constraints.Ordered`, `constraints.Integer | constraints.Float`, `comparable`, or `any`)
- **Pure & composable**: Functions have no side effects and can be chained (e.g., `__.Reduce(__.Map(__.Filter(...)))`)
- **Truthiness via zero-value comparison**: `All`, `Any`, and `Compact` treat the zero value of `T` as falsy (0 for numbers, "" for strings) using `var zero T` with `comparable`

### Import Convention
The library is designed to be imported with a double-underscore alias:
```go
import __ "github.com/farbodsalimi/goulash"
```

### Dependencies
- `golang.org/x/exp/constraints` — generic type constraints
- `github.com/stretchr/testify` — testing (assertions + subtests)
- Go 1.26

### Testing
- Uses testify for assertions
- Tests follow `TestFunctionName` with table-driven subtests
- Most tests use `t.Parallel()`
