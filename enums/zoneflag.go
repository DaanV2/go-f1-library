package enums

import "github.com/DaanV2/go-f1-library/constants"

// ZoneFlag is a flag for the zone of a track
type ZoneFlag int8

const (
	ZF_Invalid ZoneFlag = -1 // invalid / unknown
	ZF_None    ZoneFlag = 0  // none
	ZF_Green   ZoneFlag = 1  // green flag
	ZF_Blue    ZoneFlag = 2  // blue flag
	ZF_Yellow  ZoneFlag = 3  // yellow flag
)

func (zf ZoneFlag) Id() int8 {
	return int8(zf)
}

func (zf ZoneFlag) String() string {
	switch zf {
	case ZF_Invalid:
		return "Invalid"
	case ZF_None:
		return "None"
	case ZF_Green:
		return "Green"
	case ZF_Blue:
		return "Blue"
	case ZF_Yellow:
		return "Yellow"
	default:
		return constants.UNKNOWN
	}
}
