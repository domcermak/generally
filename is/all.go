package is

func All[T any](src []T, f func(T) bool) bool {
	return !Any(src, func(item T) bool {
		return !f(item)
	})
}
