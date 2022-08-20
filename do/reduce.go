package do

// Reduce reduces given slice with the given function f to a single data item.
func Reduce[T any, E any](src []T, f func(E, T) E) E {
	var acc E

	for _, item := range src {
		acc = f(acc, item)
	}

	return acc
}
