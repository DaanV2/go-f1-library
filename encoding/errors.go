package encoding

import "fmt"

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