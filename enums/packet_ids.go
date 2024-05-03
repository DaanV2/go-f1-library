package enums

type PacketId uint8

const (
	Motion              PacketId = 0  // Contains all motion data for player’s car – only sent while player is in control
	Session             PacketId = 1  // Data about the session – track, time left
	LapData             PacketId = 2  // Data about all the lap times of cars in the session
	Event               PacketId = 3  // Various notable events that happen during a session
	Participants        PacketId = 4  // List of participants in the session, mostly relevant for multiplayer
	CarSetups           PacketId = 5  // Packet detailing car setups for cars in the race
	CarTelemetry        PacketId = 6  // Telemetry data for all cars
	CarStatus           PacketId = 7  // Status data for all cars
	FinalClassification PacketId = 8  // Final classification confirmation at the end of a race
	LobbyInfo           PacketId = 9  // Information about players in a multiplayer lobby
	CarDamage           PacketId = 10 // Damage status for all cars
	SessionHistory      PacketId = 11 // Lap and tyre data for session
	TyreSets            PacketId = 12 // Extended tyre set data
	MotionEx            PacketId = 13 // Extended motion data for player car
)
