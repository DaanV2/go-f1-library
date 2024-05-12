package game

type Infringement interface {
	// Returns the infringement id
	Id() uint8
	// Returns the infringement description
	String() string
}

type Infringements interface {
	// Looks up a infringement by their id for the specific game version
	Get(infringement_id uint8) Infringement
	// Returns the maximum infringement id for the specific game version
	Max() uint8
}
