package f1_2023

import "github.com/DaanV2/go-f1-library/enums"

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

func (p *PacketHeader) HasSecondaryPlayer() bool {
	return p.SecondaryPlayerCarIndex != 255
}
