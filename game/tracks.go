package game

type Track interface {
	// Returns the Track id
	Id() uint8
	// Returns the Track's name
	String() string
}

type Tracks interface {
	// Looks up a Track by their id for the specific game version
	Get(track_id uint8) Track
	// Returns the maximum Track id for the specific game version
	Max() uint8
}
