package encoding

import "errors"

var (
	// ErrBufferNotLargeEnough is returned when the buffer is not large enough to read the requested data
	ErrBufferNotLargeEnough = errors.New("buffer not large enough")
)