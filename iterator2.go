//go:build goexperiment.rangefunc

package goiter

import (
	"iter"
)

type Iterator2[T1, T2 any] iter.Seq2[T1, T2]

func (it Iterator2[T1, T2]) Seq() iter.Seq2[T1, T2] {
	return iter.Seq2[T1, T2](it)
}

func (it Iterator2[T1, T2]) OrderBy(cmp func(*Combined[T1, T2], *Combined[T1, T2]) int) Iterator2[T1, T2] {
	return Order2By(it, cmp)
}

func (it Iterator2[T1, T2]) StableOrderBy(cmp func(*Combined[T1, T2], *Combined[T1, T2]) int) Iterator2[T1, T2] {
	return StableOrder2By(it, cmp)
}

func (it Iterator2[T1, T2]) Filter(cmp func(T1, T2) bool) Iterator2[T1, T2] {
	return Filter2(it, cmp)
}

func (it Iterator2[T1, T2]) Take(n int) Iterator2[T1, T2] {
	return Take2(it, n)
}

func (it Iterator2[T1, T2]) Skip(n int) Iterator2[T1, T2] {
	return Skip2(it, n)
}

func (it Iterator2[T1, T2]) Concat(its ...Iterator2[T1, T2]) Iterator2[T1, T2] {
	return Concat2(it, its...)
}

func (it Iterator2[T1, T2]) Reverse() Iterator2[T1, T2] {
	return Reverse2(it)
}

func (it Iterator2[T1, T2]) Count() int {
	return Count2(it)
}
