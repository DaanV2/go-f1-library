package f1_2023

import (
	"github.com/DaanV2/go-f1-library/c"
	"github.com/DaanV2/go-f1-library/encoding"
)

const (
	PacketLobbyInfoDataSize    = 1218
	PacketLobbyInfoDataVersion = 1
)

type (
	// This packet details the players currently in a multiplayer lobby.
	// It details each player's selected car, any AI involved in the game and also the ready status of each of the participants.
	PacketLobbyInfoData struct {
		Header       PacketHeader      `json:"header"`        // Header
		NumPlayers   uint8             `json:"num_players"`   // Number of players in the lobby data
		LobbyPlayers [22]LobbyInfoData `json:"lobby_players"` //
	}

	LobbyInfoData struct {
		AiControlled uint8       `json:"ai_controlled"` // Whether the vehicle is AI (1) or Human (0) controlled
		TeamId       TeamId      `json:"team_id"`       // Team id - see appendix (255 if no team currently selected)
		Nationality  Nationality `json:"nationality"`   // Nationality of the driver
		Platform     Platform    `json:"platform"`      // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
		Name         [48]uint8   `json:"name"`          // Name of participant in UTF-8 format – null terminated,  Will be truncated with ... (U+2026) if too long
		CarNumber    uint8       `json:"car_number"`    // Car number of the player
		ReadyStatus  uint8       `json:"ready_status"`  // 0 = not ready, 1 = ready, 2 = spectating
	}
)

// GetHeader returns the header of the packet
func (p PacketLobbyInfoData) GetHeader() PacketHeader {
	return p.Header
}

// Size returns the size of the packet
func (p PacketLobbyInfoData) Size() int {
	return PacketLobbyInfoDataSize
}

func (l *LobbyInfoData) PlayerName() string {
	return c.String(l.Name[:])
}

func (l *LobbyInfoData) Ready() bool {
	return l.ReadyStatus == 1
}

// ParsePacketLobbyInfoData will parse the given data into a packet
func ParsePacketLobbyInfoData(decoder *encoding.Decoder) (PacketLobbyInfoData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketLobbyInfoData{}, err
	}

	return ParsePacketLobbyInfoDataWithHeader(decoder, header)
}

// ParsePacketLobbyInfoDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketLobbyInfoDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketLobbyInfoData, error) {
	if decoder.LeftToRead() < (PacketLobbyInfoDataSize - header.Size()) {
		return PacketLobbyInfoData{}, encoding.ErrBufferNotLargeEnough
	}

	packet := PacketLobbyInfoData{
		Header:       header,
		NumPlayers:   decoder.Uint8(),
		LobbyPlayers: parseLobbyInfoData(decoder),
	}

	return packet, nil
}

func parseLobbyInfoData(decoder *encoding.Decoder) [22]LobbyInfoData {
	items := [22]LobbyInfoData{}

	for i := range items {
		items[i] = LobbyInfoData{
			AiControlled: decoder.Uint8(),
			TeamId:       TeamId(decoder.Uint8()),
			Nationality:  Nationality(decoder.Uint8()),
			Platform:     Platform(decoder.Uint8()),
			Name:         decoder.Read48(),
			CarNumber:    decoder.Uint8(),
			ReadyStatus:  decoder.Uint8(),
		}
	}

	return items
}
