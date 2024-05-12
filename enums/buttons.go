package enums

// These flags areused in the telemetry packet to determine if any buttons are being held on the controlling device. If the value below logical ANDed with the button status is set then the corresponding button is being held. These are bit fields
type Button uint32

const (
	BUT_CrossOrA        Button = 1 << 0  // Cross Or A
	BUT_TriangleOrY     Button = 1 << 1  // Triangle Or Y
	BUT_CircleOrB       Button = 1 << 2  // Circle Or B
	BUT_SquareOrX       Button = 1 << 3  // Square Or X
	BUT_DPadLeft        Button = 1 << 4  // D-Pad Left
	BUT_DPadRight       Button = 1 << 5  // D-Pad Right
	BUT_DPadUp          Button = 1 << 6  // D-Pad Up
	BUT_DPadDown        Button = 1 << 7  // D-Pad Down
	BUT_OptionsOrMenu   Button = 1 << 8  // Options Or Menu
	BUT_L1OrLB          Button = 1 << 9  // L1 Or LB
	BUT_R1OrRB          Button = 1 << 10 // R1 Or RB
	BUT_L2OrLT          Button = 1 << 11 // L2 Or LT
	BUT_R2OrRT          Button = 1 << 12 // R2 Or RT
	BUT_LeftStickClick  Button = 1 << 13 // Left Stick Click
	BUT_RightStickClick Button = 1 << 14 // Right Stick Click
	BUT_RightStickLeft  Button = 1 << 15 // Right Stick Left
	BUT_RightStickRight Button = 1 << 16 // Right Stick Right
	BUT_RightStickUp    Button = 1 << 17 // Right Stick Up
	BUT_RightStickDown  Button = 1 << 18 // Right Stick Down
	BUT_Special         Button = 1 << 19 // Special
	BUT_UDPAction1      Button = 1 << 20 // UDP Action1
	BUT_UDPAction2      Button = 1 << 21 // UDP Action2
	BUT_UDPAction3      Button = 1 << 22 // UDP Action3
	BUT_UDPAction4      Button = 1 << 23 // UDP Action4
	BUT_UDPAction5      Button = 1 << 24 // UDP Action5
	BUT_UDPAction6      Button = 1 << 25 // UDP Action6
	BUT_UDPAction7      Button = 1 << 26 // UDP Action7
	BUT_UDPAction8      Button = 1 << 27 // UDP Action8
	BUT_UDPAction9      Button = 1 << 28 // UDP Action9
	BUT_UDPAction10     Button = 1 << 29 // UDP Action10
	BUT_UDPAction11     Button = 1 << 30 // UDP Action11
	BUT_UDPAction12     Button = 1 << 31 // UDP Action12
)

// Data returns the button as a uint32
func (b Button) Data() uint32 {
	return uint32(b)
}

// With returns the button with the given flag turned on
func (b Button) With(flag Button) Button {
	return b | flag
}

// Without returns the button with the given flag turned off
func (b Button) Without(flag Button) Button {
	return b & (^flag)
}

// HasFlag returns true if the button has the given flag
func (b Button) HasFlag(flag Button) bool {
	return b&flag != 0
}

// HasAnyFlag returns true if the button has any of the given flags
func (b Button) HasAnyFlag(flags ...Button) bool {
	var mask Button
	for _, flag := range flags {
		mask = mask.With(flag)
	}

	return b.HasFlag(mask)
}
