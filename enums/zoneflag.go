package enums

type ZoneFlag int8

const (
	Invalid ZoneFlag = -1 // invalid / unknown
	None    ZoneFlag = 0
	Green   ZoneFlag = 1
	Blue    ZoneFlag = 2
	Yellow  ZoneFlag = 3
)