package general

import (
	"github.com/DaanV2/go-f1-library/encoding"
	"github.com/DaanV2/go-f1-library/enums"
)

// PacketHeader is a subset of the data sent in every packet helping to identify the packet for format, version and game year
type PacketHeader struct {
	PacketFormat     enums.PacketFormat `json:"packet_format"`      // uint16, 2023
	GameYear         uint8              `json:"game_year"`          // Game year - last two digits e.g. 23
	GameMajorVersion uint8              `json:"game_major_version"` // Game major version - "X.00"
	GameMinorVersion uint8              `json:"game_minor_version"` // Game minor version - "1.XX"
	PacketVersion    uint8              `json:"packet_version"`     // Version of this packet type, all start from 1
	PacketId         enums.PacketId     `json:"packet_id"`          // uint8, Identifier for the packet type, see below
}

// ParsePacketHeader will deserialise a byte slice into a PacketHeader, expects the slice to be 29 bytes long
func ParsePacketHeader(data []byte) PacketHeader {
	return PacketHeader{
		PacketFormat:            enums.PacketFormat(encoding.Uint16(data[0:2])),
		GameYear:                encoding.Uint8(data[2]),
		GameMajorVersion:        encoding.Uint8(data[3]),
		GameMinorVersion:        encoding.Uint8(data[4]),
		PacketVersion:           encoding.Uint8(data[5]),
		PacketId:                enums.PacketId(encoding.Uint8(data[6])),
	}
}
