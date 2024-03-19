package goiter

import (
	"cmp"
	"iter"
	"slices"
)

func Order[T cmp.Ordered](seq iter.Seq[T], desc ...bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		s := make([]T, 0)
		for each := range seq {
			s = append(s, each)
		}

		if len(desc) > 0 && desc[0] {
			slices.SortFunc(s, func(a, b T) int {
				return cmp.Compare(b, a)
			})
		} else {
			slices.Sort(s)
		}
		for _, each := range s {
			if !yield(each) {
				return
			}
		}
	}
}

func OrderK[K cmp.Ordered, V any](seq iter.Seq2[K, V], desc ...bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		kvs := make([]*KVPair[K, V], 0)
		for k, v := range seq {
			kvs = append(kvs, &KVPair[K, V]{
				K: k,
				V: v,
			})
		}

		if len(desc) > 0 && desc[0] {
			slices.SortFunc(kvs, func(a *KVPair[K, V], b *KVPair[K, V]) int {
				return cmp.Compare(b.K, a.K)
			})
		} else {
			slices.SortFunc(kvs, func(a *KVPair[K, V], b *KVPair[K, V]) int {
				return cmp.Compare(a.K, b.K)
			})
		}
		for _, each := range kvs {
			if !yield(each.K, each.V) {
				return
			}
		}
	}
}

func OrderV[K any, V cmp.Ordered](seq iter.Seq2[K, V], desc ...bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		kvs := make([]*KVPair[K, V], 0)
		for k, v := range seq {
			kvs = append(kvs, &KVPair[K, V]{
				K: k,
				V: v,
			})
		}

		if len(desc) > 0 && desc[0] {
			slices.SortFunc(kvs, func(a *KVPair[K, V], b *KVPair[K, V]) int {
				return cmp.Compare(b.V, a.V)
			})
		} else {
			slices.SortFunc(kvs, func(a *KVPair[K, V], b *KVPair[K, V]) int {
				return cmp.Compare(a.V, b.V)
			})
		}
		for _, each := range kvs {
			if !yield(each.K, each.V) {
				return
			}
		}
	}
}

func OrderBy[T any](seq iter.Seq[T], cmp func(T, T) int) iter.Seq[T] {
	return doOrderBy(seq, cmp, slices.SortFunc[[]T, T])
}

func OrderBy2[K, V any](seq iter.Seq2[K, V], cmp func(*KVPair[K, V], *KVPair[K, V]) int) iter.Seq2[K, V] {
	return doOrderBy2(seq, cmp, slices.SortFunc[[]*KVPair[K, V], *KVPair[K, V]])
}

func OrderByStable[T any](seq iter.Seq[T], cmp func(T, T) int) iter.Seq[T] {
	return doOrderBy(seq, cmp, slices.SortStableFunc[[]T, T])
}

func OrderByStable2[K, V any](seq iter.Seq2[K, V], cmp func(*KVPair[K, V], *KVPair[K, V]) int) iter.Seq2[K, V] {
	return doOrderBy2(seq, cmp, slices.SortStableFunc[[]*KVPair[K, V], *KVPair[K, V]])
}

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

func doOrderBy2[K, V any](seq iter.Seq2[K, V], cmp func(*KVPair[K, V], *KVPair[K, V]) int, sortFunc tSortFunc[[]*KVPair[K, V], *KVPair[K, V]]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		kvs := make([]*KVPair[K, V], 0)
		for k, v := range seq {
			kvs = append(kvs, &KVPair[K, V]{
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

type tSortFunc[S ~[]E, E any] func(x S, cmp func(a, b E) int)
