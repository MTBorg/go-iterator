package iterator_test

import (
	"fmt"
	"math/rand"

	"github.com/MTBorg/go-iterator"
)

type InfiniteIterator struct {
	max int
}

func (iter InfiniteIterator) Next() *int {
	n := rand.Intn(iter.max)
	return &n
}

func ExampleIterate_infiniteIterator() {
	iter := iterator.Iterate[int](InfiniteIterator{10})

	// Take 10 (pseudo-)random integers from the interval [0,10) and group them
	// into even and odd numbers
	even, odd := iter.Take(10).Partition(func(i int) bool { return i%2 == 0 })
	fmt.Println(even.Collect())
	fmt.Println(odd.Collect())
}
