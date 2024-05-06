package f1_2023

import (
	"github.com/DaanV2/go-f1-library/encoding"
	"github.com/DaanV2/go-f1-library/enums"
)

const (
	PacketEventDataSize    = 45
	PacketEventDataVersion = 1
)

type EventCode string

const (
	EC_ButtonStatus       EventCode = "BUTN" // Button status changed
	EC_ChequeredFlag      EventCode = "CHQF" // The chequered flag has been waved
	EC_DriveThroughServed EventCode = "DTSV" // Drive through penalty served
	EC_DRSdisabled        EventCode = "DRSD" // Race control have disabled DRS
	EC_DRSenabled         EventCode = "DRSE" // Race control have enabled DRS
	EC_FastestLap         EventCode = "FTLP" // When a driver achieves the fastest lap
	EC_Flashback          EventCode = "FLBK" // Flashback activated
	EC_LightsOut          EventCode = "LGOT" // Lights out
	EC_Overtake           EventCode = "OVTK" // Overtake occurred
	EC_PenaltyIssued      EventCode = "PENA" // A penalty has been issued – details in event
	EC_RaceWinner         EventCode = "RCWN" // The race winner is announced
	EC_RedFlag            EventCode = "RDFL" // Red flag shown
	EC_Retirement         EventCode = "RTMT" // When a driver retires
	EC_SessionEnded       EventCode = "SEND" // Sent when the session ends
	EC_SessionStarted     EventCode = "SSTA" // Sent when the session starts
	EC_SpeedTrapTriggered EventCode = "SPTP" // Speed trap has been triggered by fastest speed
	EC_StartLights        EventCode = "STLG" // Start lights – number shown
	EC_StopGoServed       EventCode = "SGSV" // Stop go penalty served
	EC_TeamMateInPits     EventCode = "TMPT" // Your team mate has entered the pits
)

type (
	PacketEventData struct {
		Header          PacketHeader // Header
		EventStringCode EventCode    // [4]uint8 Event string code, see below
		EventDetails    interface{}  // Event details - should be interpreted differently for each type, is a pointer to a struct of type: FastestLap, Retirement, TeamMateInPits, RaceWinner, Penalty, SpeedTrap, StartLights, DriveThroughPenaltyServed, StopGoPenaltyServed, Flashback, Buttons, Overtake
	}

	Retirement struct {
		VehicleIdx uint8 // Vehicle index of car retiring
	}

	TeamMateInPits struct {
		VehicleIdx uint8 // Vehicle index of team mate
	}

	RaceWinner struct {
		VehicleIdx uint8 // Vehicle index of the race winner
	}

	Penalty struct {
		PenaltyType      uint8 // Penalty type – see Appendices
		InfringementType uint8 // Infringement type – see Appendices
		VehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
		OtherVehicleIdx  uint8 // Vehicle index of the other car involved
		Time             uint8 // Time gained, or time spent doing action in seconds
		LapNum           uint8 // Lap the penalty occurred on
		PlacesGained     uint8 // Number of places gained by this
	}

	SpeedTrap struct {
		VehicleIdx                 uint8   // Vehicle index of the vehicle triggering speed trap
		Speed                      float32 // Top speed achieved in kilometres per hour
		IsOverallFastestInSession  uint8   // Overall fastest speed in session = 1, otherwise 0
		IsDriverFastestInSession   uint8   // Fastest speed for driver in session = 1, otherwise 0
		FastestVehicleIdxInSession uint8   // Vehicle index of the vehicle that is the fastest in this session
		FastestSpeedInSession      float32 // Speed of the vehicle that is the fastest in this session
	}

	StartLights struct {
		NumLights uint8 // Number of lights showing
	}

	DriveThroughPenaltyServed struct {
		VehicleIdx uint8 // Vehicle index of the vehicle serving drive through
	}

	StopGoPenaltyServed struct {
		VehicleIdx uint8 // Vehicle index of the vehicle serving stop go
	}

	Flashback struct {
		FlashbackFrameIdentifier uint32  // Frame identifier flashed back to
		FlashbackSessionTime     float32 // Session time flashed back to
	}

	Buttons struct {
		ButtonStatus enums.Button // uint32, Bit flags specifying which buttons are being pressed currently - see appendices
	}

	Overtake struct {
		OvertakingVehicleIdx     uint8 // Vehicle index of the vehicle overtaking
		BeingOvertakenVehicleIdx uint8 // Vehicle index of the vehicle being overtaken
	}

	FastestLap struct {
		VehicleIdx uint8   // Vehicle index of car achieving fastest lap
		LapTime    float32 // Lap time is in seconds
	}
)

// ParsePacketEventData will parse the given data into a packet
func ParsePacketEventData(decoder *encoding.Decoder) (PacketEventData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketEventData{}, err
	}

	return ParsePacketEventDataWithHeader(decoder, header)
}

// ParsePacketEventDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketEventDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketEventData, error) {
	if decoder.LeftToRead() < PacketEventDataSize {
		return PacketEventData{}, encoding.ErrBufferNotLargeEnough
	}

	esc := encoding.Read4Times(decoder.Byte)
	eventStringCode := EventCode(string(esc[:]))
	eventDetails := parseEvent(eventStringCode, decoder)

	return PacketEventData{
		Header: header,
		EventStringCode: eventStringCode,
		EventDetails: eventDetails,
	}, nil
}

func parseEvent(eventCode EventCode, decoder *encoding.Decoder) any {
	switch eventCode {
	case EC_FastestLap:
		return parseFastestLap(decoder)
	case EC_Retirement:
		return parseRetirement(decoder)
	case EC_TeamMateInPits:
		return parseTeamMateInPits(decoder)
	case EC_RaceWinner:
		return parseRaceWinner(decoder)
	case EC_PenaltyIssued:
		return parsePenalty(decoder)
	case EC_SpeedTrapTriggered:
		return parseSpeedTrap(decoder)
	case EC_StartLights:
		return parseStartLights(decoder)
	case EC_DriveThroughServed:
		return parseDriveThroughPenaltyServed(decoder)
	case EC_StopGoServed:
		return parseStopGoPenaltyServed(decoder)
	case EC_Flashback:
		return parseFlashback(decoder)
	case EC_ButtonStatus:
		return parseButtons(decoder)
	case EC_Overtake:
		return parseOvertake(decoder)
	}

	return nil
}

func parseFastestLap(decoder *encoding.Decoder) FastestLap {
	return FastestLap{
		VehicleIdx: decoder.Uint8(),
		LapTime:    decoder.Float32(),
	}
}

func parseRetirement(decoder *encoding.Decoder) Retirement {
	return Retirement{
		VehicleIdx: decoder.Uint8(),
	}
}

func parseTeamMateInPits(decoder *encoding.Decoder) TeamMateInPits {
	return TeamMateInPits{
		VehicleIdx: decoder.Uint8(),
	}
}

func parseRaceWinner(decoder *encoding.Decoder) RaceWinner {
	return RaceWinner{
		VehicleIdx: decoder.Uint8(),
	}
}

func parsePenalty(decoder *encoding.Decoder) Penalty {
	return Penalty{
		PenaltyType:      decoder.Uint8(),
		InfringementType: decoder.Uint8(),
		VehicleIdx:       decoder.Uint8(),
		OtherVehicleIdx:  decoder.Uint8(),
		Time:             decoder.Uint8(),
		LapNum:           decoder.Uint8(),
		PlacesGained:     decoder.Uint8(),
	}
}

func parseSpeedTrap(decoder *encoding.Decoder) SpeedTrap {
	return SpeedTrap{
		VehicleIdx:                 decoder.Uint8(),
		Speed:                      decoder.Float32(),
		IsOverallFastestInSession:  decoder.Uint8(),
		IsDriverFastestInSession:   decoder.Uint8(),
		FastestVehicleIdxInSession: decoder.Uint8(),
		FastestSpeedInSession:      decoder.Float32(),
	}
}

func parseStartLights(decoder *encoding.Decoder) StartLights {
	return StartLights{
		NumLights: decoder.Uint8(),
	}
}

func parseDriveThroughPenaltyServed(decoder *encoding.Decoder) DriveThroughPenaltyServed {
	return DriveThroughPenaltyServed{
		VehicleIdx: decoder.Uint8(),
	}
}

func parseStopGoPenaltyServed(decoder *encoding.Decoder) StopGoPenaltyServed {
	return StopGoPenaltyServed{
		VehicleIdx: decoder.Uint8(),
	}
}

func parseFlashback(decoder *encoding.Decoder) Flashback {
	return Flashback{
		FlashbackFrameIdentifier: decoder.Uint32(),
		FlashbackSessionTime:     decoder.Float32(),
	}
}

func parseButtons(decoder *encoding.Decoder) Buttons {
	return Buttons{
		ButtonStatus: enums.Button(decoder.Uint32()),
	}
}

func parseOvertake(decoder *encoding.Decoder) Overtake {
	return Overtake{
		OvertakingVehicleIdx:     decoder.Uint8(),
		BeingOvertakenVehicleIdx: decoder.Uint8(),
	}
}