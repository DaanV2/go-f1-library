package f1_2023

import (
	"github.com/DaanV2/go-f1-library/encoding"
	"github.com/DaanV2/go-f1-library/enums"
)

type PacketHeader struct {
	PacketFormat            enums.PacketFormat // uint16, 2023
	GameYear                uint8              // Game year - last two digits e.g. 23
	GameMajorVersion        uint8              // Game major version - "X.00"
	GameMinorVersion        uint8              // Game minor version - "1.XX"
	PacketVersion           uint8              // Version of this packet type, all start from 1
	PacketId                uint8              // Identifier for the packet type, see below
	SessionUID              uint64             // Unique identifier for the session
	SessionTime             float32            // Session timestamp
	FrameIdentifier         uint32             // Identifier for the frame the data was retrieved on
	OverallFrameIdentifier  uint32             // Overall identifier for the frame the data was retrieved on, doesn't go back after flashbacks
	PlayerCarIndex          uint8              // Index of player's car in the array
	SecondaryPlayerCarIndex uint8              // Index of secondary player's car in the array (splitscreen), 255 if no second player
}

// PacketFormat returns the format of the packet
func (p *PacketHeader) PlayerIndex() uint8 {
	return p.PlayerCarIndex
}

// SecondaryPlayerIndex returns the index of the secondary player's car in the array (splitscreen) and a bool indicating if there is a secondary player
func (p *PacketHeader) SecondaryPlayerIndex() (uint8, bool) {
	return p.SecondaryPlayerCarIndex, p.HasSecondaryPlayer()
}

func (p *PacketHeader) HasSecondaryPlayer() bool {
	return p.SecondaryPlayerCarIndex != 255
}

// DeserializePacketHeader will deserialise a byte slice into a PacketHeader, expects the slice to be 29 bytes long
func DeserializePacketHeader(buf []byte) PacketHeader {
	_ = buf[28] // Ensure buf is large enough

	return PacketHeader{
		PacketFormat:            enums.PacketFormat(encoding.Uint16(buf[0:2])),
		GameYear:                encoding.Uint8(buf[2]),
		GameMajorVersion:        encoding.Uint8(buf[3]),
		GameMinorVersion:        encoding.Uint8(buf[4]),
		PacketVersion:           encoding.Uint8(buf[5]),
		PacketId:                encoding.Uint8(buf[6]),
		SessionUID:              encoding.Uint64(buf[7:15]),
		SessionTime:             float32(encoding.Uint32(buf[15:19])),
		FrameIdentifier:         encoding.Uint32(buf[19:23]),
		OverallFrameIdentifier:  encoding.Uint32(buf[23:27]),
		PlayerCarIndex:          buf[27],
		SecondaryPlayerCarIndex: buf[28],
	}
}
