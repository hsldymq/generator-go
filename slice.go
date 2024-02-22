//go:build goexperiment.rangefunc

package goiter

import (
	"iter"
)

func Slice[T any](s []T, backward ...bool) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		if len(backward) == 0 || !backward[0] {
			for idx, elem := range s {
				if !yield(idx, elem) {
					return
				}
			}
		} else {
			for i := len(s) - 1; i >= 0; i-- {
				if !yield(i, s[i]) {
					return
				}
			}
		}
	}
}

func SliceElem[T any](s []T, backward ...bool) iter.Seq[T] {
	return PickV[int, T](Slice(s, backward...))
}

func SliceIdx[T any](s []T, backward ...bool) iter.Seq[int] {
	return PickK[int, T](Slice(s, backward...))
}
