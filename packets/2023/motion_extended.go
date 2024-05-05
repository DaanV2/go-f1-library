package f1_2023

const (
	PacketMotionExDataSize    = 217
	PacketMotionExDataVersion = 1
)

// The motion packet gives extended data for the car being driven with the goal of being able to drive a motion platform setup.
// Note: All wheel arrays have the following order: RL, RR, FL, FR
type PacketMotionExData struct {
	Header PacketHeader // Header

	SuspensionPosition     [4]float32 //
	SuspensionVelocity     [4]float32 //
	SuspensionAcceleration [4]float32 //
	WheelSpeed             [4]float32 // Speed of each wheel
	WheelSlipRatio         [4]float32 // Slip ratio for each wheel
	WheelSlipAngle         [4]float32 // Slip angles for each wheel
	WheelLatForce          [4]float32 // Lateral forces for each wheel
	WheelLongForce         [4]float32 // Longitudinal forces for each wheel
	HeightOfCOGAboveGround float32    // Height of centre of gravity above ground
	LocalVelocityX         float32    // Velocity in local space – metres/s
	LocalVelocityY         float32    // Velocity in local space
	LocalVelocityZ         float32    // Velocity in local space
	AngularVelocityX       float32    // Angular velocity x-component – radians/s
	AngularVelocityY       float32    // Angular velocity y-component
	AngularVelocityZ       float32    // Angular velocity z-component
	AngularAccelerationX   float32    // Angular acceleration x-component – radians/s/s
	AngularAccelerationY   float32    // Angular acceleration y-component
	AngularAccelerationZ   float32    // Angular acceleration z-component
	FrontWheelsAngle       float32    // Current front wheels angle in radians
	WheelVertForce         [4]float32 // Vertical forces for each wheel
}
