# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Goulash is a Go utility library that provides functional programming utilities using Go generics. The library follows a flat package structure where each utility function is implemented in its own file with corresponding tests.

## Common Commands

### Build
```bash
go build ./...
```

### Test
```bash
# Run all tests with coverage
go test -v -coverprofile=coverage.out ./...

# Using Task runner
task test
```

### Task Runner
The project uses [Task](https://taskfile.dev) for build automation:
```bash
# Run default build
task

# Run tests with coverage
task test

# Show coverage report (requires dokimi)
task show-cover
```

## Code Architecture

### Package Structure
- **Root package**: All utility functions are in the main `goulash` package
- **utils/**: Contains generic utility types and reflection helpers
- **internal/generic/**: Internal generic utilities (currently empty directory)

### Function Organization
Each utility function follows this pattern:
- `function.go` - Implementation
- `function_test.go` - Comprehensive tests using testify

### Dependencies
- `golang.org/x/exp/constraints` - For generic type constraints
- `github.com/stretchr/testify` - Testing framework
- Uses Go 1.25 features

### Key Design Patterns
1. **Generic Functions**: All functions use Go generics with appropriate constraints
2. **Functional Style**: Functions are pure and composable (e.g., `__.Reduce(__.Map(__.Filter(...)))`)
3. **Type Safety**: Heavy use of `constraints.Ordered` and reflection for type checking
4. **Flat Structure**: Each function is a top-level export in the main package

### Import Convention
The library is designed to be imported with an underscore alias:
```go
import __ "github.com/farbodsalimi/goulash"
```

### Testing Approach
- Uses testify for assertions and test structure
- Each function has comprehensive test coverage including edge cases
- Tests follow the pattern `TestFunctionName` with subtests for different scenarios

## Available Functions

The library provides comprehensive functional programming utilities including:
- **Collection operations**: Filter, Map, Reduce, FlatMap, GroupBy, Partition
- **Set operations**: Union, Intersection, Difference, Unique
- **Aggregation**: All, Any, None, Min, Max, MinMax, Sum, Product, Count
- **Transformation**: Chunk, Compact, Concat, Sort, Zip, Reverse, Flatten, Transpose
- **Slicing**: Take, Drop, TakeWhile, DropWhile, Sliding
- **Search & Find**: Contains, Find, FindIndex, IndexOf, LastIndexOf  
- **Generation**: Range, Repeat, RandomSample, Permutations, Combinations
- **Utility**: Join, Keys, Values, Ternary, ForEach, Fold, Scan