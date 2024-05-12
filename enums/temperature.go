package enums

// TemperatureChange represents the temperature change.
type TemperatureChange int8

const (
	TC_Up       TemperatureChange = 0 // Up represents the temperature change up.
	TC_Down     TemperatureChange = 1 // Down represents the temperature change down.
	TC_NoChange TemperatureChange = 2 // NoChange represents the temperature change no change.
)