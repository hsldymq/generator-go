//go:build goexperiment.rangefunc

package goiter

import (
	"iter"
)

// SourceFunc delegates data retrieval from elsewhere.
type SourceFunc[T any] func() T

// Slice returns an iterator that allows you to traverse a slice in a forward or reverse direction.
func Slice[T any](s []T, backward ...bool) iter.Seq2[int, T] {
	return SliceSource(func() []T { return s }, backward...)
}

// SliceElem only yields the elements of a slice.
func SliceElem[T any](s []T, backward ...bool) iter.Seq[T] {
	return PickV(Slice(s, backward...))
}

// SliceSource is like the Slice function, but the slice is taken from the input SourceFunc.
// You might use this function in scenarios like this:
// When a slice is encapsulated within a struct, it offers the capability to traverse elements by exposing an iterator.
// However, there might come a time when the structure replaces the original slice with a new one.
// When the already exposed iterator is traversed again, we hope it traverses the new slice instead of the old one.
// Therefore, by providing a SourceFunc, the moment of obtaining the slice is delayed until the iterator is traversed.
func SliceSource[T any](source SourceFunc[[]T], backward ...bool) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		s := source()
		if len(backward) == 0 || !backward[0] {
			for idx, elem := range s {
				if !yield(idx, elem) {
					return
				}
			}
		} else {
			for i := len(s) - 1; i >= 0; i-- {
				if !yield(i, s[i]) {
					return
				}
			}
		}
	}
}

// SliceSourceElem is the SourceFunc version of SliceElem function.
// see comments of SliceSource function for more details.
func SliceSourceElem[T any](source SourceFunc[[]T], backward ...bool) iter.Seq[T] {
	return PickV(SliceSource(source, backward...))
}

// Map returns an iterator that allows you to traverse a map.
func Map[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return MapSource(func() map[K]V { return m })
}

// MapVal only yields the values of a map.
func MapVal[K comparable, V any](m map[K]V) iter.Seq[V] {
	return PickV(Map(m))
}

// MapKey only yields the keys of a map.
func MapKey[K comparable, V any](m map[K]V) iter.Seq[K] {
	return PickK(Map(m))
}

// MapSource is the SourceFunc version of Map function.
// see comments of SliceSource function for more details.
func MapSource[K comparable, V any](source SourceFunc[map[K]V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		m := source()
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

// MapSourceVal is the SourceFunc version of MapVal function.
// see comments of SliceSource function for more details.
func MapSourceVal[K comparable, V any](source SourceFunc[map[K]V]) iter.Seq[V] {
	return PickV(MapSource(source))
}

// MapSourceKey is the SourceFunc version of MapKey function.
// see comments of SliceSource function for more details.
func MapSourceKey[K comparable, V any](source SourceFunc[map[K]V]) iter.Seq[K] {
	return PickK(MapSource(source))
}

// Channel yields the values from a channel, it will stop when the channel is closed.
func Channel[V any](c <-chan V) iter.Seq[V] {
	return ChannelSource(func() <-chan V { return c })
}

// ChannelSource is the SourceFunc version of Channel function.
// see comments of SliceSource function for more details.
func ChannelSource[V any](source SourceFunc[<-chan V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		c := source()
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}

// Empty returns an empty iterator.
func Empty[V any]() iter.Seq[V] {
	return func(yield func(V) bool) {
		return
	}
}

// Empty2 is iter.Seq2 version of Empty
func Empty2[K any, V any]() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		return
	}
}

// Count counts the number of elements yielded by the input iterator.
func Count[V any](seq iter.Seq[V]) int {
	count := 0
	for _ = range seq {
		count++
	}
	return count
}

// Count2 counts the number of elements yielded by the input iterator.
func Count2[K any, V any](seq iter.Seq2[K, V]) int {
	count := 0
	for _, _ = range seq {
		count++
	}
	return count
}
