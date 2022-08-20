package do

// Match returns true if the given slices are equal.
func Match[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, itemA := range a {
		if itemA != b[i] {
			return false
		}
	}

	return true
}

// MatchWithFunc returns true if the given slices are equal
// when items are compared with the given function f.
func MatchWithFunc[T any](a, b []T, f func(T, T) bool) bool {
	if len(a) != len(b) {
		return false
	}

	for i, itemA := range a {
		if f(itemA, b[i]) {
			return false
		}
	}

	return true
}
