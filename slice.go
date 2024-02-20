package goiter

import "iter"

func SliceIter[T any](s []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for idx, elem := range s {
			if !yield(idx, elem) {
				return
			}
		}
	}
}

func SliceIterElem[T any](s []T) iter.Seq[T] {
	return PickV[int, T](SliceIter(s))
}

func SliceIterIdx[T any](s []T) iter.Seq[int] {
	return PickK[int, T](SliceIter(s))
}
