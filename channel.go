//go:build goexperiment.rangefunc

package goiter

import "iter"

func Chan[T any](c <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}
