package extend

import (
	"github.com/domcermak/generally/do"
	"github.com/domcermak/generally/is"
)

type ExtendedSlice[T any] []T

// Slice wraps the given map type as ExtendedSlice implementing additional methods.
func Slice[T any](slice []T) ExtendedSlice[T] {
	return slice
}

// Filter filters items from the slice according to the given function f.
// Filter returns a slice of the filtered items.
// An item is added to the filtered slice if the function f returns true for the item.
func (s ExtendedSlice[T]) Filter(f func(T) bool) ExtendedSlice[T] {
	return do.Filter(s, f)
}

// Reduce reduces the slice with the given function f to a single data item.
func (s ExtendedSlice[T]) Reduce(f func(T, T) T) T {
	return do.Reduce(s, f)
}

// Map applies the given function f to each item in the slice.
// It allocates a new slice to store mapped values.
func (s ExtendedSlice[T]) Map(f func(T) T) ExtendedSlice[T] {
	return do.Map(s, f)
}

// MapWithIndex applies the given function f to each item in the slice.
// Compared to Map, the function f expects an additional argument representing the index of an item from the slice.
func (s ExtendedSlice[T]) MapWithIndex(f func(T, int) T) ExtendedSlice[T] {
	return do.MapWithIndex(s, f)
}

// Each applies the given function f to each item in the slice.
//
// Each uses the slice as the destination slice,
// thus replacing the original values in the slice with a new values according to the given function f.
// This means that no additional memory is allocated.
func (s ExtendedSlice[T]) Each(f func(T) T) ExtendedSlice[T] {
	return do.Each(s, f)
}

// EachWithIndex applies the given function f to each item in the slice.
//
// EachWithIndex uses the slice as the destination slice,
// thus replacing the original values in the slice with a new values according to the given function f.
// This means that no additional memory is allocated.
//
// Compared to Each, the f function expects an additional argument representing the index of an item from the slice.
func (s ExtendedSlice[T]) EachWithIndex(f func(T, int) T) ExtendedSlice[T] {
	return do.EachWithIndex(s, f)
}

// All returns true if all items from the slice passed through the given function f return true.
func (s ExtendedSlice[T]) All(f func(T) bool) bool {
	return is.All(s, f)
}

// Any returns true if at least one item from the slice passed through the given function f returns true.
func (s ExtendedSlice[T]) Any(f func(T) bool) bool {
	return is.Any(s, f)
}

// None returns true if no item from the slice passed through the given function f returns true.
func (s ExtendedSlice[T]) None(f func(T) bool) bool {
	return is.None(s, f)
}

// Contains returns true any item from the slice passed to the given function f returns true.
func (s ExtendedSlice[T]) Contains(f func(T) bool) bool {
	return s.Any(f)
}

// MemCopy creates a shallow copy of the slice.
func (s ExtendedSlice[T]) MemCopy() ExtendedSlice[T] {
	dst := make(ExtendedSlice[T], len(s))
	if len(dst) == 0 {
		return dst
	}

	copy(dst, s)

	return dst
}

// Copy creates a deep copy of the slice.
func (s ExtendedSlice[T]) Copy(f func(T) T) ExtendedSlice[T] {
	return s.Map(f)
}

// Raw unwraps the slice from the ExtendedSlice type.
func (s ExtendedSlice[T]) Raw() []T {
	return s
}

// Length returns the length of the slice.
func (s ExtendedSlice[T]) Length() int {
	return len(s)
}
