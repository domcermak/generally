package is

import "domcermak/stl/do"

func Equal[T comparable](a, b []T) bool {
	return do.Match(a, b)
}
