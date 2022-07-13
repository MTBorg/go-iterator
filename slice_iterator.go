package iterator

// FromSlice returns an Iterator that iterates over a slice.
func FromSlice[T any](l []T) Iterator[T] {
	return Iterator[T]{newSliceIterator(l)}
}

type sliceIterator[T any] struct {
	data []T
	idx  int
}

func newSliceIterator[T any](l []T) *sliceIterator[T] {
	return &sliceIterator[T]{l, 0}
}

func (iter *sliceIterator[T]) Next() *T {
	if iter.idx >= len(iter.data) {
		return nil
	}

	i := iter.idx
	iter.idx += 1

	return &iter.data[i]
}
