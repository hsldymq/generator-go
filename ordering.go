//go:build goexperiment.rangefunc

package goiter

import (
	"cmp"
	"iter"
	"slices"
)

type Tuple[T1, T2 any] struct {
	V1 T1
	V2 T2
}

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

// OrderV1 sorts the 2-tuples of the input iterator by first element and returns a new iterator whose elements are arranged in ascending or descending order.
// If the second parameter is true, the tuples are arranged in descending order.
// For example:
//
//	since iter.Map(map[string]int{"bob":2, "eve":3, "alice":1}) yields the 2-tuples in arbitrary order
//	then OrderV1(iter.Map(map[string]int{"bob":2, "eve":3, "alice":1}))       will yield (alice, 1) (bob 2) (eve 3)
//	and  OrderV1(iter.Map(map[string]int{"bob":2, "eve":3, "alice":1}), true) will yield (eve 3) (bob 2) (alice, 1).
func OrderV1[T1 cmp.Ordered, T2 any](seq iter.Seq2[T1, T2], desc ...bool) iter.Seq2[T1, T2] {
	var cmpFunc func(a *Tuple[T1, T2], b *Tuple[T1, T2]) int

	if len(desc) > 0 && desc[0] {
		cmpFunc = func(a *Tuple[T1, T2], b *Tuple[T1, T2]) int {
			return cmp.Compare(b.V1, a.V1)
		}
	} else {
		cmpFunc = func(a *Tuple[T1, T2], b *Tuple[T1, T2]) int {
			return cmp.Compare(a.V1, b.V1)
		}
	}

	return doOrderBy2(seq, cmpFunc, slices.SortFunc[[]*Tuple[T1, T2], *Tuple[T1, T2]])
}

// OrderV2 is like OrderV1, but it sorts by the second element of the 2-tuples.
func OrderV2[T1 any, T2 cmp.Ordered](seq iter.Seq2[T1, T2], desc ...bool) iter.Seq2[T1, T2] {
	var cmpFunc func(a *Tuple[T1, T2], b *Tuple[T1, T2]) int
	if len(desc) > 0 && desc[0] {
		cmpFunc = func(a *Tuple[T1, T2], b *Tuple[T1, T2]) int {
			return cmp.Compare(b.V2, a.V2)
		}
	} else {
		cmpFunc = func(a *Tuple[T1, T2], b *Tuple[T1, T2]) int {
			return cmp.Compare(a.V2, b.V2)
		}
	}

	return doOrderBy2(seq, cmpFunc, slices.SortFunc[[]*Tuple[T1, T2], *Tuple[T1, T2]])
}

// OrderBy accepts a comparison function and returns a new iterator that yields elements sorted by the comparison function.
func OrderBy[T any](seq iter.Seq[T], cmp func(T, T) int) iter.Seq[T] {
	return doOrderBy(seq, cmp, slices.SortFunc[[]T, T])
}

// OrderBy2 is the iter.seq2 version of OrderBy.
func OrderBy2[T1, T2 any](seq iter.Seq2[T1, T2], cmp func(*Tuple[T1, T2], *Tuple[T1, T2]) int) iter.Seq2[T1, T2] {
	return doOrderBy2(seq, cmp, slices.SortFunc[[]*Tuple[T1, T2], *Tuple[T1, T2]])
}

// StableOrderBy is like OrderBy, but it uses a stable sort algorithm.
func StableOrderBy[T any](seq iter.Seq[T], cmp func(T, T) int) iter.Seq[T] {
	return doOrderBy(seq, cmp, slices.SortStableFunc[[]T, T])
}

// StableOrderBy2 is like OrderBy2, but it uses a stable sort algorithm.
func StableOrderBy2[T1, T2 any](seq iter.Seq2[T1, T2], cmp func(*Tuple[T1, T2], *Tuple[T1, T2]) int) iter.Seq2[T1, T2] {
	return doOrderBy2(seq, cmp, slices.SortStableFunc[[]*Tuple[T1, T2], *Tuple[T1, T2]])
}

type tSortFunc[S ~[]T, T any] func(x S, cmp func(a, b T) int)

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

func doOrderBy2[T1, T2 any](seq iter.Seq2[T1, T2], cmp func(*Tuple[T1, T2], *Tuple[T1, T2]) int, sortFunc tSortFunc[[]*Tuple[T1, T2], *Tuple[T1, T2]]) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		tuples := make([]*Tuple[T1, T2], 0)
		for v1, v2 := range seq {
			tuples = append(tuples, &Tuple[T1, T2]{
				V1: v1,
				V2: v2,
			})
		}

		sortFunc(tuples, cmp)
		for _, each := range tuples {
			if !yield(each.V1, each.V2) {
				return
			}
		}
	}
}
