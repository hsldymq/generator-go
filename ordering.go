package goiter

import (
	"cmp"
	"iter"
	"slices"
)

// Order sorts the elements of the input iterator and returns a new iterator whose elements are arranged in ascending or descending order.
// If the second parameter is true, the elements are arranged in descending order.
// For example:
//
//	since iter.SliceElem([]int{2, 3, 1})) yields 2 3 1
//	then Order(iter.SliceElem([]int{2, 3, 1}))       will yield 1 2 3
//	and  Order(iter.SliceElem([]int{2, 3, 1}), true) will yield 3 2 1.
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

// OrderK sorts the key-values of the input iterator by key and returns a new iterator whose elements are arranged in ascending or descending order.
// If the second parameter is true, the elements are arranged in descending order by key.
// For example:
//
//	since iter.Map(map[string]int{"bob":2, "eve":3, "alice":1}) yields the key-values in arbitrary order
//	then OrderK(iter.Map(map[string]int{"bob":2, "eve":3, "alice":1}))       will yield (alice, 1) (bob 2) (eve 3)
//	and  OrderK(iter.Map(map[string]int{"bob":2, "eve":3, "alice":1}), true) will yield (eve 3) (bob 2) (alice, 1).
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

// OrderV is like OrderK, but it sorts by values.
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

// OrderBy accepts a comparison function and returns a new iterator that yields elements sorted by the comparison function.
func OrderBy[T any](seq iter.Seq[T], cmp func(T, T) int) iter.Seq[T] {
	return doOrderBy(seq, cmp, slices.SortFunc[[]T, T])
}

// OrderBy2 is the iter.seq2 version of OrderBy.
func OrderBy2[K, V any](seq iter.Seq2[K, V], cmp func(*KV[K, V], *KV[K, V]) int) iter.Seq2[K, V] {
	return doOrderBy2(seq, cmp, slices.SortFunc[[]*KV[K, V], *KV[K, V]])
}

// StableOrderBy is like OrderBy, but it uses a stable sort algorithm.
func StableOrderBy[T any](seq iter.Seq[T], cmp func(T, T) int) iter.Seq[T] {
	return doOrderBy(seq, cmp, slices.SortStableFunc[[]T, T])
}

// StableOrderBy2 is like OrderBy2, but it uses a stable sort algorithm.
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
