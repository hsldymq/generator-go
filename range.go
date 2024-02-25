//go:build goexperiment.rangefunc

package goiter

import (
	"iter"
	"math"
)

func Range[T tInt](start, stop T) iter.Seq[T] {
	return RangeStep(start, stop, 1)
}

// RangeStep extends the ability to range over integers, allowing iteration from any integer and stepping forward or backward.
// It is similar to Python's range function, but with some differences:
//  1. stepSize does not accept negative numbers. Whether iterating forward or backward, stepSize must be positive.
//     so you don't need to consider adjusting the sign of step according to the direction of iteration, It is the absolute value of the step parameter of Python range function.
//  2. Providing a value less than 0 for stepSize will not return an error, but will not yield any values.
func RangeStep[T tInt, S tInt](start, stop T, stepSize S) iter.Seq[T] {
	if stepSize <= 0 {
		// 0 will lead to infinite loops
		return func(yield func(T) bool) {}
	}

	step := uint64(stepSize)
	inc := true
	if start > stop {
		inc = false
	}

	if willOverflow(start, step, inc) {
		return func(yield func(T) bool) {
			yield(start)
		}
	}

	return func(yield func(T) bool) {
		curr := start
		for {
			if !yield(curr) {
				return
			}

			if inc {
				next := curr + T(step)
				if next >= stop || next < start || next <= curr {
					return
				}
				curr = next
			} else {
				next := curr - T(step)
				if next <= stop || next > start || next >= curr {
					return
				}
				curr = next
			}
		}
	}
}

func willOverflow[T tInt](v T, step uint64, inc bool) bool {
	tMax := int64(intMax(v))
	tMin := int64(intMin(v))

	if tMax != math.MaxInt64 && step >= uint64(tMax-tMin+1) {
		return true
	}
	if inc && v+T(step) < v {
		return true
	}
	if !inc && v-T(step) > v {
		return true
	}

	return false
}

type tInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
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
