package goulash

func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return [][]T{slice}
	}

	chunks := [][]T{}
	for i := 0; i < len(slice); i += size {
		end := min(i+size, len(slice))
		chunks = append(chunks, append([]T{}, slice[i:end]...))
	}
	return chunks
}
