package settype

import (
	"domcermak/stl/extend"
)

type Set[T comparable] map[T]bool

// Empty creates an empty Set instance.
func Empty[T comparable]() Set[T] {
	return make(map[T]bool)
}

// FromSlice creates a Set from slice, thus removing duplicated items.
func FromSlice[T comparable](slice []T) Set[T] {
	set := Empty[T]()
	for _, item := range slice {
		set.Add(item)
	}

	return set
}

// Add adds an item to the set.
func (s Set[T]) Add(item T) {
	s[item] = true
}

// Contains returns true if the given item matches any item in the Set.
func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

// Delete deletes an item from the set and returns true.
// Delete returns false if the given item is not present in the set.
func (s Set[T]) Delete(item T) bool {
	if s.Contains(item) {
		delete(s, item)
		return true
	}

	return false
}

// Intersect returns intersection of the given set with the set.
func (s Set[T]) Intersect(other Set[T]) Set[T] {
	intersect := Empty[T]()
	for item := range s {
		if other.Contains(item) {
			intersect.Add(item)
		}
	}
	for item := range other {
		if s.Contains(item) {
			intersect.Add(item)
		}
	}

	return intersect
}

// Union returns union of the given set and the set.
func (s Set[T]) Union(other Set[T]) Set[T] {
	union := other.Copy()
	for item := range other {
		union.Add(item)
	}

	return union
}

// Minus subtracts the given set from the set.
func (s Set[T]) Minus(other Set[T]) Set[T] {
	union := s.Copy()
	for item := range other {
		union.Delete(item)
	}

	return union
}

// Items returns ExtendedSlice of items present in the set.
func (s Set[T]) Items() extend.ExtendedSlice[T] {
	items := make([]T, len(s))
	i := 0
	for item := range s {
		items[i] = item
		i++
	}

	return extend.Slice(items)
}

// Copy creates a new instance of Set and copies the items from the set to the new set instance.
// The new Set instance is returned then.
func (s Set[T]) Copy() Set[T] {
	return FromSlice(s.Items())
}

// Map copies the set items, applies the function f to each item
// and returns the new set with the mapped values.
func (s Set[T]) Map(f func(T) T) Set[T] {
	return FromSlice(extend.Slice(s.Items()).Map(f))
}

// Filter filters items from the set according to the given function f.
// Filter returns a set of the filtered items.
// An item is added to the filtered set if the function f returns true for the item.
func (s Set[T]) Filter(f func(T) bool) Set[T] {
	return FromSlice(extend.Slice(s.Items()).Filter(f))
}

// Reduce reduces the set items with the given function f to a single data item.
func (s Set[T]) Reduce(f func(T, T) T) T {
	return extend.Slice(s.Items()).Reduce(f)
}

// All returns true if all items from the set passed through the given function f return true.
func (s Set[T]) All(f func(T) bool) bool {
	return extend.Slice(s.Items()).All(f)
}

// Any returns true if at least one item from the set passed through the given function f returns true.
func (s Set[T]) Any(f func(T) bool) bool {
	return extend.Slice(s.Items()).Any(f)
}

// None returns true if no item from the set passed through the given function f returns true.
func (s Set[T]) None(f func(T) bool) bool {
	return extend.Slice(s.Items()).None(f)
}

// Equal return true if the given set contains the same items as the set.
func (s Set[T]) Equal(other Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}

	return s.All(func(item T) bool {
		return other.Contains(item)
	})
}

// Size returns the number of items in the set.
func (s Set[T]) Size() int {
	return s.Items().Length()
}
