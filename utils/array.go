package utils

func LastIndex[T any](arr []T) int {
	return len(arr) - 1
}

func Last[T any](arr []T) *T {
	return &arr[LastIndex[T](arr)]
}
