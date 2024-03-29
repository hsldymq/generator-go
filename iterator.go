//go:build goexperiment.rangefunc

package goiter

import (
    "iter"
)

type SeqX[T any] interface {
    ~func(yield func(T) bool)
}

type Iterator[T any] iter.Seq[T]

func (it Iterator[T]) Seq() iter.Seq[T] {
    return iter.Seq[T](it)
}

func (it Iterator[T]) OrderBy(cmp func(T, T) int) Iterator[T] {
    return OrderBy(it, cmp)
}

func (it Iterator[T]) StableOrderBy(cmp func(T, T) int) Iterator[T] {
    return StableOrderBy(it, cmp)
}

func (it Iterator[T]) Filter(predicate func(T) bool) Iterator[T] {
    return Filter(it, predicate)
}

func (it Iterator[T]) Take(n int) Iterator[T] {
    return Take(it, n)
}

func (it Iterator[T]) TakeLast(n int) Iterator[T] {
    return TakeLast(it, n)
}

func (it Iterator[T]) Skip(n int) Iterator[T] {
    return Skip(it, n)
}

func (it Iterator[T]) SkipLast(n int) Iterator[T] {
    return SkipLast(it, n)
}

func (it Iterator[T]) Concat(its ...Iterator[T]) Iterator[T] {
    return Concat(it, its...)
}

func (it Iterator[T]) Reverse() Iterator[T] {
    return Reverse(it)
}

func (it Iterator[T]) Count() int {
    return Count(it)
}

func (it Iterator[T]) ToSlice() []T {
    return ToSlice(it)
}
