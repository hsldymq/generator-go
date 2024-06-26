package goiter

// SourceFunc delegates data retrieval from elsewhere.
type SourceFunc[T any] func() T

// Slice returns an iterator that allows you to traverse a slice in a forward or reverse direction.
func Slice[S ~[]T, T any](s S, backward ...bool) Iterator2[int, T] {
    return SliceSource(func() S { return s }, backward...)
}

// SliceElems only yields the elements of a slice.
func SliceElems[S ~[]T, T any](s S, backward ...bool) Iterator[T] {
    return SliceSourceElems(func() S { return s }, backward...)
}

// SliceSource is like the Slice function, but the slice is taken from the input SourceFunc.
// You might use this function in scenarios like this:
// When a slice is encapsulated within a struct, it offers the capability to traverse elements by exposing an iterator.
// However, there might come a time when the structure replaces the original slice with a new one.
// When the already exposed iterator is traversed again, we hope it traverses the new slice instead of the old one.
// Therefore, by providing a SourceFunc, the moment of obtaining the slice is delayed until the iterator is traversed.
func SliceSource[S ~[]T, T any](source SourceFunc[S], backward ...bool) Iterator2[int, T] {
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

// SliceSourceElems is like SliceElems function, it serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func SliceSourceElems[S ~[]T, T any](source SourceFunc[S], backward ...bool) Iterator[T] {
    return func(yield func(T) bool) {
        s := source()
        if len(backward) == 0 || !backward[0] {
            for _, elem := range s {
                if !yield(elem) {
                    return
                }
            }
        } else {
            for i := len(s) - 1; i >= 0; i-- {
                if !yield(s[i]) {
                    return
                }
            }
        }
    }
}

// Map returns an iterator that allows you to traverse a map.
func Map[K comparable, V any](m map[K]V) Iterator2[K, V] {
    return MapSource(func() map[K]V { return m })
}

// MapKeys yields only keys of a map in arbitrary order.
func MapKeys[K comparable, V any](m map[K]V) Iterator[K] {
    return MapSourceKeys(func() map[K]V { return m })
}

// MapVals yields only values of a map in arbitrary order.
func MapVals[K comparable, V any](m map[K]V) Iterator[V] {
    return MapSourceVals(func() map[K]V { return m })
}

// MapSource is like Map function, it serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func MapSource[K comparable, V any](source SourceFunc[map[K]V]) Iterator2[K, V] {
    return func(yield func(K, V) bool) {
        m := source()
        for key, val := range m {
            if !yield(key, val) {
                return
            }
        }
    }
}

// MapSourceKeys is like MapKeys function, it serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func MapSourceKeys[K comparable, V any](source SourceFunc[map[K]V]) Iterator[K] {
    return func(yield func(K) bool) {
        m := source()
        for key := range m {
            if !yield(key) {
                return
            }
        }
    }
}

// MapSourceVals is like MapVals function, it serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func MapSourceVals[K comparable, V any](source SourceFunc[map[K]V]) Iterator[V] {
    return func(yield func(V) bool) {
        m := source()
        for _, val := range m {
            if !yield(val) {
                return
            }
        }
    }
}

// Chan yields the values from a channel, it will stop when the channel is closed.
func Chan[T any](c <-chan T) Iterator[T] {
    return ChanSource(func() <-chan T { return c })
}

// ChanSource is like Chan function, it serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func ChanSource[T any](source SourceFunc[<-chan T]) Iterator[T] {
    return func(yield func(T) bool) {
        c := source()
        for v := range c {
            if !yield(v) {
                return
            }
        }
    }
}

// SeqSource serves similar purposes as SliceSource, the difference is that the SourceFunc returns an iter.Seq-like iterator.
// see comments of SliceSource function for more details.
func SeqSource[TIter SeqX[T], T any](source SourceFunc[TIter]) Iterator[T] {
    return func(yield func(T) bool) {
        seq := source()
        for v := range seq {
            if !yield(v) {
                return
            }
        }
    }
}

// Seq2Source serves similar purposes as SliceSource.
// see comments of SliceSource function for more details.
func Seq2Source[TIter Seq2X[T1, T2], T1, T2 any](source SourceFunc[TIter]) Iterator2[T1, T2] {
    return func(yield func(T1, T2) bool) {
        seq := source()
        for v1, v2 := range seq {
            if !yield(v1, v2) {
                return
            }
        }
    }
}

// Items returns an iterator that simply yields the input values.
// So goiter.Items[any](1, true, 1.5, "hello") will yield 1, true, 1.5, "hello".
func Items[T any](t ...T) Iterator[T] {
    return SliceElems(t)
}

// Empty returns an empty iterator.
func Empty[T any]() Iterator[T] {
    return func(yield func(T) bool) {
        return
    }
}

// Empty2 is Iterator2 version of Empty
func Empty2[T1 any, T2 any]() Iterator2[T1, T2] {
    return func(yield func(T1, T2) bool) {
        return
    }
}
