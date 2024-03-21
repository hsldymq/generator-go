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
func T2[InK, InV, OutK, OutV any](
	seq iter.Seq2[InK, InV],
	transformer func(InK, InV) (OutK, OutV),
) iter.Seq2[OutK, OutV] {
	return func(yield func(OutK, OutV) bool) {
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
func T12[In, OutK, OutV any](
	seq iter.Seq[In],
	transformer func(In) (OutK, OutV),
) iter.Seq2[OutK, OutV] {
	return func(yield func(OutK, OutV) bool) {
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
func T21[InK, InV, Out any](
	seq iter.Seq2[InK, InV],
	transformer func(InK, InV) Out,
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
	return func(yield func(*Zipped[T1, T2]) bool) {
		p1, stop1 := iter.Pull(seq1)
		defer stop1()
		p2, stop2 := iter.Pull(seq2)
		defer stop2()

		for {
			v1, ok1 := p1()
			v2, ok2 := p2()
			if !ok1 || !ok2 {
				return
			}

			v := &Zipped[T1, T2]{
				V1: v1,
				V2: v2,
			}
			if !yield(v) {
				return
			}
		}
	}
}

func ZipAs[In1, In2, Out1, Out2 any](seq1 iter.Seq[In1], seq2 iter.Seq[In2], transformer func(In1, In2) (Out1, Out2)) iter.Seq[*Zipped[Out1, Out2]] {
	return func(yield func(*Zipped[Out1, Out2]) bool) {
		p1, stop1 := iter.Pull(seq1)
		defer stop1()
		p2, stop2 := iter.Pull(seq2)
		defer stop2()

		for {
			in1, ok1 := p1()
			in2, ok2 := p2()
			if !ok1 || !ok2 {
				return
			}

			out1, out2 := transformer(in1, in2)
			v := &Zipped[Out1, Out2]{
				V1: out1,
				V2: out2,
			}
			if !yield(v) {
				return
			}
		}
	}
}

func ToSlice[T any](seq iter.Seq[T]) []T {
	var result []T
	for each := range seq {
		result = append(result, each)
	}
	return result
}

func ToMap[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	result := make(map[K]V)
	for k, v := range seq {
		result[k] = v
	}
	return result
}

func ToMapBy[InK any, InV any, OutK comparable, OutV any](
	seq iter.Seq2[InK, InV],
	transformer func(InK, InV) (OutK, OutV),
) map[OutK]OutV {
	result := make(map[OutK]OutV)
	for k, v := range seq {
		kk, vv := transformer(k, v)
		result[kk] = vv
	}
	return result
}

func ToMapByV[V any, OutK comparable, OutV any](
	seq iter.Seq[V],
	transformer func(V) (OutK, OutV),
) map[OutK]OutV {
	result := make(map[OutK]OutV)
	for v := range seq {
		kk, vv := transformer(v)
		result[kk] = vv
	}
	return result
}
