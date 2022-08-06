# go-iterator

![go workflow](https://github.com/MTBorg/go-iterator/actions/workflows/go.yml/badge.svg)

`go-iterator` provides a generic implementation of the
[iterator](https://en.wikipedia.org/wiki/Iterator_pattern) pattern.

It provides the two types

```go
// An Iterable implements how to iterate over something one element at a time.
type Iterable[T any] interface {
	Next() *T
}

// Iterator iterates over an iterable object.
type Iterator[T any] struct {
	iterable Iterable[T]
}
```

`Iterator` provides some common iterator methods (e.g. `Filter`, `Take`,
`ForEach`) and can be instantiated using slices and maps (see `FromSlice` and
`FromMap`) as well as any type that implements the `Iterable` interface.

## Examples

### Basic iterator operations

```go
even := FromSlice([]int{1,2,3,4}).Filter(func(i int) bool {return i % 2 == 0})
fmt.Println(even.Collect()) // [2 4]
count := FromSlice([]int{1,2,3,4}).Count()
fmt.Println(count) // 4
reversed := FromSlice([]int{1,2,3,4}).Reverse()
fmt.Println(reversed.Collect()) // [4 3 2 1]

FromMap(map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
}).ForEach(func (i int){fmt.Printf("%d ", i)}) // 1 2 3 (order not guaranteed)

i1, i2 := FromSlice([]string{"hello", "world"}), FromSlice([]string{"goodbye", "now"})
fmt.Println(i1.Chain(i2).Collect()) // ["hello" "world" "goodbye" "now"]
```

### Custom iterable implementations

#### Binary tree traversal

```go
package iterator

import "fmt"

type BinaryTree struct {
	root Node
}

type BinaryTreeIterator struct {
	stack []*Node
}

func NewBinaryTreeIterator(tree BinaryTree) *BinaryTreeIterator {
	return &BinaryTreeIterator{
		stack: []*Node{&tree.root},
	}
}

// Next traverses the binary tree in a left depth-first order
func (iter *BinaryTreeIterator) Next() *Node {
	if len(iter.stack) == 0 {
		return nil
	}
	next := iter.stack[0]
	iter.stack = iter.stack[1:]
	if next.Right != nil {
		iter.stack = append([]*Node{next.Right}, iter.stack...)
	}
	if next.Left != nil {
		iter.stack = append([]*Node{next.Left}, iter.stack...)
	}
	return next
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func ExampleIterate() {
	tree := BinaryTree{
		root: Node{
			5,
			&Node{
				3,
				&Node{Value: 2},
				&Node{Value: 4},
			},
			&Node{
				9,
				&Node{Value: 7},
				&Node{
					10,
					nil,
					&Node{Value: 11},
				},
			},
		},
	}

	// Traverse tree
	values := []int{}
	iter := Iterate[Node](NewBinaryTreeIterator(tree))
	iter.ForEach(func(n Node) { values = append(values, n.Value) })
	fmt.Println(values)

	// Traverse tree in reverse order
	values = []int{}
	iter = Iterate[Node](NewBinaryTreeIterator(tree))
	iter.Reverse().ForEach(func(n Node) { values = append(values, n.Value) })
	fmt.Println(values)

	// Filter all node whose values are greater than 5
	values = []int{}
	iter = Iterate[Node](NewBinaryTreeIterator(tree))
	iter.Filter(func(n Node) bool { return n.Value > 5 }).ForEach(func(n Node) { values = append(values, n.Value) })
	fmt.Println(values)
	// Output:
	// [5 3 2 4 9 7 10 11]
	// [11 10 7 9 4 2 3 5]
	// [9 7 10 11]
}
```

#### Infinite Iterator

```go
package iterator

import (
	"fmt"
	"math/rand"
)

type InfiniteIterator struct {
	max int
}

func (iter InfiniteIterator) Next() *int {
	n := rand.Intn(iter.max)
	return &n
}

func ExampleIterate_infiniteIterator() {
	iter := Iterate[int](InfiniteIterator{10})

	// Take 10 (pseudo-)random integers from the interval [0,10) and group them
	// into even and odd numbers
	even, odd := iter.Take(10).Partition(func(i int) bool { return i%2 == 0 })
	fmt.Println(even.Collect())
	fmt.Println(odd.Collect())
}
```
