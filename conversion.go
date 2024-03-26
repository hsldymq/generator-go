//go:build goexperiment.rangefunc

package goiter

// ToSlice converts an iterator to a slice.
func ToSlice[T any](it Iterator[T]) []T {
	var result []T
	for each := range it {
		result = append(result, each)
	}
	return result
}

// ToMap converts an iterator that yields 2-tuple to a map, where the first element of the tuple is the key and the second element is the value.
func ToMap[T1 comparable, T2 any](it Iterator2[T1, T2]) map[T1]T2 {
	result := make(map[T1]T2)
	for key, val := range it {
		result[key] = val
	}
	return result
}

// ToMapAs transform every element provided from the input iterator to a key-value pair, and then returns a map.
func ToMapAs[T any, OutK comparable, OutV any](
	it Iterator[T],
	transformer func(T) (OutK, OutV),
) map[OutK]OutV {
	result := make(map[OutK]OutV)
	for v := range it {
		key, val := transformer(v)
		result[key] = val
	}
	return result
}

// ToMap2As is similar to ToMapAs, but it takes 2-Tuple from the input iterator.
func ToMap2As[InT1 any, InT2 any, OutK comparable, OutV any](
	it Iterator2[InT1, InT2],
	transformer func(InT1, InT2) (OutK, OutV),
) map[OutK]OutV {
	result := make(map[OutK]OutV)
	for v1, v2 := range it {
		key, val := transformer(v1, v2)
		result[key] = val
	}
	return result
}
