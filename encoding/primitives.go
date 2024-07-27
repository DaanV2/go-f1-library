package encoding

import (
	"encoding/binary"
	"math"
)

// Uint8 converts the given byte into a uint8, as formatted by the spec
func Uint8(buf byte) uint8 {
	return uint8(buf)
}

// Uint16 converts the given bytes into a uint16, as formatted by the spec
func Uint16(buf []byte) uint16 {
	return binary.LittleEndian.Uint16(buf)
}

// Uint32 converts the given bytes into a uint32, as formatted by the spec
func Uint32(buf []byte) uint32 {
	return binary.LittleEndian.Uint32(buf)
}

// Uint64 converts the given bytes into a uint64, as formatted by the spec
func Uint64(buf []byte) uint64 {
	return binary.LittleEndian.Uint64(buf)
}

// Int8 converts the given byte into a int8, as formatted by the spec
func Int8(buf byte) int8 {
	return int8(buf)
}

// Int16 converts the given bytes into a int16, as formatted by the spec
func Int16(buf []byte) int16 {
	return int16(binary.LittleEndian.Uint16(buf))
}

// Int32 converts the given bytes into a int32, as formatted by the spec
func Int32(buf []byte) int32 {
	return int32(binary.LittleEndian.Uint32(buf))
}

// Int64 converts the given bytes into a int64, as formatted by the spec
func Int64(buf []byte) int64 {
	return int64(binary.LittleEndian.Uint64(buf))
}

// Float32 converts the given bytes into a float32, as formatted by the spec
func Float32(buf []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(buf))
}

// Float64 converts the given bytes into a float64, as formatted by the spec
func Float64(buf []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(buf))
}
