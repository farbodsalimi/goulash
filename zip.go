package goulash

type Pair[T any, U any] struct {
	First  T
	Second U
}

func Zip[T any, U any](ts []T, us []U) []Pair[T, U] {
	minLen := Min(len(ts), len(us))

	result := make([]Pair[T, U], minLen)
	for i := 0; i < minLen; i++ {
		result[i] = Pair[T, U]{
			First:  ts[i],
			Second: us[i],
		}
	}

	return result
}
