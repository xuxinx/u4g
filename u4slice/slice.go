package u4slice

import "math"

func Chunks(s []interface{}, size int) [][]interface{} {
	l := int(math.Ceil(float64(len(s)) / float64(size)))
	chunks := make([][]interface{}, l)

	idx := 0
	for i := 0; i < l; i++ {
		idx = i * size
		if i != l-1 {
			chunks[i] = s[idx : idx+size]
		} else {
			chunks[i] = s[idx:]
		}
	}

	return chunks
}
