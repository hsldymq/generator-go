//go:build goexperiment.rangefunc

package goiter

import "iter"

func Combiner[T1, T2 any](v1 T1, v2 T2) *Combined[T1, T2] {
    return &Combined[T1, T2]{
        V1: v1,
        V2: v2,
    }
}

type Combined[T1, T2 any] struct {
    V1 T1
    V2 T2
}

type Zipped[T1, T2 any] struct {
    V1  T1
    OK1 bool
    V2  T2
    OK2 bool
}

// Combine returns an iterator that yields combined values, where each value contains the elements of the 2-Tuple provided by the input iterator.
func Combine[TIter Seq2X[T1, T2], T1, T2 any](iterator TIter) Iterator[*Combined[T1, T2]] {
    return Transform21(iterator, Combiner[T1, T2])
}

// Zip is like python's zip function, it takes two iterators and returns an iterator that yields 2-tuples,
// where the first element of the 2-tuple is from the first iterator, and the second element is from the second iterator
// when two iterators have different lengths, the resulting iterator will stop when the shorter one stops.
// for example:
//
//	it1 yields  1   2   3   4   5
//	it2 yields "a" "b" "c"
//	Zip(it1, it2) will yield (1, "a") (2, "b") (3, "c")
func Zip[TIter1 SeqX[T1], TIter2 SeqX[T2], T1, T2 any](
    iterator1 TIter1,
    iterator2 TIter2,
) Iterator2[T1, T2] {
    return func(yield func(T1, T2) bool) {
        p1, stop1 := iter.Pull(iter.Seq[T1](iterator1))
        defer stop1()
        p2, stop2 := iter.Pull(iter.Seq[T2](iterator2))
        defer stop2()

        for {
            v1, ok1 := p1()
            v2, ok2 := p2()
            if !ok1 || !ok2 {
                return
            }

            if !yield(v1, v2) {
                return
            }
        }
    }
}

// ZipAs is a more general version of Zip.
// if exhaust parameter is true, the resulting iterator will not stop until both input iterators stop, and Zipped.OK1 and Zipped.OK2 will be false when the corresponding iterator stops.
func ZipAs[TIter1 SeqX[T1], TIter2 SeqX[T2], TOut, T1, T2 any](
    iterator1 TIter1,
    iterator2 TIter2,
    transformer func(*Zipped[T1, T2]) TOut,
    exhaust bool,
) Iterator[TOut] {
    return func(yield func(TOut) bool) {
        p1, stop1 := iter.Pull(iter.Seq[T1](iterator1))
        defer stop1()
        p2, stop2 := iter.Pull(iter.Seq[T2](iterator2))
        defer stop2()

        for {
            in1, ok1 := p1()
            in2, ok2 := p2()
            if !ok1 && !ok2 {
                return
            }
            if (!ok1 || !ok2) && !exhaust {
                return
            }

            out := transformer(&Zipped[T1, T2]{
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

// Concat returns an iterator that allows you to traverse multiple iterators in sequence.
func Concat[TIter SeqX[T], T any](
    iterator TIter,
    more ...TIter,
) Iterator[T] {
    if len(more) == 0 {
        return Iterator[T](iterator)
    }

    return func(yield func(T) bool) {
        for v := range iterator {
            if !yield(v) {
                return
            }
        }
        for _, it := range more {
            for v := range it {
                if !yield(v) {
                    return
                }
            }
        }
    }
}

// Concat2 returns an iterator that allows you to traverse multiple iterators in sequence.
func Concat2[TIter Seq2X[T1, T2], T1 any, T2 any](
    iterator TIter,
    more ...TIter,
) Iterator2[T1, T2] {
    if len(more) == 0 {
        return Iterator2[T1, T2](iterator)
    }

    return func(yield func(T1, T2) bool) {
        for v1, v2 := range iterator {
            if !yield(v1, v2) {
                return
            }
        }

        for _, eachIt := range more {
            for v1, v2 := range eachIt {
                if !yield(v1, v2) {
                    return
                }
            }
        }
    }
}
