package do

// Filter filters items from a given slice according to the given function f.
// Filter returns a slice of the filtered items.
// An item is added to the filtered slice if the function f returns true for the item.
func Filter[T any](src []T, f func(T) bool) []T {
	dst := make([]T, 0)
	for _, item := range src {
		if f(item) {
			dst = append(dst, item)
		}
	}

	return dst
}
