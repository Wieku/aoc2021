package util

func Reverse[T any](slice []T) {
	n := len(slice)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}
