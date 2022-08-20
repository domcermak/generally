package do

// Each function is a higher order function applying a given function f to each item in a given slice.
func Each[T any](slice []T, f func(T)) {
	for _, item := range slice {
		f(item)
	}
}

// EachWithIndex function is a higher order function applying a given function f to each item in a given slice.
// Compared to Each function, the f function expects an additional argument representing the index of an item from the given slice.
func EachWithIndex[T any](slice []T, f func(T, int)) {
	for i, item := range slice {
		f(item, i)
	}
}
