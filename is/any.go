package is

func Any[T any](src []T, f func(T) bool) bool {
	for _, item := range src {
		if f(item) {
			return true
		}
	}

	return false
}
