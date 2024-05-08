package f1_2023

import "github.com/DaanV2/go-f1-library/encoding"

const (
	PacketMotionExDataSize    = 217
	PacketMotionExDataVersion = 1
)

// The motion packet gives extended data for the car being driven with the goal of being able to drive a motion platform setup.
// Note: All wheel arrays have the following order: RL, RR, FL, FR
type PacketMotionExData struct {
	Header                 PacketHeader `json:"header"`
	SuspensionPosition     [4]float32   `json:"suspension_position"`        // Suspension position
	SuspensionVelocity     [4]float32   `json:"suspension_velocity"`        // Suspension velocity
	SuspensionAcceleration [4]float32   `json:"suspension_acceleration"`    // Suspension acceleration
	WheelSpeed             [4]float32   `json:"wheel_speed"`                // Speed of each wheel
	WheelSlipRatio         [4]float32   `json:"wheel_slip_ratio"`           // Slip ratio for each wheel
	WheelSlipAngle         [4]float32   `json:"wheel_slip_angle"`           // Slip angles for each wheel
	WheelLatForce          [4]float32   `json:"wheel_lat_force"`            // Lateral forces for each wheel
	WheelLongForce         [4]float32   `json:"wheel_long_force"`           // Longitudinal forces for each wheel
	HeightOfCOGAboveGround float32      `json:"height_of_cog_above_ground"` // Height of centre of gravity above ground
	LocalVelocityX         float32      `json:"local_velocity_x"`           // Velocity in local space – metres/s
	LocalVelocityY         float32      `json:"local_velocity_y"`           // Velocity in local space – metres/s
	LocalVelocityZ         float32      `json:"local_velocity_z"`           // Velocity in local space – metres/s
	AngularVelocityX       float32      `json:"angular_velocity_x"`         // Angular velocity x-component – radians/s
	AngularVelocityY       float32      `json:"angular_velocity_y"`         // Angular velocity y-component – radians/s
	AngularVelocityZ       float32      `json:"angular_velocity_z"`         // Angular velocity z-component – radians/s
	AngularAccelerationX   float32      `json:"angular_acceleration_x"`     // Angular acceleration x-component – radians/s/s
	AngularAccelerationY   float32      `json:"angular_acceleration_y"`     // Angular acceleration y-component – radians/s/s
	AngularAccelerationZ   float32      `json:"angular_acceleration_z"`     // Angular acceleration z-component – radians/s/s
	FrontWheelsAngle       float32      `json:"front_wheels_angle"`         // Current front wheels angle in radians
	WheelVertForce         [4]float32   `json:"wheel_vert_force"`           // Vertical forces for each wheel
}

// ParsePacketMotionExData will parse the given data into a packet
func ParsePacketMotionExData(decoder *encoding.Decoder) (PacketMotionExData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketMotionExData{}, err
	}

	return ParsePacketMotionExDataWithHeader(decoder, header)
}

// ParsePacketMotionExDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketMotionExDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketMotionExData, error) {
	if decoder.LeftToRead() < PacketMotionExDataSize {
		return PacketMotionExData{}, encoding.ErrBufferNotLargeEnough
	}

	packet := PacketMotionExData{
		Header:                 header,
		SuspensionPosition:     encoding.Read4Times(decoder.Float32),
		SuspensionVelocity:     encoding.Read4Times(decoder.Float32),
		SuspensionAcceleration: encoding.Read4Times(decoder.Float32),
		WheelSpeed:             encoding.Read4Times(decoder.Float32),
		WheelSlipRatio:         encoding.Read4Times(decoder.Float32),
		WheelSlipAngle:         encoding.Read4Times(decoder.Float32),
		WheelLatForce:          encoding.Read4Times(decoder.Float32),
		WheelLongForce:         encoding.Read4Times(decoder.Float32),
		HeightOfCOGAboveGround: decoder.Float32(),
		LocalVelocityX:         decoder.Float32(),
		LocalVelocityY:         decoder.Float32(),
		LocalVelocityZ:         decoder.Float32(),
		AngularVelocityX:       decoder.Float32(),
		AngularVelocityY:       decoder.Float32(),
		AngularVelocityZ:       decoder.Float32(),
		AngularAccelerationX:   decoder.Float32(),
		AngularAccelerationY:   decoder.Float32(),
		AngularAccelerationZ:   decoder.Float32(),
		FrontWheelsAngle:       decoder.Float32(),
		WheelVertForce:         encoding.Read4Times(decoder.Float32),
	}

	return packet, nil
}
