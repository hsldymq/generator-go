//go:build goexperiment.rangefunc

package goiter

import (
	"iter"
	"slices"
)

type Iterator2[T1, T2 any] iter.Seq2[T1, T2]

func (it Iterator2[T1, T2]) Seq() iter.Seq2[T1, T2] {
	return iter.Seq2[T1, T2](it)
}

func (it Iterator2[T1, T2]) OrderBy(cmp func(*Combined[T1, T2], *Combined[T1, T2]) int) Iterator2[T1, T2] {
	return doOrderBy2(it, cmp, slices.SortFunc[[]*Combined[T1, T2]])
}

func (it Iterator2[T1, T2]) StableOrderBy(cmp func(*Combined[T1, T2], *Combined[T1, T2]) int) Iterator2[T1, T2] {
	return doOrderBy2(it, cmp, slices.SortStableFunc[[]*Combined[T1, T2]])
}

func (it Iterator2[T1, T2]) Count() int {
	count := 0
	for _, _ = range it {
		count++
	}
	return count
}
