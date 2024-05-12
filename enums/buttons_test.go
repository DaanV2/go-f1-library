package enums_test

import (
	"testing"

	"github.com/DaanV2/go-f1-library/enums"
	"github.com/stretchr/testify/require"
)

var buttons = []enums.Button{
	enums.BUT_CrossOrA,
	enums.BUT_TriangleOrY,
	enums.BUT_CircleOrB,
	enums.BUT_SquareOrX,
	enums.BUT_DPadLeft,
	enums.BUT_DPadRight,
	enums.BUT_DPadUp,
	enums.BUT_DPadDown,
	enums.BUT_OptionsOrMenu,
	enums.BUT_L1OrLB,
	enums.BUT_R1OrRB,
	enums.BUT_L2OrLT,
	enums.BUT_R2OrRT,
	enums.BUT_LeftStickClick,
	enums.BUT_RightStickClick,
	enums.BUT_RightStickLeft,
	enums.BUT_RightStickRight,
	enums.BUT_RightStickUp,
	enums.BUT_RightStickDown,
	enums.BUT_Special,
	enums.BUT_UDPAction1,
	enums.BUT_UDPAction2,
	enums.BUT_UDPAction3,
	enums.BUT_UDPAction4,
	enums.BUT_UDPAction5,
	enums.BUT_UDPAction6,
	enums.BUT_UDPAction7,
	enums.BUT_UDPAction8,
	enums.BUT_UDPAction9,
	enums.BUT_UDPAction10,
	enums.BUT_UDPAction11,
	enums.BUT_UDPAction12,
}

func Test_Buttons_Flags(t *testing.T) {
	for _, button := range buttons {
		if !button.HasFlag(button) {
			t.Errorf("Button %v should have flag %v", button, button)
		}
	}
}

func Test_Buttons_With(t *testing.T) {
	for i := range buttons {
		if i == 0 {
			continue
		}
		var flags enums.Button

		for _, b := range buttons[:i] {
			flags = flags.With(b)
		}

		for _, b := range buttons[:i] {
			if !flags.HasFlag(b) {
				t.Errorf("Button %v should have flag %v", flags, b)
			}
		}

		require.True(t, flags.HasAnyFlag(buttons[:i]...))
	}
}