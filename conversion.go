//go:build goexperiment.rangefunc

package goiter

// ToSlice converts an iterator to a slice.
func ToSlice[TIter SeqX[T], T any](iterator TIter) []T {
	var result []T
	for each := range iterator {
		result = append(result, each)
	}
	return result
}

// ToMap converts an iterator that yields 2-tuple to a map, where the first element of the tuple is the key and the second element is the value.
func ToMap[TIter Seq2X[T1, T2], T1 comparable, T2 any](iterator TIter) map[T1]T2 {
	result := make(map[T1]T2)
	for key, val := range iterator {
		result[key] = val
	}
	return result
}

// ToMapAs transform every element provided from the input iterator to a key-value pair, and then returns a map.
func ToMapAs[TIter SeqX[T], TK comparable, TV any, T any](
	iterator TIter,
	transformer func(T) (TK, TV),
) map[TK]TV {
	result := make(map[TK]TV)
	for v := range iterator {
		key, val := transformer(v)
		result[key] = val
	}
	return result
}

// ToMap2As is similar to ToMapAs, but it takes 2-Tuple from the input iterator.
func ToMap2As[TIter Seq2X[T1, T2], TK comparable, TV any, T1 any, T2 any](
	iterator TIter,
	transformer func(T1, T2) (TK, TV),
) map[TK]TV {
	result := make(map[TK]TV)
	for v1, v2 := range iterator {
		key, val := transformer(v1, v2)
		result[key] = val
	}
	return result
}
