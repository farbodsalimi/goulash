# Goulash

[![Pipeline](https://github.com/farbodsalimi/goulash/actions/workflows/go.yml/badge.svg)](https://github.com/farbodsalimi/goulash/actions/workflows/go.yml)
[![Release](https://github.com/farbodsalimi/goulash/actions/workflows/release.yml/badge.svg)](https://github.com/farbodsalimi/goulash/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/farbodsalimi/goulash?r)](https://goreportcard.com/report/github.com/farbodsalimi/goulash)
[![Maintainability](https://api.codeclimate.com/v1/badges/41c658c64749a3221cf5/maintainability)](https://codeclimate.com/github/farbodsalimi/goulash/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/41c658c64749a3221cf5/test_coverage)](https://codeclimate.com/github/farbodsalimi/goulash/test_coverage)

Goulash offers a variety of useful utility functions for everyday programming tasks by leveraging Go generics and keeping the functional programming paradigm in mind.

**Table of Contents**

- [Goulash](#goulash)
	- [Quick Start](#quick-start)
		- [Installation](#installation)
		- [Example](#example)
	- [Functions](#functions)
		- [All](#all)
		- [Any](#any)
		- [Chunk](#chunk)
		- [Compact](#compact)
		- [Concat](#concat)
		- [Contains](#contains)
		- [Difference](#difference)
		- [Filter](#filter)
		- [FlatMap](#flatmap)
		- [Fold](#fold)
		- [ForEach](#foreach)
		- [GroupBy](#groupby)
		- [Intersection](#intersection)
		- [Join](#join)
		- [Keys](#keys)
		- [Map](#map)
		- [Max](#max)
		- [Min](#min)
		- [MinMax](#minmax)
		- [Reduce](#reduce)
		- [Sort](#sort)
		- [Ternary](#ternary)
		- [Union](#union)
		- [Unique](#unique)
		- [Values](#values)

## Quick Start

### Installation

```sh
go get github.com/farbodsalimi/goulash
```

### Example

```go
package main

import (
	"fmt"

	__ "github.com/farbodsalimi/goulash"
)

type Player struct {
	id    int
	name  string
	score float64
	bonus float64
}

func main() {
	players := []Player{
		{id: 1, name: "David", score: 20, bonus: 15},
		{id: 2, name: "Jessica", score: 10, bonus: 25},
		{id: 3, name: "Alex", score: 40, bonus: 45},
		{id: 4, name: "Tom", score: 30, bonus: 25},
	}

	totalScore := __.Reduce(__.Map(__.Filter(players, func(p Player) bool {
		return len(p.name) > 4
	}), func(p Player) float64 {
		return p.score + p.bonus
	}), func(acc, newScore float64) float64 {
		return acc + newScore
	}, 0)

	fmt.Println(totalScore)
}
```

## Functions

### All

```go
result := __.All([]float64{3.1, 4.1, 5.1, 1.1, 0})
fmt.Println(result) // false
```

### Any

```go
result := __.Any([]float64{3.1, 4.1, 5.1, 1.1, 0})
fmt.Println(result) // true
```

### Chunk

```go
chunks := __.Chunk([]int{1, 2, 3, 4}, 2)
fmt.Println(chunks) // [[1 2] [3 4]]
```

### Compact

```go
compacted := __.Compact([]int{0, 1, 2, 3})
fmt.Println(compacted) // [1 2 3]
```

### Concat

```go
concatenated := __.Concat([]int{1, 2, 3}, []int{4, 5, 6, 7})
fmt.Println(concatenated) // [1 2 3 4 5 6 7]
```

### Contains

```go
result := __.Contains([]float64{3.1, 4.1, 5.1, 1.1, -2.1}, -2.1)
fmt.Println(result) // true
```

### Difference

```go
diff := __.Difference([]int{1, 2, 3}, []int{1, 2})
fmt.Println(diff) // [3]
```

### Filter

```go
filtered := __.Filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
	return n%2 == 1
})
fmt.Println(filtered) // [1 3 5]
```

### FlatMap

```go
flatten := __.FlatMap([]string{"it's sunny in", "", "California"}, func(s string) []string {
	return strings.Split(s, " ")
})
fmt.Println(flatten) // ["it's", "sunny", "in", "", "California"]
```

### Fold

```go
folded := __.Fold([]int{1, 2, 3}, 0, func(a int, b int) int {
	return a + b
})
fmt.Println(folded) // 6
```

### ForEach

```go
var forEachResult [][]int
__.ForEach([]int{1, 2, 3}, func(value int, args ...any) {
	index := args[0].(int)
	forEachResult = append(forEachResult, []int{index, value})
})
fmt.Println(forEachResult) // [[0 1] [1 2] [2 3]]
```

### GroupBy

```go
grouped := __.GroupBy([]float64{6.1, 4.2, 6.3}, math.Floor)
fmt.Println(grouped) // map[4:[4.2] 6:[6.1 6.3]]
```

### Intersection

```go
intersected := __.Intersection([]string{"a", "b", "c", "d", "e"}, []string{"d", "e"})
fmt.Println(intersected) // ["d", "e"]
```

### Join

```go
joined := __.Join([]uint{1, 2, 3, 4}, ", ")
fmt.Println(joined) // 1, 2, 3, 4
```

```go
joined := __.Join([]float64{1.1, 2.1, 3.1, 1.1, 2.1, 3.1}, "~")
fmt.Println(joined) // 1.1~2.1~3.1~1.1~2.1~3.1
```

```go
joined := __.Join([]string{"a", "b", "c", "d"}, "*")
fmt.Println(joined) // a*b*c*d
```

### Keys

```go
keys := __.Keys(map[string]string{"key1": "value", "key2": "value", "key3": "value"})
fmt.Println(keys) // ["key1", "key2", "key3"
```

### Map

```go
mapped := __.Map([]int{1, 2, 3}, math.Pow10)
fmt.Println(mapped) // [10 100 1000]
```

### Max

```go
maxResult := __.Max([]float32{6.1, 4.2, 6.3}...)
fmt.Println(maxResult) // 6.3
```

### Min

```go
minResult := __.Min([]float64{6.1, 4.2, 6.3}...)
fmt.Println(minResult) // 4.2
```

### MinMax

```go
min, max := __.MinMax([]float64{6.1, 4.2, 6.3}...)
fmt.Println(min, max) // 4.2 6.3
```

### Reduce

```go
reduced := __.Reduce([]uint{6, 7, 8}, func(a uint, b uint) uint {
	return a + b
}, 0)
fmt.Println(reduced) // 21
```

### Sort

```go
sorted := __.Sort([]int{6, 1, 2, 3, -1, 0, 4, 7, 5})
fmt.Println(sorted) // [-1 0 1 2 3 4 5 6 7]
```

### Ternary

```go
result := __.Ternary((2 + 2) == 4, "yup", "nope")
fmt.Println(result) // yup
```

### Union

```go
unified := __.Union([][]int{{1, 2}, {2, 3, 4}, {3, 4, 5, 6, 7}}...)
fmt.Println(unified) // [1 2 3 4 5 6 7]
```

### Unique

```go
uniq := __.Unique([]int{1, 1, 1, 1, 2, 3})
fmt.Println(uniq) // [1 2 3]
```

### Values

```go
values := __.Values(map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"})
fmt.Println(values) // ["value1", "value2", "value3"
```
