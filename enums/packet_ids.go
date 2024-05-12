package enums

import "github.com/DaanV2/go-f1-library/constants"

// PacketId is a flag for the packet id.
type PacketId uint8

const (
	PID_Motion              PacketId = 0  // Contains all motion data for player’s car – only sent while player is in control
	PID_Session             PacketId = 1  // Data about the session – track, time left
	PID_LapData             PacketId = 2  // Data about all the lap times of cars in the session
	PID_Event               PacketId = 3  // Various notable events that happen during a session
	PID_Participants        PacketId = 4  // List of participants in the session, mostly relevant for multiplayer
	PID_CarSetups           PacketId = 5  // Packet detailing car setups for cars in the race
	PID_CarTelemetry        PacketId = 6  // Telemetry data for all cars
	PID_CarStatus           PacketId = 7  // Status data for all cars
	PID_FinalClassification PacketId = 8  // Final classification confirmation at the end of a race
	PID_LobbyInfo           PacketId = 9  // Information about players in a multiplayer lobby
	PID_CarDamage           PacketId = 10 // Damage status for all cars
	PID_SessionHistory      PacketId = 11 // Lap and tyre data for session
	PID_TyreSets            PacketId = 12 // Extended tyre set data
	PID_MotionEx            PacketId = 13 // Extended motion data for player car
)

func (p PacketId) Id() uint8 {
	return uint8(p)
}

func (p PacketId) String() string {
	switch p {
	case PID_Motion:
		return "Motion"
	case PID_Session:
		return "Session"
	case PID_LapData:
		return "LapData"
	case PID_Event:
		return "Event"
	case PID_Participants:
		return "Participants"
	case PID_CarSetups:
		return "CarSetups"
	case PID_CarTelemetry:
		return "CarTelemetry"
	case PID_CarStatus:
		return "CarStatus"
	case PID_FinalClassification:
		return "FinalClassification"
	case PID_LobbyInfo:
		return "LobbyInfo"
	case PID_CarDamage:
		return "CarDamage"
	case PID_SessionHistory:
		return "SessionHistory"
	case PID_TyreSets:
		return "TyreSets"
	case PID_MotionEx:
		return "MotionEx"
	}

	return constants.UNKNOWN
}