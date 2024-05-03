package f1_2023

type PacketHeader struct {
	// 2023
	PacketFormat uint16 `json:"packet_format"`
	// Game year - last two digits e.g. 23
	GameYear uint8 `json:"game_year"`
	// Game major version - "X.00"
	GameMajorVersion uint8 `json:"game_major_version"`
	// Game minor version - "1.XX"
	GameMinorVersion uint8 `json:"game_minor_version"`
	// Version of this packet type, all start from 1
	PacketVersion uint8 `json:"packet_version"`
	// Identifier for the packet type, see below
	PacketId uint8 `json:"packet_id"`
	// Unique identifier for the session
	SessionUID uint64 `json:"session_uid"`
	// Session timestamp
	SessionTime float32 `json:"session_time"`
	// Identifier for the frame the data was retrieved on
	FrameIdentifier uint32 `json:"frame_identifier"`
	// Overall identifier for the frame the data was retrieved on, doesn't go back after flashbacks
	OverallFrameIdentifier uint32 `json:"overall_frame_identifier"`
	// Index of player's car in the array
	PlayerCarIndex uint8 `json:"player_car_index"`
	// Index of secondary player's car in the array (splitscreen), 255 if no second player
	SecondaryPlayerCarIndex uint8 `json:"secondary_player_car_index"`
}

func (p *PacketHeader) HasSecondaryPlayer() bool {
	return p.SecondaryPlayerCarIndex != 255
}