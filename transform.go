//go:build goexperiment.rangefunc

package goiter

import "iter"

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
