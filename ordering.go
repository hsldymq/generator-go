//go:build goexperiment.rangefunc

package goiter

import (
	"cmp"
	"slices"
)

// Order sorts the elements of the input iterator and returns a new iterator whose elements are arranged in ascending or descending order.
// If the second parameter is true, the elements are arranged in descending order.
// For example:
//
//	since iter.SliceElem([]int{2, 3, 1})) yields 2 3 1
//	then Order(iter.SliceElem([]int{2, 3, 1}))       will yield 1 2 3
//	and  Order(iter.SliceElem([]int{2, 3, 1}), true) will yield 3 2 1.
//
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func Order[TIter SeqX[T], T cmp.Ordered](
	iterator TIter,
	desc ...bool,
) Iterator[T] {
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

	return doOrderBy(iterator, cmpFunc, slices.SortFunc[[]T, T])
}

// OrderV1 sorts the 2-tuples of the input iterator by first element and returns a new iterator whose elements are arranged in ascending or descending order.
// If the second parameter is true, the tuples are arranged in descending order.
// For example:
//
//	since iter.Map(map[string]int{"bob":2, "eve":3, "alice":1}) yields the 2-tuples in arbitrary order
//	then OrderV1(iter.Map(map[string]int{"bob":2, "eve":3, "alice":1}))       will yield (alice, 1) (bob 2) (eve 3)
//	and  OrderV1(iter.Map(map[string]int{"bob":2, "eve":3, "alice":1}), true) will yield (eve 3) (bob 2) (alice, 1).
//
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func OrderV1[TIter Seq2X[T1, T2], T1 cmp.Ordered, T2 any](
	iterator TIter,
	desc ...bool,
) Iterator2[T1, T2] {
	var cmpFunc func(a *Combined[T1, T2], b *Combined[T1, T2]) int

	if len(desc) > 0 && desc[0] {
		cmpFunc = func(a *Combined[T1, T2], b *Combined[T1, T2]) int {
			return cmp.Compare(b.V1, a.V1)
		}
	} else {
		cmpFunc = func(a *Combined[T1, T2], b *Combined[T1, T2]) int {
			return cmp.Compare(a.V1, b.V1)
		}
	}

	return doOrderBy2(iterator, cmpFunc, slices.SortFunc[[]*Combined[T1, T2], *Combined[T1, T2]])
}

// OrderV2 is like OrderV1, but it sorts by the second element of the 2-tuples.
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func OrderV2[TIter Seq2X[T1, T2], T1 any, T2 cmp.Ordered](
	iterator TIter,
	desc ...bool,
) Iterator2[T1, T2] {
	var cmpFunc func(a *Combined[T1, T2], b *Combined[T1, T2]) int
	if len(desc) > 0 && desc[0] {
		cmpFunc = func(a *Combined[T1, T2], b *Combined[T1, T2]) int {
			return cmp.Compare(b.V2, a.V2)
		}
	} else {
		cmpFunc = func(a *Combined[T1, T2], b *Combined[T1, T2]) int {
			return cmp.Compare(a.V2, b.V2)
		}
	}

	return doOrderBy2(iterator, cmpFunc, slices.SortFunc[[]*Combined[T1, T2], *Combined[T1, T2]])
}

// OrderBy accepts a comparison function and returns a new iterator that yields elements sorted by the comparison function.
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func OrderBy[TIter SeqX[T], T any](
	iterator TIter,
	cmp func(T, T) int,
) Iterator[T] {
	return doOrderBy(iterator, cmp, slices.SortFunc[[]T, T])
}

// Order2By is the Iterator2 version of OrderBy.
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func Order2By[TIter Seq2X[T1, T2], T1, T2 any](
	iterator TIter,
	cmp func(*Combined[T1, T2], *Combined[T1, T2]) int,
) Iterator2[T1, T2] {
	return doOrderBy2(iterator, cmp, slices.SortFunc[[]*Combined[T1, T2], *Combined[T1, T2]])
}

// StableOrderBy is like OrderBy, but it uses a stable sort algorithm.
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func StableOrderBy[TIter SeqX[T], T any](
	iterator TIter,
	cmp func(T, T) int,
) Iterator[T] {
	return doOrderBy(iterator, cmp, slices.SortStableFunc[[]T, T])
}

// StableOrder2By is like Order2By, but it uses a stable sort algorithm.
// be careful, if this function is used on iterators that has massive amount of data, it might consume a lot of memory.
func StableOrder2By[TIter Seq2X[T1, T2], T1, T2 any](
	iterator TIter,
	cmp func(*Combined[T1, T2], *Combined[T1, T2]) int,
) Iterator2[T1, T2] {
	return doOrderBy2(iterator, cmp, slices.SortStableFunc[[]*Combined[T1, T2], *Combined[T1, T2]])
}

type tSortFunc[S ~[]T, T any] func(x S, cmp func(a, b T) int)

func doOrderBy[TIter SeqX[T], T any](
	iterator TIter,
	cmp func(T, T) int,
	sortFunc tSortFunc[[]T, T],
) Iterator[T] {
	return func(yield func(T) bool) {
		s := make([]T, 0)
		for each := range iterator {
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

func doOrderBy2[TIter Seq2X[T1, T2], T1, T2 any](
	iterator TIter,
	cmp func(*Combined[T1, T2], *Combined[T1, T2]) int,
	sortFunc tSortFunc[[]*Combined[T1, T2], *Combined[T1, T2]],
) Iterator2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		tuples := make([]*Combined[T1, T2], 0)
		for v1, v2 := range iterator {
			tuples = append(tuples, &Combined[T1, T2]{
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
