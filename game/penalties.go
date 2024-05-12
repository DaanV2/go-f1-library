package game

type Penalty interface {
	// Returns the penalty id
	Id() uint8
	// Returns the penalty description
	String() string
}

type Penalties interface {
	// Looks up a penalty by their id for the specific game version
	Get(penality_id uint8) Penalty
	// Returns the maximum penalty id for the specific game version
	Max() uint8
}
