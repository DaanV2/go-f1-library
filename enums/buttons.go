package enums

// These flags are used in the telemetry packet to determine if any buttons are being held on the controlling device. If the value below logical ANDed with the button status is set then the corresponding button is being held.
type Button uint32

const (
	CrossOrA        Button = 0x00000001
	TriangleOrY     Button = 0x00000002
	CircleOrB       Button = 0x00000004
	SquareOrX       Button = 0x00000008
	DPadLeft        Button = 0x00000010
	DPadRight       Button = 0x00000020
	DPadUp          Button = 0x00000040
	DPadDown        Button = 0x00000080
	OptionsOrMenu   Button = 0x00000100
	L1OrLB          Button = 0x00000200
	R1OrRB          Button = 0x00000400
	L2OrLT          Button = 0x00000800
	R2OrRT          Button = 0x00001000
	LeftStickClick  Button = 0x00002000
	RightStickClick Button = 0x00004000
	RightStickLeft  Button = 0x00008000
	RightStickRight Button = 0x00010000
	RightStickUp    Button = 0x00020000
	RightStickDown  Button = 0x00040000
	Special         Button = 0x00080000
	UDPAction1      Button = 0x00100000
	UDPAction2      Button = 0x00200000
	UDPAction3      Button = 0x00400000
	UDPAction4      Button = 0x00800000
	UDPAction5      Button = 0x01000000
	UDPAction6      Button = 0x02000000
	UDPAction7      Button = 0x04000000
	UDPAction8      Button = 0x08000000
	UDPAction9      Button = 0x10000000
	UDPAction10     Button = 0x20000000
	UDPAction11     Button = 0x40000000
	UDPAction12     Button = 0x80000000
)

func (b Button) HasFlag(flag Button) bool {
	return b&flag != 0
}