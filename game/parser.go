package game

import "github.com/DaanV2/go-f1-library/encoding"

type PacketParser interface {
	ParsePacket(decoder *encoding.Decoder) (any, error)
}