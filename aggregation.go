//go:build goexperiment.rangefunc

package goiter

import "iter"

// Count counts the number of elements yielded by the input iterator.
func Count[T any](seq iter.Seq[T]) int {
    count := 0
    for _ = range seq {
        count++
    }
    return count
}

// Count2 counts the number of elements yielded by the input iterator.
func Count2[T1 any, T2 any](seq iter.Seq2[T1, T2]) int {
    count := 0
    for _, _ = range seq {
        count++
    }
    return count
}

// Fold is basically Reduce function in functional programming.
// so you want to sum up 1 to 10, using Fold, you can do it like this:
//
//	sum := goiter.Fold(goiter.Range(0, 11), 0, func(acc, v int) int { return acc + v })
func Fold[T any, Acc any](seq iter.Seq[T], init Acc, folder func(Acc, T) Acc) Acc {
    var result = init
    for v := range seq {
        result = folder(result, v)
    }
    return result
}

// Scan is similar to Fold, but unlike Fold, it reduces a complete sequence to a single value,
// Scan returns an iterator that will yield the folded value of each round.
func Scan[T any, Acc any](seq iter.Seq[T], init Acc, folder func(Acc, T) Acc) iter.Seq[Acc] {
    return func(yield func(Acc) bool) {
        next, stop := iter.Pull(seq)
        defer stop()

        acc := init
        for {
            v, ok := next()
            if !ok {
                return
            }

            acc = folder(acc, v)
            if !yield(acc) {
                return
            }
        }
    }
}
