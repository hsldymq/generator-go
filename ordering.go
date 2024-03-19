package goiter

import (
	"cmp"
	"iter"
	"slices"
)

func Order[T cmp.Ordered](seq iter.Seq[T], desc ...bool) iter.Seq[T] {
	var cmpFunc func(a T, b T) int
	if len(desc) > 0 && desc[0] {
		cmpFunc = func(a, b T) int {
			return cmp.Compare(b, a)
		}
	} else {
		cmpFunc = func(a, b T) int {
			return cmp.Compare(a, b)
		}
	}

	return doOrderBy(seq, cmpFunc, slices.SortFunc[[]T, T])
}

func OrderK[K cmp.Ordered, V any](seq iter.Seq2[K, V], desc ...bool) iter.Seq2[K, V] {
	var cmpFunc func(a *KV[K, V], b *KV[K, V]) int

	if len(desc) > 0 && desc[0] {
		cmpFunc = func(a *KV[K, V], b *KV[K, V]) int {
			return cmp.Compare(b.K, a.K)
		}
	} else {
		cmpFunc = func(a *KV[K, V], b *KV[K, V]) int {
			return cmp.Compare(a.K, b.K)
		}
	}

	return doOrderBy2(seq, cmpFunc, slices.SortFunc[[]*KV[K, V], *KV[K, V]])
}

func OrderV[K any, V cmp.Ordered](seq iter.Seq2[K, V], desc ...bool) iter.Seq2[K, V] {
	var cmpFunc func(a *KV[K, V], b *KV[K, V]) int
	if len(desc) > 0 && desc[0] {
		cmpFunc = func(a *KV[K, V], b *KV[K, V]) int {
			return cmp.Compare(b.V, a.V)
		}
	} else {
		cmpFunc = func(a *KV[K, V], b *KV[K, V]) int {
			return cmp.Compare(a.V, b.V)
		}
	}

	return doOrderBy2(seq, cmpFunc, slices.SortFunc[[]*KV[K, V], *KV[K, V]])
}

func OrderBy[T any](seq iter.Seq[T], cmp func(T, T) int) iter.Seq[T] {
	return doOrderBy(seq, cmp, slices.SortFunc[[]T, T])
}

func OrderBy2[K, V any](seq iter.Seq2[K, V], cmp func(*KV[K, V], *KV[K, V]) int) iter.Seq2[K, V] {
	return doOrderBy2(seq, cmp, slices.SortFunc[[]*KV[K, V], *KV[K, V]])
}

func StableOrderBy[T any](seq iter.Seq[T], cmp func(T, T) int) iter.Seq[T] {
	return doOrderBy(seq, cmp, slices.SortStableFunc[[]T, T])
}

func StableOrderBy2[K, V any](seq iter.Seq2[K, V], cmp func(*KV[K, V], *KV[K, V]) int) iter.Seq2[K, V] {
	return doOrderBy2(seq, cmp, slices.SortStableFunc[[]*KV[K, V], *KV[K, V]])
}

type tSortFunc[S ~[]E, E any] func(x S, cmp func(a, b E) int)

func doOrderBy[T any](seq iter.Seq[T], cmp func(T, T) int, sortFunc tSortFunc[[]T, T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		s := make([]T, 0)
		for each := range seq {
			s = append(s, each)
		}

		sortFunc(s, cmp)
		for _, each := range s {
			if !yield(each) {
				return
			}
		}
	}
}

func doOrderBy2[K, V any](seq iter.Seq2[K, V], cmp func(*KV[K, V], *KV[K, V]) int, sortFunc tSortFunc[[]*KV[K, V], *KV[K, V]]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		kvs := make([]*KV[K, V], 0)
		for k, v := range seq {
			kvs = append(kvs, &KV[K, V]{
				K: k,
				V: v,
			})
		}

		sortFunc(kvs, cmp)
		for _, each := range kvs {
			if !yield(each.K, each.V) {
				return
			}
		}
	}
}
