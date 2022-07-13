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

func ExampleIterate_binaryTreeIterator() {
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
