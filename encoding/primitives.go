package encoding

import (
	"encoding/binary"
	"math"
)

func Uint8(buf byte) uint8 {
	return uint8(buf)
}

func Uint16(buf []byte) uint16 {
	return binary.LittleEndian.Uint16(buf)
}

func Uint32(buf []byte) uint32 {
	return binary.LittleEndian.Uint32(buf)
}

func Uint64(buf []byte) uint64 {
	return binary.LittleEndian.Uint64(buf)
}

func Float32(buf []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(buf))
}

func Float64(buf []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(buf))
}
