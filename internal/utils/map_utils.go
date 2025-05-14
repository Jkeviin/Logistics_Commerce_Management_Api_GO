package utils

// MapToSlice converts a map into a slice of values.
func MapToSlice[K comparable, V any](m map[K]V) []V {
	var slice []V
	for _, value := range m {
		slice = append(slice, value)
	}
	return slice
}
