package pkg

func Contains[T comparable](slice []T, expected T) bool {
	for _, s := range slice {
		if s == expected {
			return true
		}
	}
	return false
}