package f1_2023

import (
	"github.com/DaanV2/go-f1-library/encoding"
	"github.com/DaanV2/go-f1-library/internal/c"
)

const (
	PacketParticipantsDataFrequency = 5 // Frequency: Every 5 seconds
	PacketParticipantsDataSize      = 1306
	PacketParticipantsDataVersion   = 1
)

type (
	// This is a list of participants in the race. If the vehicle is controlled by AI, then the name will be the driver name.
	// If this is a multiplayer game, the names will be the Steam Id on PC, or the LAN name if appropriate.
	PacketParticipantsData struct {
		Header        PacketHeader        `json:"header"`          // Header
		NumActiveCars uint8               `json:"num_active_cars"` // Number of active cars in the data – should match number of cars on HUD
		Participants  [22]ParticipantData `json:"participants"`
	}

	ParticipantData struct {
		AiControlled    uint8       `json:"ai_controlled"`     // Whether the vehicle is AI (1) or Human (0) controlled
		DriverId        Driver      `json:"driver_id"`         // uint8, Driver id - see appendix, 255 if network human
		NetworkId       uint8       `json:"network_id"`        // Network id – unique identifier for network players
		TeamId          TeamId      `json:"team_id"`           // uint8, Team id - see appendix
		MyTeam          uint8       `json:"my_team"`           // My team flag – 1 = My Team, 0 = otherwise
		RaceNumber      uint8       `json:"race_number"`       // Race number of the car
		Nationality     Nationality `json:"nationality"`       // uint8, Nationality of the driver
		Name            [48]uint8   `json:"name"`              // Name of participant in UTF-8 format – null terminated Will be truncated with … (U+2026) if too long
		YourTelemetry   uint8       `json:"your_telemetry"`    // The player's UDP setting, 0 = restricted, 1 = public
		ShowOnlineNames uint8       `json:"show_online_names"` // The player's show online names setting, 0 = off, 1 = on
		Platform        Platform    `json:"platform"`          // uint8, 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
	}
)

// GetHeader returns the header of the packet
func (p PacketParticipantsData) GetHeader() PacketHeader {
	return p.Header
}

// Size returns the size of the packet
func (p PacketParticipantsData) Size() int {
	return PacketParticipantsDataSize
}

func (p *ParticipantData) ParticipantName() string {
	return c.String(p.Name[:])
}

// ParsePacketParticipantsData will parse the given data into a packet
func ParsePacketParticipantsData(decoder *encoding.Decoder) (PacketParticipantsData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketParticipantsData{}, err
	}

	return ParsePacketParticipantsDataWithHeader(decoder, header)
}

// ParsePacketParticipantsDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketParticipantsDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketParticipantsData, error) {
	if decoder.LeftToRead() < (PacketParticipantsDataSize - header.Size()) {
		return PacketParticipantsData{}, encoding.NewBufferNotLargeEnoughError(PacketCarDamageDataSize, decoder.LeftToRead())
	}

	return PacketParticipantsData{
		Header:        header,
		NumActiveCars: decoder.Uint8(),
		Participants:  parseParticipantData(decoder),
	}, nil
}

func parseParticipantData(decoder *encoding.Decoder) [22]ParticipantData {
	items := [22]ParticipantData{}

	for i := range items {
		items[i] = ParticipantData{
			AiControlled:    decoder.Uint8(),
			DriverId:        Driver(decoder.Uint8()),
			NetworkId:       decoder.Uint8(),
			TeamId:          TeamId(decoder.Uint8()),
			MyTeam:          decoder.Uint8(),
			RaceNumber:      decoder.Uint8(),
			Nationality:     Nationality(decoder.Uint8()),
			Name:            decoder.Read48(),
			YourTelemetry:   decoder.Uint8(),
			ShowOnlineNames: decoder.Uint8(),
			Platform:        Platform(decoder.Uint8()),
		}
	}

	return items
}
