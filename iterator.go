package iterator

// An Iterable implements how to iterate over something one element at a time.
type Iterable[T any] interface {
	Next() *T
}

// Iterator iterates over an iterable object.
// Iterator is not safe for concurrent use.
// All methods consume the iterator.
type Iterator[T any] struct {
	iterable Iterable[T]
}

func (iter Iterator[T]) Next() *T {
	return iter.iterable.Next()
}

// Iterate creates an Iterator for an iterable object
func Iterate[T any](iterable Iterable[T]) Iterator[T] {
	return Iterator[T]{iterable}
}

// Map creates an iterator by mapping the elements of another iterator.
func Map[T any, U any](iterator Iterator[T], fn func(T) U) Iterator[U] {
	var res []U
	for {
		next := iterator.Next()
		if next == nil {
			break
		}
		res = append(res, fn(*next))
	}
	return FromSlice(res)
}

// Collect collects the iterator's elements into a slice []T.
func (iter Iterator[T]) Collect() []T {
	var res []T
	for {
		next := iter.Next()
		if next == nil {
			return res
		}
		res = append(res, *next)
	}
}

// Filter creates an Iterator[T] that filters iter's elements using a boolean
// predicate.
func (iter Iterator[T]) Filter(fn func(T) bool) Iterator[T] {
	var res []T
	for {
		next := iter.Next()
		if next == nil {
			break
		}
		if fn(*next) {
			res = append(res, *next)
		}
	}
	return FromSlice(res)
}

// Count returns the count of the iterator's elements.
func (iter Iterator[T]) Count() int {
	count := 0
	for {
		next := iter.Next()
		if next == nil {
			return count
		}
		count += 1
	}
}

// Take creates a new iterator over iter's n first elements.
func (iter Iterator[T]) Take(n int) Iterator[T] {
	var res []T
	taken := 0
	for {
		if taken >= n {
			break
		}
		next := iter.Next()
		if next == nil {
			break
		}

		taken += 1
		res = append(res, *next)
	}
	return FromSlice(res)
}

// Nth returns the nth element of the iterator (or nil if n exceeds the count).
func (iter Iterator[T]) Nth(n int) *T {
	i := 0
	for {
		next := iter.Next()
		if next == nil {
			return nil
		}

		if i == n {
			return next
		}
		i += 1
	}
}

// Last returns the last element of the iterator (or nil if no elements
// remain).
func (iter Iterator[T]) Last() *T {
	var current *T
	for {
		next := iter.Next()
		if next == nil {
			return current
		}
		current = next
	}
}

// ForEach calls a function for each of the iterator's elements.
func (iter Iterator[T]) ForEach(fn func(T)) {
	for {
		next := iter.Next()
		if next == nil {
			break
		}
		fn(*next)
	}
}

// Skip creates an iterator that skips the n first elements of iter.
func (iter Iterator[T]) Skip(n int) Iterator[T] {
	iter.Take(n)
	return iter
}

// StepBy creates a new iterator that iterates over iter's elements, skipping n
// elements each time.
func (iter Iterator[T]) StepBy(n int) Iterator[T] {
	var res []T
	for {
		next := iter.Next()
		if next == nil {
			break
		}
		res = append(res, *next)
		iter.Skip(n - 1)
	}
	return FromSlice(res)
}

// Partition partitions the iterator's elements into two sets using a boolean
// predicate (true, false respectively) and returns two iterators that iterates
// over the respective sets.
func (iter Iterator[T]) Partition(fn func(T) bool) (Iterator[T], Iterator[T]) {
	var l1, l2 []T
	for {
		next := iter.Next()
		if next == nil {
			break
		}
		if fn(*next) {
			l1 = append(l1, *next)
		} else {
			l2 = append(l2, *next)
		}
	}
	return FromSlice(l1), FromSlice(l2)
}

// Reverse creates an iterator that iterates of iter's elements in reverse
// order.
func (iter Iterator[T]) Reverse() Iterator[T] {
	vals := iter.Collect()
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}
	return FromSlice(vals)
}

// Chain concatenates the elements of two iterators and creates a new iterator
// iterating over the concatenation.
func (iter Iterator[T]) Chain(iter2 Iterator[T]) Iterator[T] {
	return FromSlice(append(iter.Collect(), iter2.Collect()...))
}
