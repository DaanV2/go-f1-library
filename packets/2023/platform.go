package f1_2023

type Platform uint8

const (
	Steam       Platform = 1   // Steam
	PlayStation Platform = 3   // PlayStation
	Xbox        Platform = 4   // Xbox
	Origin      Platform = 6   // Origin
	Unknown     Platform = 255 // Unknown
)

func (d Platform) Id() uint8 {
	return uint8(d)
}

func (d Platform) String() string {
	switch d {
	case Steam:
		return "Steam"
	case PlayStation:
		return "PlayStation"
	case Xbox:
		return "Xbox"
	case Origin:
		return "Origin"
	}

	return "UNKNOWN"
}
