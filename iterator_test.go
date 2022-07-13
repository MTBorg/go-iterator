package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockIterator[T any] struct {
	mock.Mock
}

func (iter *mockIterator[T]) Next() *T {
	args := iter.Mock.Called()
	return args.Get(0).(*T)
}

func TestCollect(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return((*int)(nil)).Once()
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	assert.Equal(t, []int{1, 2, 3}, iter.Collect())
}

func TestFilter(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return(ptr(4)).Once()
	m.On("Next").Return((*int)(nil)).Once()
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	assert.Equal(t, []int{2, 4}, iter.Filter(func(i int) bool { return i%2 == 0 }).Collect())
}

func TestCount(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return((*int)(nil)).Once()
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	assert.Equal(t, 3, iter.Count())
}

func TestMap(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return((*int)(nil)).Once()
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	double := func(n int) float64 { return float64(n) * 2 }
	assert.Equal(t, []float64{2.0, 4.0, 6.0}, Map(iter, double).Collect())
}

func TestTake(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return(ptr(4)).Once()
	m.On("Next").Return((*int)(nil)).Once()
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	assert.Equal(t, []int{1, 2}, iter.Take(2).Collect())
	assert.Equal(t, []int{3}, iter.Take(1).Collect())
	assert.Equal(t, []int{4}, iter.Take(2).Collect())
}

func TestSliceIterator(t *testing.T) {
	iter := FromSlice([]int{1, 2, 3})
	assert.Equal(t, []int{1, 2, 3}, iter.Collect())
}

func TestMapIterator(t *testing.T) {
	iter := FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	assert.ElementsMatch(t, []int{1, 2, 3}, iter.Collect())
}

func TestNth(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return(ptr(4)).Once()
	m.On("Next").Return((*int)(nil)).Once()
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	assert.Equal(t, 1, value(iter.Nth(0)))
	assert.Equal(t, 2, value(iter.Nth(0)))
	assert.Equal(t, 4, value(iter.Nth(1)))
	assert.Equal(t, (*int)(nil), iter.Nth(0))
}

func TestLast(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return(ptr(4)).Once()
	m.On("Next").Return((*int)(nil))
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	assert.Equal(t, 4, value(iter.Last()))
}

func TestForEach(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return(ptr(4)).Once()
	m.On("Next").Return((*int)(nil))
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	res := 0
	iter.ForEach(func(i int) { res += i })

	assert.Equal(t, 10, res)
}

func TestSkip(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return(ptr(4)).Once()
	m.On("Next").Return((*int)(nil))
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	iter.Skip(2)

	assert.Equal(t, []int{3, 4}, iter.Collect())
}

func TestStepBy(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return(ptr(4)).Once()
	m.On("Next").Return(ptr(5)).Once()
	m.On("Next").Return(ptr(6)).Once()
	m.On("Next").Return((*int)(nil))
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	assert.Equal(t, []int{1, 3, 5}, iter.StepBy(2).Collect())
}

func TestPartition(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return(ptr(4)).Once()
	m.On("Next").Return((*int)(nil))
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	even, odd := iter.Partition(func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2, 4}, even.Collect())
	assert.Equal(t, []int{1, 3}, odd.Collect())
}

func TestReverse(t *testing.T) {
	m := &mockIterator[int]{}
	m.On("Next").Return(ptr(1)).Once()
	m.On("Next").Return(ptr(2)).Once()
	m.On("Next").Return(ptr(3)).Once()
	m.On("Next").Return(ptr(4)).Once()
	m.On("Next").Return((*int)(nil))
	defer m.AssertExpectations(t)
	iter := Iterate[int](m)

	assert.Equal(t, []int{4, 3, 2, 1}, iter.Reverse().Collect())
}

func TestChain(t *testing.T) {
	m1 := &mockIterator[int]{}
	m1.On("Next").Return(ptr(1)).Once()
	m1.On("Next").Return(ptr(2)).Once()
	m1.On("Next").Return((*int)(nil))
	defer m1.AssertExpectations(t)
	iter1 := Iterate[int](m1)

	m2 := &mockIterator[int]{}
	m2.On("Next").Return(ptr(3)).Once()
	m2.On("Next").Return(ptr(4)).Once()
	m2.On("Next").Return((*int)(nil))
	defer m2.AssertExpectations(t)
	iter2 := Iterate[int](m2)

	assert.Equal(t, []int{1, 2, 3, 4}, iter1.Chain(iter2).Collect())
}

func ptr[T any](t T) *T {
	return &t
}

func value[T any](t *T) T {
	if t == nil {
		var res T
		return res
	}
	return *t
}

func ExampleIterator_Chain() {
	i1 := FromSlice([]int{1, 2})
	i2 := FromSlice([]int{3, 4})
	l := i1.Chain(i2).Collect()
	fmt.Println(l)
	// Output: [1 2 3 4]
}

func ExampleIterator_Filter() {
	even := FromSlice([]int{1, 2, 3, 4}).
		Filter(func(i int) bool { return i%2 == 0 }).
		Collect()
	fmt.Println(even)
	// Output: [2 4]
}

func ExampleIterator_ForEach() {
	FromSlice([]int{1, 2, 3}).ForEach(func(i int) { fmt.Println(i) })
	// Output: 1
	// 2
	// 3
}

func ExampleIterator_Last() {
	last := FromSlice([]int{1, 2, 3}).Last()
	fmt.Println(*last)
	// Output: 3
}

func ExampleIterator_Nth() {
	nth := FromSlice([]int{1, 2, 3, 4}).Nth(2)
	fmt.Println(*nth)
	// Output: 3
}

func ExampleIterator_Count() {
	count := FromSlice([]int{1, 2, 3}).Count()
	fmt.Println(count)
	// Output: 3
}

func ExampleIterator_Partition() {
	even, odd := FromSlice([]int{1, 2, 3, 4}).
		Partition(func(i int) bool { return i%2 == 0 })
	fmt.Println(even.Collect(), odd.Collect())
	// Output: [2 4] [1 3]
}

func ExampleIterator_Reverse() {
	reversed := FromSlice([]int{1, 2, 3}).Reverse().Collect()
	fmt.Println(reversed)
	// Output: [3 2 1]
}

func ExampleIterator_Skip() {
	l := FromSlice([]int{1, 2, 3, 4}).Skip(2).Collect()
	fmt.Println(l)
	// Output: [3 4]
}

func ExampleIterator_StepBy() {
	l := FromSlice([]int{1, 2, 3, 4}).StepBy(2).Collect()
	fmt.Println(l)
	// Output: [1 3]
}

func ExampleIterator_Take() {
	l := FromSlice([]int{1, 2, 3, 4}).Take(2).Collect()
	fmt.Println(l)
	// Output: [1 2]
}

func ExampleMap() {
	iter := Map(FromSlice([]int{1, 2, 3}), func(i int) float64 { return float64(i) * 2 })
	fmt.Println(iter.Collect())
	// Output: [2 4 6]
}

func ExampleFromMap() {
	iter := FromMap(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	iter.ForEach(func(i int) { fmt.Println(i) })
	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleFromSlice() {
	iter := FromSlice([]int{1, 2, 3})
	iter.ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 1
	// 2
	// 3
}
