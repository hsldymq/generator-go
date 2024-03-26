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

type ZippedE[T1, T2 any] struct {
	V1  T1
	OK1 bool
	V2  T2
	OK2 bool
}

// Combine returns an iterator that yields combined values, where each value contains the elements of the 2-Tuple provided by the input iterator.
func Combine[T1, T2 any](it Iterator2[T1, T2]) Iterator[*Combined[T1, T2]] {
	return Transform21(it, Combiner[T1, T2])
}

// Zip is like python's zip function, it takes two iterators and returns an iterator of combined structs,
// where the i-th struct contains the i-th element from each of the argument iterators.
// when two iterators have different lengths, the resulting iterator will stop when the shorter one stops.
// for example:
//
//	it1 yields  1   2   3   4   5
//	it2 yields "a" "b" "c"
//	Zip(it1, it2) will yield {1, "a"} {2, "b"} {3, "c"}
func Zip[T1, T2 any](it1 Iterator[T1], it2 Iterator[T2]) Iterator[*Combined[T1, T2]] {
	return ZipAs(it1, it2, func(zipped *ZippedE[T1, T2]) *Combined[T1, T2] {
		return &Combined[T1, T2]{
			V1: zipped.V1,
			V2: zipped.V2,
		}
	})
}

// ZipAs is a more general version of Zip.
// if exhaust parameter is true, the resulting iterator will not stop until both input iterators stop, and ZippedE.OK1 and ZippedE.OK2 will be false when the corresponding iterator stops.
func ZipAs[InT1, InT2, Out any](it1 Iterator[InT1], it2 Iterator[InT2], transformer func(*ZippedE[InT1, InT2]) Out, exhaust ...bool) Iterator[Out] {
	return func(yield func(Out) bool) {
		shouldExhaust := false
		if len(exhaust) > 0 {
			shouldExhaust = exhaust[0]
		}

		p1, stop1 := iter.Pull(it1.Seq())
		defer stop1()
		p2, stop2 := iter.Pull(it2.Seq())
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

// Concat returns an iterator that allows you to traverse multiple iterators in sequence.
func Concat[T any](it Iterator[T], its ...Iterator[T]) Iterator[T] {
	if len(its) == 0 {
		return it
	}

	return func(yield func(T) bool) {
		for v := range it {
			if !yield(v) {
				return
			}
		}
		for _, eachIt := range its {
			for v := range eachIt {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Concat2 returns an iterator that allows you to traverse multiple iterators in sequence.
func Concat2[T1 any, T2 any](it Iterator2[T1, T2], its ...Iterator2[T1, T2]) Iterator2[T1, T2] {
	if len(its) == 0 {
		return it
	}

	return func(yield func(T1, T2) bool) {
		for v1, v2 := range it {
			if !yield(v1, v2) {
				return
			}
		}

		for _, eachIt := range its {
			for v1, v2 := range eachIt {
				if !yield(v1, v2) {
					return
				}
			}
		}
	}
}
