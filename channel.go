//go:build goexperiment.rangefunc

package goiter

import "iter"

func ChanIter[T any](c <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}
