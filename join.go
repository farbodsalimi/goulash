package goulash

import (
	"fmt"
	"strings"
)

// Join creates and returns a new string by concatenating all of the elements in a slice, separated by a specified delimiter
func Join[T any](slice []T, delimiter string) string {
	parts := make([]string, len(slice))
	for i, v := range slice {
		parts[i] = fmt.Sprint(v)
	}
	return strings.Join(parts, delimiter)
}
