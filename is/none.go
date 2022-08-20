package is

func None[T any](src []T, f func(T) bool) bool {
	return All(src, func(item T) bool {
		return !f(item)
	})
}
