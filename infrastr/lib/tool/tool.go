package tool

func Include[T comparable](i T, ss []T) int {
	for k, s := range ss {
		if i == s {
			return k
		}
	}
	return -1
}
