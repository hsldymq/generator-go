//go:build goexperiment.rangefunc

package goiter

import "iter"

type KV[K, V any] struct {
	K K
	V V
}

type Zipped[T1, T2 any] struct {
	V1 T1
	V2 T2
}

type ZippedE[T1, T2 any] struct {
	V1  T1
	OK1 bool
	V2  T2
	OK2 bool
}

// PickK yields the keys of a sequence of key-value pairs.
func PickK[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return T21(seq, func(k K, _ V) K {
		return k
	})
}

// PickV yields the values of a sequence of key-value pairs.
func PickV[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return T21(seq, func(_ K, v V) V {
		return v
	})
}

// SwapKV yields key-value pairs after swapping the position of the keys and values obtained by the input iterator.
func SwapKV[K, V any](seq iter.Seq2[K, V]) iter.Seq2[V, K] {
	return T2(seq, func(k K, v V) (V, K) {
		return v, k
	})
}

// CombineKV yields KVPairs after combining the keys and values obtained from the input iterator.
func CombineKV[K, V any](seq iter.Seq2[K, V]) iter.Seq[*KV[K, V]] {
	return T21(seq, func(k K, v V) *KV[K, V] {
		return &KV[K, V]{K: k, V: v}
	})
}

// T1 return a transforming iterator, where T stands for transform.
// It applies the transformer function to the values obtained from the input iterator, and then yields the result.
func T1[In, Out any](
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
			if !yield(transformer(v)) {
				return
			}
		}
	}
}

// T2 is similar to T1, but it obtains the key-value pairs from the input iterator and yields new key-value pairs after transformation.
func T2[InT1, InT2, OutT1, OutT2 any](
	seq iter.Seq2[InT1, InT2],
	transformer func(InT1, InT2) (OutT1, OutT2),
) iter.Seq2[OutT1, OutT2] {
	return func(yield func(OutT1, OutT2) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()
		for {
			k, v, ok := next()
			if !ok {
				return
			}
			if !yield(transformer(k, v)) {
				return
			}
		}
	}
}

// T12 is similar to T1, but it obtains the values from the input iterator and yields new key-value pairs after transformation.
func T12[In, OutT1, OutT2 any](
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
			if !yield(transformer(v)) {
				return
			}
		}
	}
}

// T21 is similar to T2, but it only yields transform values without keys.
func T21[InT1, InT2, Out any](
	seq iter.Seq2[InT1, InT2],
	transformer func(InT1, InT2) Out,
) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()
		for {
			k, v, ok := next()
			if !ok {
				return
			}
			if !yield(transformer(k, v)) {
				return
			}
		}
	}
}

func Zip[T1, T2 any](seq1 iter.Seq[T1], seq2 iter.Seq[T2]) iter.Seq[*Zipped[T1, T2]] {
	return ZipAs(seq1, seq2, func(zipped *ZippedE[T1, T2]) *Zipped[T1, T2] {
		return &Zipped[T1, T2]{
			V1: zipped.V1,
			V2: zipped.V2,
		}
	})
}

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

// ToMap converts an iterator of key-value pairs to a map.
func ToMap[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	result := make(map[K]V)
	for k, v := range seq {
		result[k] = v
	}
	return result
}

// ToMapBy takes every element from the input iterator, applies the transformer function to it, and then stores the result in a map.
func ToMapBy[T any, OutK comparable, OutV any](
	seq iter.Seq[T],
	transformer func(T) (OutK, OutV),
) map[OutK]OutV {
	result := make(map[OutK]OutV)
	for v := range seq {
		kk, vv := transformer(v)
		result[kk] = vv
	}
	return result
}

// ToMapBy2 is similar to ToMapBy, but it takes 2-tuple from the input iterator.
func ToMapBy2[InT1 any, InT2 any, OutK comparable, OutV any](
	seq iter.Seq2[InT1, InT2],
	transformer func(InT1, InT2) (OutK, OutV),
) map[OutK]OutV {
	result := make(map[OutK]OutV)
	for k, v := range seq {
		kk, vv := transformer(k, v)
		result[kk] = vv
	}
	return result
}
