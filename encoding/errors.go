package encoding

import "fmt"

// BufferNotLargeEnoughError is used to display, and tell the user about
// The buffer that isn't large enough, and how much it was expecting.
type BufferNotLargeEnoughError struct {
	ExpectedSize int
	ActualSize   int
}

func (e BufferNotLargeEnoughError) Error() string {
	return fmt.Sprintf("buffer not large enough, expected %d, got %d", e.ExpectedSize, e.ActualSize)
}

func NewBufferNotLargeEnoughError(expectedSize, actualSize int) BufferNotLargeEnoughError {
	return BufferNotLargeEnoughError{
		ExpectedSize: expectedSize,
		ActualSize:   actualSize,
	}
}