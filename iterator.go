//go:build goexperiment.rangefunc

package goiter

import "iter"

type Iterator[T any] iter.Seq[T]

type Iterator2[T1, T2 any] iter.Seq2[T1, T2]
