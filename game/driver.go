package game

// Driver is a generic interface for a driver type
type Driver interface {
	// Returns the driver id
	Id() uint8
	// Returns the driver's name
	String() string
}

// Drivers is a generic interface for a collection of drivers
type Drivers interface {
	// Looks up a driver by their id for the specific game version
	Get(driver_id uint8) Driver
	// Returns the maximum driver id for the specific game version
	Max() uint8
}
