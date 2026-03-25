package goulash

// Sliding creates a sliding window of the specified size over a slice.
// It returns a slice of slices, where each sub-slice contains consecutive elements
// from the original slice. For example, Sliding([1,2,3,4,5], 3) returns [[1,2,3], [2,3,4], [3,4,5]].
func Sliding[T any](slice []T, size int) [][]T {
	if size <= 0 || len(slice) == 0 {
		return [][]T{}
	}

	if size > len(slice) {
		return [][]T{append([]T{}, slice...)}
	}

	var result [][]T
	for i := 0; i <= len(slice)-size; i++ {
		window := append([]T{}, slice[i:i+size]...)
		result = append(result, window)
	}
	return result
}
