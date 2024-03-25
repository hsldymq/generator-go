//go:build goexperiment.rangefunc

package goiter

import "iter"

type Combined[T1, T2 any] struct {
    V1 T1
    V2 T2
}

type ZippedE[T1, T2 any] struct {
    V1  T1
    OK1 bool
    V2  T2
    OK2 bool
}

// PickV1 returns an iterator that yields the first element of each 2-tuple provided by the input iterator.
func PickV1[T1, T2 any](seq iter.Seq2[T1, T2]) iter.Seq[T1] {
    return Transform21(seq, func(v1 T1, _ T2) T1 {
        return v1
    })
}

// PickV2 returns an iterator that yields the second element of each 2-tuple provided by the input iterator.
func PickV2[T1, T2 any](seq iter.Seq2[T1, T2]) iter.Seq[T2] {
    return Transform21(seq, func(_ T1, v2 T2) T2 {
        return v2
    })
}

// Swap returns an iterator that yields new 2-tuples by swapping the positions of the elements within each 2-Tuple provided by the input iterator.
func Swap[T1, T2 any](seq iter.Seq2[T1, T2]) iter.Seq2[T2, T1] {
    return Transform2(seq, func(v1 T1, v2 T2) (T2, T1) {
        return v2, v1
    })
}

// Combine returns an iterator that yields combined values, where each value contains the elements of the 2-Tuple provided by the input iterator.
func Combine[T1, T2 any](seq iter.Seq2[T1, T2]) iter.Seq[*Combined[T1, T2]] {
    return Transform21(seq, func(v1 T1, v2 T2) *Combined[T1, T2] {
        return &Combined[T1, T2]{
            V1: v1,
            V2: v2,
        }
    })
}

// Transform returns an iterator, it yields new values by applying the transformer function to each value provided by the input iterator.
func Transform[In, Out any](
    seq iter.Seq[In],
    transformer func(In) Out,
) iter.Seq[Out] {
    return func(yield func(Out) bool) {
        next, stop := iter.Pull(seq)
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

// Transform2 is the iter.Seq2 version of Transform function.
func Transform2[InT1, InT2, OutT1, OutT2 any](
    seq iter.Seq2[InT1, InT2],
    transformer func(InT1, InT2) (OutT1, OutT2),
) iter.Seq2[OutT1, OutT2] {
    return func(yield func(OutT1, OutT2) bool) {
        next, stop := iter.Pull2(seq)
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

// Transform12 is similar to Transform, but it yields 2-tuple values after transformation instead of single-values.
func Transform12[In, OutT1, OutT2 any](
    seq iter.Seq[In],
    transformer func(In) (OutT1, OutT2),
) iter.Seq2[OutT1, OutT2] {
    return func(yield func(OutT1, OutT2) bool) {
        next, stop := iter.Pull(seq)
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

// Transform21 is similar to Transform2, but it only yields transform single-values instead of 2-tuple values
func Transform21[InT1, InT2, Out any](
    seq iter.Seq2[InT1, InT2],
    transformer func(InT1, InT2) Out,
) iter.Seq[Out] {
    return func(yield func(Out) bool) {
        next, stop := iter.Pull2(seq)
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

// Zip is like python's zip function, it takes two iterators and returns an iterator of combined structs,
// where the i-th struct contains the i-th element from each of the argument iterators.
// when two iterators have different lengths, the resulting iterator will stop when the shorter one stops.
// for example:
//
//	seq1 yields  1   2   3   4   5
//	seq2 yields "a" "b" "c"
//	Zip(seq1, seq2) will yield {1, "a"} {2, "b"} {3, "c"}
func Zip[T1, T2 any](seq1 iter.Seq[T1], seq2 iter.Seq[T2]) iter.Seq[*Combined[T1, T2]] {
    return ZipAs(seq1, seq2, func(zipped *ZippedE[T1, T2]) *Combined[T1, T2] {
        return &Combined[T1, T2]{
            V1: zipped.V1,
            V2: zipped.V2,
        }
    })
}

// ZipAs is a more general version of Zip.
// if exhaust parameter is true, the resulting iterator will not stop until both input iterators stop, and ZippedE.OK1 and ZippedE.OK2 will be false when the corresponding iterator stops.
func ZipAs[InT1, InT2, Out any](seq1 iter.Seq[InT1], seq2 iter.Seq[InT2], transformer func(*ZippedE[InT1, InT2]) Out, exhaust ...bool) iter.Seq[Out] {
    return func(yield func(Out) bool) {
        shouldExhaust := false
        if len(exhaust) > 0 {
            shouldExhaust = exhaust[0]
        }

        p1, stop1 := iter.Pull(seq1)
        defer stop1()
        p2, stop2 := iter.Pull(seq2)
        defer stop2()

        for {
            in1, ok1 := p1()
            in2, ok2 := p2()
            if !ok1 && !ok2 {
                return
            }
            if (!ok1 || !ok2) && !shouldExhaust {
                return
            }

            out := transformer(&ZippedE[InT1, InT2]{
                V1:  in1,
                OK1: ok1,
                V2:  in2,
                OK2: ok2,
            })
            if !yield(out) {
                return
            }
        }
    }
}

// ToSlice converts an iterator to a slice.
func ToSlice[T any](seq iter.Seq[T]) []T {
    var result []T
    for each := range seq {
        result = append(result, each)
    }
    return result
}

// ToMap converts an iterator that yields 2-tuple to a map, where the first element of the tuple is the key and the second element is the value.
func ToMap[T1 comparable, T2 any](seq iter.Seq2[T1, T2]) map[T1]T2 {
    result := make(map[T1]T2)
    for key, val := range seq {
        result[key] = val
    }
    return result
}

// ToMapBy transform every element provided from the input iterator to a key-value pair, and then returns a map.
func ToMapBy[T any, OutK comparable, OutV any](
    seq iter.Seq[T],
    transformer func(T) (OutK, OutV),
) map[OutK]OutV {
    result := make(map[OutK]OutV)
    for v := range seq {
        key, val := transformer(v)
        result[key] = val
    }
    return result
}

// ToMapBy2 is similar to ToMapBy, but it takes 2-Tuple from the input iterator.
func ToMapBy2[InT1 any, InT2 any, OutK comparable, OutV any](
    seq iter.Seq2[InT1, InT2],
    transformer func(InT1, InT2) (OutK, OutV),
) map[OutK]OutV {
    result := make(map[OutK]OutV)
    for v1, v2 := range seq {
        key, val := transformer(v1, v2)
        result[key] = val
    }
    return result
}
