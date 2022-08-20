package do

// Each function is a higher order function applying a given function f to each item in a given slice.
//
// Each function uses the given slice as the destination slice,
// thus changing replacing the original values in the slice with a new values according to the given function f.
// This means that no additional memory is allocated.
func Each[T any](slice []T, f func(T) T) []T {
	async(slice, slice, f)

	return slice
}

// EachWithIndex function is a higher order function applying a given function f to each item in a given slice.
//
// EachWithIndex function uses the given slice as the destination slice,
// thus replacing the original values in the slice with a new values according to the given function f.
// This means that no additional memory is allocated.
//
// Compared to Each function, the f function expects an additional argument representing the index of an item from the given slice.
func EachWithIndex[T any](slice []T, f func(T, int) T) []T {
	asyncWithIndex(slice, slice, f)

	return slice
}
