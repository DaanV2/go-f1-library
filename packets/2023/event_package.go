package f1_2023

import "github.com/DaanV2/go-f1-library/enums"

const (
	PacketEventDataSize    = 45
	PacketEventDataVersion = 1
)

type EventCode string

const (
	EC_SessionStarted     EventCode = "SSTA" // Sent when the session starts
	EC_SessionEnded       EventCode = "SEND" // Sent when the session ends
	EC_FastestLap         EventCode = "FTLP" // When a driver achieves the fastest lap
	EC_Retirement         EventCode = "RTMT" // When a driver retires
	EC_DRSenabled         EventCode = "DRSE" // Race control have enabled DRS
	EC_DRSdisabled        EventCode = "DRSD" // Race control have disabled DRS
	EC_TeamMateInPits     EventCode = "TMPT" // Your team mate has entered the pits
	EC_ChequeredFlag      EventCode = "CHQF" // The chequered flag has been waved
	EC_RaceWinner         EventCode = "RCWN" // The race winner is announced
	EC_PenaltyIssued      EventCode = "PENA" // A penalty has been issued – details in event
	EC_SpeedTrapTriggered EventCode = "SPTP" // Speed trap has been triggered by fastest speed
	EC_StartLights        EventCode = "STLG" // Start lights – number shown
	EC_LightsOut          EventCode = "LGOT" // Lights out
	EC_DriveThroughServed EventCode = "DTSV" // Drive through penalty served
	EC_StopGoServed       EventCode = "SGSV" // Stop go penalty served
	EC_Flashback          EventCode = "FLBK" // Flashback activated
	EC_ButtonStatus       EventCode = "BUTN" // Button status changed
	EC_RedFlag            EventCode = "RDFL" // Red flag shown
	EC_Overtake           EventCode = "OVTK" // Overtake occurred
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
