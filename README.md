# goulash

[![Pipeline](https://github.com/farbodsalimi/goulash/actions/workflows/go.yml/badge.svg)](https://github.com/farbodsalimi/goulash/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/farbodsalimi/goulash)](https://goreportcard.com/report/github.com/farbodsalimi/goulash)

goulash provides a bunch of useful functional programming helpers leveraging generics.

**Table of Contents**

- [goulash](#goulash)
  - [Functions](#functions)
    - [Chunk](#chunk)
    - [Compact](#compact)
    - [Concat](#concat)
    - [Filter](#filter)
    - [Difference](#difference)
    - [ForEach](#foreach)
    - [GroupBy](#groupby)
    - [Map](#map)
    - [Max](#max)
    - [Min](#min)
    - [MinMax](#minmax)
    - [Reduce](#reduce)
    - [Sort](#sort)
    - [Union](#union)
    - [Unique](#unique)

## Functions

### Chunk

```go
	chunks := goulash.Chunk([]int{1, 2, 3, 4}, 2)
	fmt.Println(chunks) // [[1 2] [3 4]]
```

### Compact

```go
	compacted := goulash.Compact([]int{0, 1, 2, 3})
	fmt.Println(compacted) // [1 2 3]
```

### Concat

```go
	concatenated := goulash.Concat([]int{1, 2, 3}, []int{4, 5, 6, 7})
	fmt.Println(concatenated) // [1 2 3 4 5 6 7]
```

### Filter

```go
	filtered := goulash.Filter([]int{1, 2, 3, 4, 5, 6},
		func(n int) bool {
			return n%2 == 1
		})
	fmt.Println(filtered) // [1 3 5]
```

### Difference

```go
	diff := goulash.Difference([]int{1, 2, 3}, []int{1, 2})
	fmt.Println(diff) // [3]
```

### ForEach

```go
	var forEachResult [][]int
	goulash.ForEach([]int{1, 2, 3}, func(value int, args ...any) {
		index := args[0].(int)
		forEachResult = append(forEachResult, []int{index, value})
	})
	fmt.Println(forEachResult) // [[0 1] [1 2] [2 3]]
```

### GroupBy

```go
	grouped := goulash.GroupBy([]float64{6.1, 4.2, 6.3}, math.Floor)
	fmt.Println(grouped) // map[4:[4.2] 6:[6.1 6.3]]
```

### Map

```go
	mapped := goulash.Map([]int{1, 2, 3}, math.Pow10)
	fmt.Println(mapped) // [10 100 1000]
```

### Max

```go
	maxResult := goulash.Max([]float32{6.1, 4.2, 6.3}...)
	fmt.Println(maxResult) // 6.3
```

### Min

```go
	minResult := goulash.Min([]float64{6.1, 4.2, 6.3}...)
	fmt.Println(minResult) // 4.2
```

### MinMax

```go
	min, max := goulash.MinMax([]float64{6.1, 4.2, 6.3}...)
	fmt.Println(min, max) // 4.2 6.3
```

### Reduce

```go
	reduced := goulash.Reduce([]uint{6, 7, 8}, func(a uint, b uint) uint {
		return a + b
	}, 0)
	fmt.Println(reduced) // 21
```

### Sort

```go
	sorted := goulash.Sort([]int{6, 1, 2, 3, -1, 0, 4, 7, 5})
	fmt.Println(sorted) // [-1 0 1 2 3 4 5 6 7]
```

### Union

```go
	unified := goulash.Union([][]int{{1, 2}, {2, 3, 4}, {3, 4, 5, 6, 7}}...)
	fmt.Println(unified) // [1 2 3 4 5 6 7]
```

### Unique

```go
	uniq := goulash.Unique([]int{1, 1, 1, 1, 2, 3})
	fmt.Println(uniq) // [1 2 3]
```
