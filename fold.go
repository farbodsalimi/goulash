package goulash

func Fold[T any](slice []T, empty T, fn func(a T, b T) T) T {
	curr := empty
	for _, value := range slice {
		curr = fn(curr, value)
	}
	return curr
}
