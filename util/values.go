package util

// Numbers is a generic interface for a number type
type Numbers interface {
	uint8 | uint16 | uint32 | uint64
}

// Collection is a generic interface for a collection of items
type Collection[T any, U Numbers] interface {
	// Get returns the item with the given id
	Get(id U) T
	// Max returns the maximum id in the collection
	Max() U
}

// Values returns a slice of all values in the Collection
func Values[T any, U Numbers](data Collection[T, U]) []T {
	max := data.Max()

	values := make([]T, 0, max)
	for i := U(0); i < max; i++ {
		values = append(values, data.Get(i))
	}
	return values
}