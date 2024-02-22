//go:build goexperiment.rangefunc

package goiter

import "iter"

func Map[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func MapVal[K comparable, V any](m map[K]V) iter.Seq[V] {
	return PickV[K, V](Map(m))
}

func MapKey[K comparable, V any](m map[K]V) iter.Seq[K] {
	return PickK[K, V](Map(m))
}
