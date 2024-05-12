package enums

// PacketFormat represents the packet format which is based on the year of the game
type PacketFormat uint16

const (
	PF_F1_2022 PacketFormat = 2022 // packet format f1 2022
	PF_F1_2023 PacketFormat = 2023 // packet format f1 2023
)