//go:build goexperiment.rangefunc

package goiter

type tInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64
}

func intMin[T tInt](v T) T {
	ones := ^T(0)
	if ones < 0 {
		return ^(ones ^ (1 << (countBits(ones) - 1)))
	}
	return 0
}

func intMax[T tInt](v T) T {
	ones := ^T(0)
	if ones < 0 {
		return ones ^ (1 << (countBits(ones) - 1))
	}
	return ones
}

func countBits[T tInt](v T) int {
	v = 1
	for _, bits := range [4]int{8, 16, 32} {
		if v<<bits == 0 {
			return bits
		}
	}
	return 64
}
