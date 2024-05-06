package f1_2023

import "github.com/DaanV2/go-f1-library/encoding"

const (
	PacketMotionDataSize    = 1349
	PacketMotionDataVersion = 1
)

type (
	PacketMotionData struct {
		Header        PacketHeader      // Header
		CarMotionData [22]CarMotionData // Data for all cars on track
	}

	CarMotionData struct {
		WorldPositionX     float32 // World space X position - metres
		WorldPositionY     float32 // World space Y position
		WorldPositionZ     float32 // World space Z position
		WorldVelocityX     float32 // Velocity in world space X – metres/s
		WorldVelocityY     float32 // Velocity in world space Y
		WorldVelocityZ     float32 // Velocity in world space Z
		WorldForwardDirX   int16   // World space forward X direction (normalised)
		WorldForwardDirY   int16   // World space forward Y direction (normalised)
		WorldForwardDirZ   int16   // World space forward Z direction (normalised)
		WorldRightDirX     int16   // World space right X direction (normalised)
		WorldRightDirY     int16   // World space right Y direction (normalised)
		WorldRightDirZ     int16   // World space right Z direction (normalised)
		GForceLateral      float32 // Lateral G-Force component
		GForceLongitudinal float32 // Longitudinal G-Force component
		GForceVertical     float32 // Vertical G-Force component
		Yaw                float32 // Yaw angle in radians
		Pitch              float32 // Pitch angle in radians
		Roll               float32 // Roll angle in radians
	}
)

// ParsePacketMotionData will parse the given data into a packet
func ParsePacketMotionData(decoder *encoding.Decoder) (PacketMotionData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketMotionData{}, err
	}

	return ParsePacketMotionDataWithHeader(decoder, header)
}

// ParsePacketMotionDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketMotionDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketMotionData, error) {
	if decoder.LeftToRead() < PacketMotionDataSize {
		return PacketMotionData{}, encoding.ErrBufferNotLargeEnough
	}

	packet := PacketMotionData{
		Header:        header,
		CarMotionData: parseCarMotionData(decoder),
	}

	return packet, nil
}

func parseCarMotionData(decoder *encoding.Decoder) [22]CarMotionData {
	items := [22]CarMotionData{}

	for i := range items {
		items[i] = CarMotionData{
			WorldPositionX:     decoder.Float32(),
			WorldPositionY:     decoder.Float32(),
			WorldPositionZ:     decoder.Float32(),
			WorldVelocityX:     decoder.Float32(),
			WorldVelocityY:     decoder.Float32(),
			WorldVelocityZ:     decoder.Float32(),
			WorldForwardDirX:   decoder.Int16(),
			WorldForwardDirY:   decoder.Int16(),
			WorldForwardDirZ:   decoder.Int16(),
			WorldRightDirX:     decoder.Int16(),
			WorldRightDirY:     decoder.Int16(),
			WorldRightDirZ:     decoder.Int16(),
			GForceLateral:      decoder.Float32(),
			GForceLongitudinal: decoder.Float32(),
			GForceVertical:     decoder.Float32(),
			Yaw:                decoder.Float32(),
			Pitch:              decoder.Float32(),
			Roll:               decoder.Float32(),
		}
	}

	return items
}
