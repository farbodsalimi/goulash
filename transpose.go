package goulash

func Transpose[T any](matrix [][]T) [][]T {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]T{}
	}

	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]T, cols)
	for i := range result {
		result[i] = make([]T, rows)
	}

	for i := range rows {
		if len(matrix[i]) < cols {
			cols = len(matrix[i])
		}
		for j := range cols {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}
