package goiter

import (
    "iter"
)

// PickV1 returns an iterator that yields the first element of each 2-tuple provided by the input iterator.
// For example:
//  iterator := goiter.Slice([]string{"a", "b", "c"})       // iterator will yield (1, "a") (2, "b") (3, "c")
//  newIterator := goiter.PickV1(iterator)                  // after calling PickV1, newIterator will yield 1 2 3
func PickV1[TIter Seq2X[T1, T2], T1, T2 any](iterator TIter) Iterator[T1] {
    return Transform21(iterator, func(v1 T1, _ T2) T1 {
        return v1
    })
}

// PickV2 returns an iterator that yields the second element of each 2-tuple provided by the input iterator.
// For example:
//  iterator := goiter.Slice([]string{"a", "b", "c"})       // iterator will yield (1, "a") (2, "b") (3, "c")
//  newIterator := goiter.PickV2(iterator)                  // after calling PickV2, newIterator will yield "a" "b" "c"
func PickV2[TIter Seq2X[T1, T2], T1, T2 any](iterator TIter) Iterator[T2] {
    return Transform21(iterator, func(_ T1, v2 T2) T2 {
        return v2
    })
}

// Swap returns an iterator that yields new 2-tuples by swapping the positions of the elements within each 2-Tuple provided by the input iterator.
// For example:
//  iterator := goiter.Slice([]string{"a", "b", "c"})       // iterator will yield (1, "a") (2, "b") (3, "c")
//  newIterator := goiter.Swap(iterator)                    // after calling Swap, newIterator will yield ("a", 1) ("b", 2) ("c", 3)
func Swap[TIter Seq2X[T1, T2], T1, T2 any](iterator TIter) Iterator2[T2, T1] {
    return Transform2(iterator, func(v1 T1, v2 T2) (T2, T1) {
        return v2, v1
    })
}

// Transform returns an iterator, it yields new values by applying the transformer function to each value provided by the input iterator.
// For example:
//  iterator := goiter.SliceElems([]int{1, 2, 3})               // iterator will yield 1 2 3
//  newIterator := goiter.Transform(iterator, strconv.Itoa)     // after calling Transform, newIterator will yield "1" "2" "3"
func Transform[TIter SeqX[T], TOut, T any](
    iterator TIter,
    transformer func(T) TOut,
) Iterator[TOut] {
    return func(yield func(TOut) bool) {
        next, stop := iter.Pull(iter.Seq[T](iterator))
        defer stop()
        for {
            v, ok := next()
            if !ok {
                return
            }
            out := transformer(v)
            if !yield(out) {
                return
            }
        }
    }
}

// Transform2 is the iter.Seq2 version of Transform function, it transforms each 2-tuple values from the input iterator and yields the transformed 2-tuple values.
func Transform2[TIter Seq2X[T1, T2], TOut1, TOut2, T1, T2 any](
    iterator TIter,
    transformer func(T1, T2) (TOut1, TOut2),
) Iterator2[TOut1, TOut2] {
    return func(yield func(TOut1, TOut2) bool) {
        next, stop := iter.Pull2(iter.Seq2[T1, T2](iterator))
        defer stop()
        for {
            v1, v2, ok := next()
            if !ok {
                return
            }
            out1, out2 := transformer(v1, v2)
            if !yield(out1, out2) {
                return
            }
        }
    }
}

// Transform12 is similar to Transform, but it transforms each value from the input iterator to 2-tuple values.
// For example:
//  iterator := goiter.SliceElems([]string{"hello", "golang"})               // iterator will yield "hello" "golang"
//  newIterator := goiter.Transform12(iterator, func(s string) (string, int) {
//      return s, len(s)
//  })     // after calling Transform12, newIterator will yield ("hello", 5) ("golang", 6)
func Transform12[TIter SeqX[T], OutT1, OutT2, T any](
    iterator TIter,
    transformer func(T) (OutT1, OutT2),
) Iterator2[OutT1, OutT2] {
    return func(yield func(OutT1, OutT2) bool) {
        next, stop := iter.Pull(iter.Seq[T](iterator))
        defer stop()
        for {
            v, ok := next()
            if !ok {
                return
            }
            out1, out2 := transformer(v)
            if !yield(out1, out2) {
                return
            }
        }
    }
}

// Transform21 is similar to Transform12, but it is reversed, it transforms each 2-tuple value from the input iterator to single-values.
func Transform21[TIter Seq2X[T1, T2], TOut, T1, T2 any](
    iterator TIter,
    transformer func(T1, T2) TOut,
) Iterator[TOut] {
    return func(yield func(TOut) bool) {
        next, stop := iter.Pull2(iter.Seq2[T1, T2](iterator))
        defer stop()
        for {
            v1, v2, ok := next()
            if !ok {
                return
            }
            out := transformer(v1, v2)
            if !yield(out) {
                return
            }
        }
    }
}
