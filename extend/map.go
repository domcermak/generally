package extend

type ExtendedMap[T comparable, E any] map[T]E

// Map wraps the given map type as ExtendedMap implementing additional methods.
func Map[T comparable, E any](m map[T]E) ExtendedMap[T, E] {
	return m
}

// Keys method returns keys from the map.
func (m ExtendedMap[T, E]) Keys() ExtendedSlice[T] {
	var keys []T
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

// Values method returns values from the map.
func (m ExtendedMap[T, E]) Values() ExtendedSlice[E] {
	var values []E
	for _, value := range m {
		values = append(values, value)
	}

	return values
}
