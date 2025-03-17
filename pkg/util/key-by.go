package util

func KeyBy[K comparable, V any](collection []V, iteratee func(item V) K) map[K]V {
	result := make(map[K]V, len(collection))

	for i := range collection {
		k := iteratee(collection[i])
		result[k] = collection[i]
	}

	return result
}
