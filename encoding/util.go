package encoding

// Read4Times reads a value from a function 4 times and returns an array of the values
func Read4Times[T any](f func() T) [4]T {
	return [4]T{f(), f(), f(), f()}
}

// Read8Times reads a value from a function 8 times and returns an array of the values
func Read8Times[T any](f func() T) [8]T {
	return [8]T{f(), f(), f(), f(), f(), f(), f(), f()}
}