package c

import "slices"

// String turns a null-terminated byte array into a string
func String(data []byte) string {
	index := slices.Index(data, 0)
	if index == -1 {
		return string(data)
	}
	return string(data[:index])
}