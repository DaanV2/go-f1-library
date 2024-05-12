package f1_2023

import (
	"encoding/json"

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
		Header          PacketHeader `json:"header"`            // Header
		EventStringCode EventCode    `json:"event_string_code"` // [4]uint8 Event string code, see below
		EventDetails    interface{}  `json:"event_details"`     // Event details - should be interpreted differently for each type, is a pointer to a struct of type: FastestLap, Retirement, TeamMateInPits, RaceWinner, Penalty, SpeedTrap, StartLights, DriveThroughPenaltyServed, StopGoPenaltyServed, Flashback, Buttons, Overtake
	}

	Buttons struct {
		ButtonStatus enums.Button `json:"button_status"` // uint32, Bit flags specifying which buttons are being pressed currently - see appendices
	}

	DriveThroughPenaltyServed struct {
		VehicleIdx uint8 `json:"vehicle_index"` // Vehicle index of the vehicle serving drive through
	}

	Flashback struct {
		FlashbackFrameIdentifier uint32  `json:"flashback_frame_identifier"` // Frame identifier flashed back to
		FlashbackSessionTime     float32 `json:"flashback_session_time"`     // Session time flashed back to
	}

	FastestLap struct {
		VehicleIdx uint8   `json:"vehicle_index"` // Vehicle index of car achieving fastest lap
		LapTime    float32 `json:"lap_time"`      // Lap time is in seconds
	}

	Overtake struct {
		OvertakingVehicleIdx     uint8 `json:"overtaking_vehicle_idx"`      // Vehicle index of the vehicle overtaking
		BeingOvertakenVehicleIdx uint8 `json:"being_overtaken_vehicle_idx"` // Vehicle index of the vehicle being overtaken
	}

	Penalty struct {
		PenaltyType      PenaltyType      `json:"penalty_type"`      // uint8, Penalty type – see Appendices
		InfringementType InfringementType `json:"infringement_type"` // uint8, Infringement type – see Appendices
		VehicleIdx       uint8            `json:"vehicle_index"`     // Vehicle index of the car the penalty is applied to
		OtherVehicleIdx  uint8            `json:"other_vehicle_idx"` // Vehicle index of the other car involved
		Time             uint8            `json:"time"`              // Time gained, or time spent doing action in seconds
		LapNum           uint8            `json:"lap_num"`           // Lap the penalty occurred on
		PlacesGained     uint8            `json:"places_gained"`     // Number of places gained by this
	}

	RaceWinner struct {
		VehicleIdx uint8 `json:"vehicle_index"` // Vehicle index of the race winner
	}

	Retirement struct {
		VehicleIdx uint8 `json:"vehicle_index"` // Vehicle index of car retiring
	}

	StartLights struct {
		NumLights uint8 `json:"num_lights"` // Number of lights showing
	}

	SpeedTrap struct {
		VehicleIdx                 uint8   `json:"vehicle_index"`                  // Vehicle index of the vehicle triggering speed trap
		Speed                      float32 `json:"speed"`                          // Top speed achieved in kilometres per hour
		IsOverallFastestInSession  uint8   `json:"is_overall_fastest_in_session"`  // Overall fastest speed in session = 1, otherwise 0
		IsDriverFastestInSession   uint8   `json:"is_driver_fastest_in_session"`   // Fastest speed for driver in session = 1, otherwise 0
		FastestVehicleIdxInSession uint8   `json:"fastest_vehicle_idx_in_session"` // Vehicle index of the vehicle that is the fastest in this session
		FastestSpeedInSession      float32 `json:"fastest_speed_in_session"`       // Speed of the vehicle that is the fastest in this session
	}

	TeamMateInPits struct {
		VehicleIdx uint8 `json:"vehicle_index"` // Vehicle index of team mate
	}

	StopGoPenaltyServed struct {
		VehicleIdx uint8 `json:"vehicle_index"` // Vehicle index of the vehicle serving stop go
	}
)

// GetHeader returns the header of the packet
func (p PacketEventData) GetHeader() PacketHeader {
	return p.Header
}

// Size returns the size of the packet
func (p PacketEventData) Size() int {
	return PacketEventDataSize
}

type subSection struct {
	Header          PacketHeader `json:"header"`
	EventStringCode EventCode    `json:"event_string_code"`
}

type typedEventDetails[T any] struct {
	EventDetails T `json:"event_details"`
}

func (p *PacketEventData) UnmarshalJSON(data []byte) error {
	var sub subSection
	if err := json.Unmarshal(data, &sub); err != nil {
		return err
	}

	p.Header = sub.Header
	p.EventStringCode = sub.EventStringCode

	switch sub.EventStringCode {
	case EC_FastestLap:
		return unmarshalEvent[FastestLap](p, data)
	case EC_Retirement:
		return unmarshalEvent[Retirement](p, data)
	case EC_TeamMateInPits:
		return unmarshalEvent[TeamMateInPits](p, data)
	case EC_RaceWinner:
		return unmarshalEvent[RaceWinner](p, data)
	case EC_PenaltyIssued:
		return unmarshalEvent[Penalty](p, data)
	case EC_SpeedTrapTriggered:
		return unmarshalEvent[SpeedTrap](p, data)
	case EC_StartLights:
		return unmarshalEvent[StartLights](p, data)
	case EC_DriveThroughServed:
		return unmarshalEvent[DriveThroughPenaltyServed](p, data)
	case EC_StopGoServed:
		return unmarshalEvent[StopGoPenaltyServed](p, data)
	case EC_Flashback:
		return unmarshalEvent[Flashback](p, data)
	case EC_ButtonStatus:
		return unmarshalEvent[Buttons](p, data)
	case EC_Overtake:
		return unmarshalEvent[Overtake](p, data)
	}

	return nil
}

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
	if decoder.LeftToRead() < (PacketEventDataSize - header.Size()) {
		return PacketEventData{}, encoding.ErrBufferNotLargeEnough
	}

	esc := encoding.Read4Times(decoder.Byte)
	eventStringCode := EventCode(string(esc[:]))
	eventDetails := parseEvent(eventStringCode, decoder)

	return PacketEventData{
		Header:          header,
		EventStringCode: eventStringCode,
		EventDetails:    eventDetails,
	}, nil
}

func unmarshalEvent[T any](p *PacketEventData, data []byte) error {
	var details typedEventDetails[T]
	if err := json.Unmarshal(data, &details); err != nil {
		return err
	}
	p.EventDetails = details.EventDetails
	return nil
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
		PenaltyType:      PenaltyType(decoder.Uint8()),
		InfringementType: InfringementType(decoder.Uint8()),
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
