package enums

// TemperatureFormat represents the temperature format.
type TemperatureFormat uint8

const (
	TF_Celsius    TemperatureFormat = 0 // Celsius represents the Celsius temperature format.
	TF_Fahrenheit TemperatureFormat = 1 // Fahrenheit represents the Fahrenheit temperature format.
)