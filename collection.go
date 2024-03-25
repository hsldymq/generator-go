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
	return PickV2(Slice(s, backward...))
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

// SliceSourceElem is like SliceElem function, it serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func SliceSourceElem[T any](source SourceFunc[[]T], backward ...bool) iter.Seq[T] {
	return PickV2(SliceSource(source, backward...))
}

// Map returns an iterator that allows you to traverse a map.
func Map[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return MapSource(func() map[K]V { return m })
}

// MapVal yields only values of a map in arbitrary order.
func MapVal[K comparable, V any](m map[K]V) iter.Seq[V] {
	return PickV2(Map(m))
}

// MapKey yields only keys of a map in arbitrary order.
func MapKey[K comparable, V any](m map[K]V) iter.Seq[K] {
	return PickV1(Map(m))
}

// MapSource is like Map function, it serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func MapSource[K comparable, V any](source SourceFunc[map[K]V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		m := source()
		for key, val := range m {
			if !yield(key, val) {
				return
			}
		}
	}
}

// MapSourceVal is like MapVal function, it serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func MapSourceVal[K comparable, V any](source SourceFunc[map[K]V]) iter.Seq[V] {
	return PickV2(MapSource(source))
}

// MapSourceKey is like MapKey function, it serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func MapSourceKey[K comparable, V any](source SourceFunc[map[K]V]) iter.Seq[K] {
	return PickV1(MapSource(source))
}

// Channel yields the values from a channel, it will stop when the channel is closed.
func Channel[T any](c <-chan T) iter.Seq[T] {
	return ChannelSource(func() <-chan T { return c })
}

// ChannelSource is like Channel function, it serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func ChannelSource[T any](source SourceFunc[<-chan T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		c := source()
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}

// IterSource serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func IterSource[T any](source SourceFunc[iter.Seq[T]]) iter.Seq[T] {
	return func(yield func(T) bool) {
		seq := source()
		for v := range seq {
			if !yield(v) {
				return
			}
		}
	}
}

// Empty returns an empty iterator.
func Empty[T any]() iter.Seq[T] {
	return func(yield func(T) bool) {
		return
	}
}

// Empty2 is iter.Seq2 version of Empty
func Empty2[T1 any, T2 any]() iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		return
	}
}
