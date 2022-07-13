package iterator

// FromMap returns an Iterator that iterates over a map.
// The order of the elements should not be assumed to be consistent for
// different but identical invocations.
func FromMap[K comparable, V any](m map[K]V) Iterator[V] {
	return Iterator[V]{newMapIterator(m)}
}

type mapIterator[K comparable, V any] struct {
	keys []K
	data map[K]V
	idx  int
}

func newMapIterator[K comparable, V any](m map[K]V) *mapIterator[K, V] {
	iter := mapIterator[K, V]{}
	iter.data = m
	iter.keys = make([]K, 0, len(m))
	for key := range m {
		iter.keys = append(iter.keys, key)
	}
	return &iter
}

func (iter *mapIterator[K, V]) Next() *V {
	if iter.idx >= len(iter.keys) {
		return nil
	}

	val := iter.data[iter.keys[iter.idx]]
	iter.idx += 1

	return &val
}
