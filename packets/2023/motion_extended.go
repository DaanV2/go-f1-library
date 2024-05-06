package f1_2023

import "github.com/DaanV2/go-f1-library/encoding"

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
