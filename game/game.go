package game

// Game is a generic interface for a game version
type Game struct {
	Drivers       Drivers // Drivers
	Nationalities Nationalities // Nationalities
	Teams         Teams // Teams
	Tracks        Tracks // Tracks
}
