package f1_2023

import (
	"github.com/DaanV2/go-f1-library/encoding"
	"github.com/DaanV2/go-f1-library/enums"
)

const (
	PacketHeaderSize = 29
)

type PacketHeader struct {
	PacketFormat            enums.PacketFormat `json:"packet_format"`              // uint16, 2023
	GameYear                uint8              `json:"game_year"`                  // Game year - last two digits e.g. 23
	GameMajorVersion        uint8              `json:"game_major_version"`         // Game major version - "X.00"
	GameMinorVersion        uint8              `json:"game_minor_version"`         // Game minor version - "1.XX"
	PacketVersion           uint8              `json:"packet_version"`             // Version of this packet type, all start from 1
	PacketId                enums.PacketId     `json:"packet_id"`                  // uint8, Identifier for the packet type, see below
	SessionUID              uint64             `json:"session_uid"`                // Unique identifier for the session
	SessionTime             float32            `json:"session_time"`               // Session timestamp
	FrameIdentifier         uint32             `json:"frame_identifier"`           // Identifier for the frame the data was retrieved on
	OverallFrameIdentifier  uint32             `json:"overall_frame_identifier"`   // Overall identifier for the frame the data was retrieved on, doesn't go back after flashbacks
	PlayerCarIndex          uint8              `json:"player_car_index"`           // Index of player's car in the array
	SecondaryPlayerCarIndex uint8              `json:"secondary_player_car_index"` // Index of secondary player's car in the array (splitscreen), 255 if no second player
}

// Size returns the size of the packet
func (p PacketHeader) Size() int {
	return PacketHeaderSize
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

// ParsePacketHeader will deserialise a byte slice into a PacketHeader, expects the slice to be 29 bytes long
func ParsePacketHeader(decoder *encoding.Decoder) (PacketHeader, error) {
	if decoder.LeftToRead() < PacketHeaderSize {
		return PacketHeader{}, encoding.NewBufferNotLargeEnoughError(PacketCarDamageDataSize, decoder.LeftToRead())
	}

	return PacketHeader{
		PacketFormat:            enums.PacketFormat(decoder.Uint16()),
		GameYear:                decoder.Uint8(),
		GameMajorVersion:        decoder.Uint8(),
		GameMinorVersion:        decoder.Uint8(),
		PacketVersion:           decoder.Uint8(),
		PacketId:                enums.PacketId(decoder.Uint8()),
		SessionUID:              decoder.Uint64(),
		SessionTime:             float32(decoder.Uint32()),
		FrameIdentifier:         decoder.Uint32(),
		OverallFrameIdentifier:  decoder.Uint32(),
		PlayerCarIndex:          decoder.Uint8(),
		SecondaryPlayerCarIndex: decoder.Uint8(),
	}, nil
}
