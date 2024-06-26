package goiter

import "iter"

// Count counts the number of elements yielded by the input iterator.
func Count[TIter SeqX[T], T any](iterator TIter) int {
    count := 0
    for _ = range iterator {
        count++
    }
    return count
}

// Count2 counts the number of elements yielded by the input iterator.
func Count2[TIter Seq2X[T1, T2], T1 any, T2 any](iterator TIter) int {
    count := 0
    for _, _ = range iterator {
        count++
    }
    return count
}

// Reduce is basically Reduce function in functional programming.
// The following example uses Reduce to sum up the numbers from 1 to 10:
//
//  sum := goiter.Reduce(goiter.Range(1, 10), 0, func(acc, v int) int {
//      return acc + v
// 	})
func Reduce[TIter SeqX[T], TAcc any, T any](
    iterator TIter,
    init TAcc,
    folder func(TAcc, T) TAcc,
) TAcc {
    var result = init
    for v := range iterator {
        result = folder(result, v)
    }
    return result
}

// Scan is similar to Reduce function, but it returns an iterator that will yield the reduced value of each round.
// So, the following code will create an iterator that yields 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, where each value is the sum of numbers from 1 to the current number.
//  iterator := goiter.Scan(goiter.Range(1, 10), 0, func(acc, v int) int {
//      return acc + v
// 	})
func Scan[TIter SeqX[T], TAcc any, T any](
    iterator TIter,
    init TAcc,
    folder func(TAcc, T) TAcc,
) Iterator[TAcc] {
    return func(yield func(TAcc) bool) {
        next, stop := iter.Pull(iter.Seq[T](iterator))
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
