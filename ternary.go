package goulash

// Ternary evaluates a bool expression and returns the result of one of the two expressions, depending on whether the bool expression evaluates to true or false
func Ternary[T any](condition bool, exprIfTrue T, exprIfFalse T) T {
	if condition {
		return exprIfTrue
	}
	return exprIfFalse
}
