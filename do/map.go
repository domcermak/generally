package do

// Map function is a higher order function applying a given function f to each item in a given slice.
//
// It allocates a new slice to store mapped values.
func Map[T any, E any](slice []T, f func(T) E) []E {
	dst := make([]E, len(slice))
	async(slice, dst, f)

	return dst
}

// MapWithIndex function is a higher order function applying a given function f to each item in a given slice.
//
// Compared to Map function, the f function expects an additional argument representing the index of an item from the given slice.
func MapWithIndex[T any, E any](slice []T, f func(T, int) E) []E {
	dst := make([]E, len(slice))
	asyncWithIndex(slice, dst, f)

	return dst
}
