package goiter

import "iter"

func PickK[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return Transform21(seq, func(k K, _ V) K {
		return k
	})
}

func PickV[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return Transform21(seq, func(_ K, v V) V {
		return v
	})
}

func Transform11[In, Out any](
	seq iter.Seq[In],
	transformFunc func(In) Out,
) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v, ok := next()
			if !ok {
				return
			}
			if !yield(transformFunc(v)) {
				return
			}
		}
	}
}

func Transform12[In, OutK, OutV any](
	seq iter.Seq[In],
	transformFunc func(In) (OutK, OutV),
) iter.Seq2[OutK, OutV] {
	return func(yield func(OutK, OutV) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v, ok := next()
			if !ok {
				return
			}
			if !yield(transformFunc(v)) {
				return
			}
		}
	}
}

func Transform21[InK, InV, Out any](
	seq iter.Seq2[InK, InV],
	transformFunc func(InK, InV) Out,
) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()
		for {
			k, v, ok := next()
			if !ok {
				return
			}
			if !yield(transformFunc(k, v)) {
				return
			}
		}
	}
}

func Transform22[InK, InV, OutK, OutV any](
	seq iter.Seq2[InK, InV],
	transformFunc func(InK, InV) (OutK, OutV),
) iter.Seq2[OutK, OutV] {
	return func(yield func(OutK, OutV) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()
		for {
			k, v, ok := next()
			if !ok {
				return
			}
			if !yield(transformFunc(k, v)) {
				return
			}
		}
	}
}
