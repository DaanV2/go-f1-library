package game

type Team interface {
	// Returns the Team id
	Id() uint8
	// Returns the Team's name
	String() string
}

type Teams interface {
	// Looks up a Team by their id for the specific game version
	Get(team_id uint8) Team
	// Returns the maximum Team id for the specific game version
	Max() uint8
}
