package encoding

type Decoder struct {
	buffer []byte
	index  int
}

func NewDecoder(buffer []byte) *Decoder {
	return &Decoder{
		buffer: buffer,
		index:  0,
	}
}

// Len returns the length of the buffer
func (d *Decoder) Len() int {
	return len(d.buffer)
}

// Index returns the current index of the decoder
func (d *Decoder) Index() int {
	return d.index
}

// Seek sets the index of the decoder
func (d *Decoder) Seek(offset int) {
	d.index = offset
}

// Skip moves the index of the decoder by the offset
func (d *Decoder) Skip(offset int) {
	d.index += offset
}

// LeftToRead returns the amount of bytes left to read
func (d *Decoder) LeftToRead() int {
	return d.Len() - d.Index()
}

func (d *Decoder) IsLargeEnough(size int) error {
	if d.Len() < size {
		return BufferNotLargeEnoughError{ExpectedSize: size, ActualSize: d.LeftToRead()}
	}

	return nil
}

func (d *Decoder) HasEnoughLeft(still_to_read int) error {
	if d.LeftToRead() < still_to_read {
		return BufferNotLargeEnoughError{ExpectedSize: still_to_read, ActualSize: d.LeftToRead()}
	}

	return nil
}

// Read reads the buffer into the given byte slice
func (d *Decoder) Read(p []byte) int {
	n := copy(p, d.buffer[d.index:])
	d.index += n
	return n
}

func (d *Decoder) Read48() [48]byte {
	var v [48]byte
	d.Read(v[:])
	return v
}

// Uint8 reads a byte from the buffer, and moves the index by 1
func (d *Decoder) Byte() byte {
	v := d.buffer[d.index]
	d.index++
	return v
}

// Uint8 reads a uint8 from the buffer, and moves the index by 1
func (d *Decoder) Uint8() uint8 {
	v := Uint8(d.buffer[d.index])
	d.index++
	return v
}

// Uint16 reads a uint16 from the buffer, and moves the index by 2
func (d *Decoder) Uint16() uint16 {
	v := Uint16(d.buffer[d.index:])
	d.index += 2
	return v
}

// Uint32 reads a uint32 from the buffer, and moves the index by 4
func (d *Decoder) Uint32() uint32 {
	v := Uint32(d.buffer[d.index:])
	d.index += 4
	return v
}

// Uint64 reads a uint64 from the buffer, and moves the index by 8
func (d *Decoder) Uint64() uint64 {
	v := Uint64(d.buffer[d.index:])
	d.index += 8
	return v
}

// Int8 reads a uint8 from the buffer, and moves the index by 1
func (d *Decoder) Int8() int8 {
	v := Int8(d.buffer[d.index])
	d.index++
	return v
}

// Int16 reads a uint16 from the buffer, and moves the index by 2
func (d *Decoder) Int16() int16 {
	v := Int16(d.buffer[d.index:])
	d.index += 2
	return v
}

// Int32 reads a uint32 from the buffer, and moves the index by 4
func (d *Decoder) Int32() int32 {
	v := Int32(d.buffer[d.index:])
	d.index += 4
	return v
}

// Int64 reads a uint64 from the buffer, and moves the index by 8
func (d *Decoder) Int64() int64 {
	v := Int64(d.buffer[d.index:])
	d.index += 8
	return v
}

// Float32 reads a float32 from the buffer, and moves the index by 4
func (d *Decoder) Float32() float32 {
	v := Float32(d.buffer[d.index:])
	d.index += 4
	return v
}

// Float64 reads a float64 from the buffer, and moves the index by 8
func (d *Decoder) Float64() float64 {
	v := Float64(d.buffer[d.index:])
	d.index += 8
	return v
}
