package goiter

import "iter"

func MapIter[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func MapIterVal[K comparable, V any](m map[K]V) iter.Seq[V] {
	return PickV[K, V](MapIter(m))
}

func MapIterKey[K comparable, V any](m map[K]V) iter.Seq[K] {
	return PickK[K, V](MapIter(m))
}
