package f1_2023

import "github.com/DaanV2/go-f1-library/c"

const (
	PacketParticipantsDataFrequency = 5 // Frequency: Every 5 seconds
	PacketParticipantsDataSize      = 1306
	PacketParticipantsDataVersion   = 1
)

type (
	// This is a list of participants in the race. If the vehicle is controlled by AI, then the name will be the driver name.
	// If this is a multiplayer game, the names will be the Steam Id on PC, or the LAN name if appropriate.
	PacketParticipantsData struct {
		Header PacketHeader // Header

		NumActiveCars uint8 // Number of active cars in the data – should match number of cars on HUD
		Participants  [22]ParticipantData
	}

	ParticipantData struct {
		AiControlled    uint8       // Whether the vehicle is AI (1) or Human (0) controlled
		DriverId        Driver      // uint8, Driver id - see appendix, 255 if network human
		NetworkId       uint8       // Network id – unique identifier for network players
		TeamId          TeamId      // uint8, Team id - see appendix
		MyTeam          uint8       // My team flag – 1 = My Team, 0 = otherwise
		RaceNumber      uint8       // Race number of the car
		Nationality     Nationality // uint8, Nationality of the driver
		Name            [48]uint8   // Name of participant in UTF-8 format – null terminated Will be truncated with … (U+2026) if too long
		YourTelemetry   uint8       // The player's UDP setting, 0 = restricted, 1 = public
		ShowOnlineNames uint8       // The player's show online names setting, 0 = off, 1 = on
		Platform        Platform    // uint8, 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
	}
)

func (p *ParticipantData) ParticipantName() string {
	return c.String(p.Name[:])
}
