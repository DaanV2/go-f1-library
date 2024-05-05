package game

type Nationality interface {
	// Returns the nationality id
	Id() uint8
	// Returns the nationality's name
	String() string
}

type Nationalities interface {
	// Looks up a nationality by their id for the specific game version
	Get(nationality_id uint8) Nationality
	// Returns the maximum nationality id for the specific game version
	Max() uint8
}
