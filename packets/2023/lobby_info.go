package f1_2023

import "github.com/DaanV2/go-f1-library/c"

const (
	PacketLobbyInfoDataSize    = 1218
	PacketLobbyInfoDataVersion = 1
)

type (
	// This packet details the players currently in a multiplayer lobby.
	// It details each player's selected car, any AI involved in the game and also the ready status of each of the participants.
	PacketLobbyInfoData struct {
		Header PacketHeader // Header

		// Packet specific data
		NumPlayers   uint8 // Number of players in the lobby data
		LobbyPlayers [22]LobbyInfoData
	}

	LobbyInfoData struct {
		AiControlled uint8       // Whether the vehicle is AI (1) or Human (0) controlled
		TeamId       TeamId      // Team id - see appendix (255 if no team currently selected)
		Nationality  Nationality // Nationality of the driver
		Platform     Platform    // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
		Name         [48]uint8   // Name of participant in UTF-8 format – null terminated,  Will be truncated with ... (U+2026) if too long
		CarNumber    uint8       // Car number of the player
		ReadyStatus  uint8       // 0 = not ready, 1 = ready, 2 = spectating
	}
)

func (l *LobbyInfoData) PlayerName() string {
	return c.String(l.Name[:])
}

func (l *LobbyInfoData) Ready() bool {
	return l.ReadyStatus == 1
}