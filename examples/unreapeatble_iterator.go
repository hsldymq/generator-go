package examples

import (
	"fmt"
	"github.com/hsldymq/goiter"
)

func OnceDemo() {
	data := []int{1, 2, 3, 4, 5, 6}

	// this is an example for Once function.
	// iterator returned by Once can be only used once.
	iterator := goiter.Once(goiter.SliceElem(data))

	// break the loop or finish the iteration will make the iterator unusable.
	for v := range iterator {
		// only the first 3 elements will be printed: 1 2 3
		fmt.Printf("%d ", v)
		if v == 3 {
			break
		}
	}
	fmt.Println()

	// so this won't print anything.
	for v := range iterator {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func ContinuableOnceDemo() {
	data := []int{1, 2, 3, 4, 5, 6}

	// this is an example for ContinuableOnce function.
	iterator := goiter.ContinuableOnce(goiter.SliceElem(data))

	// break the loop midway will not make the iterator unusable until all elements have been yielded.
	for v := range iterator {
		// only the first 3 elements will be printed: 1 2 3
		fmt.Printf("%d ", v)
		if v == 3 {
			break
		}
	}
	fmt.Println()

	// so this will print the remaining elements: 4 5 6
	for v := range iterator {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// and this won't print anything because all elements have been yielded, iterator is unusable now.
	for v := range iterator {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
