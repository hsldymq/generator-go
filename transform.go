//go:build goexperiment.rangefunc

package goiter

import "iter"

func PickK[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return T21(seq, func(k K, _ V) K {
		return k
	})
}

func PickV[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return T21(seq, func(_ K, v V) V {
		return v
	})
}

func SwapKV[K, V any](seq iter.Seq2[K, V]) iter.Seq2[V, K] {
	return T2(seq, func(k K, v V) (V, K) {
		return v, k
	})
}

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
