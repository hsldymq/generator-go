//go:build goexperiment.rangefunc

package goiter

import (
	"iter"
)

func Int[T tInt, S sInt](start, stop T, _step ...S) iter.Seq[T] {
	step := int64(1)
	inc := true
	if start > stop {
		inc = false
		step = -1
	}

	if len(_step) > 0 {
		step = int64(_step[0])
	}

	emptyFunc := func(yield func(T) bool) {}
	if step == 0 {
		return emptyFunc
	}
	if (inc && step < 0) || (!inc && step > 0) {
		return emptyFunc
	}

	return func(yield func(T) bool) {
		curr := start
		for {
			if !yield(curr) {
				return
			}

			if inc {
				s := T(step)
				if s > intMax(curr) {
					return
				}
				next := curr + s
				// TODO
				_ = next
			} else {

			}
		}
	}
}

type tInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | sInt
}

type sInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func intMin[T tInt](v T) T {
	allOnes := ^T(0)
	if allOnes < 0 {
		return allOnes
	}
	return 0
}

func intMax[T tInt](v T) T {
	allOnes := ^T(0)
	if allOnes < 0 {
		return allOnes ^ (1 << (countBits(allOnes) - 1))
	}

	return allOnes
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
