package f1_2023

import "github.com/DaanV2/go-f1-library/encoding"

const (
	PacketCarSetupDataFrequency = 2
	PacketCarSetupDataSize      = 1107
	PacketCarSetupDataVersion   = 1
)

type (
	// This packet details the car setups for each vehicle in the session. Note that in multiplayer games,
	// other player cars will appear as blank, you will only be able to see your own car setup,
	// regardless of the "Your Telemetry" setting. Spectators will also not be able to see any car setups.
	PacketCarSetupData struct {
		Header PacketHeader // Header

		CarSetups [22]CarSetupData
	}

	CarSetupData struct {
		FrontWing              uint8   // Front wing aero
		RearWing               uint8   // Rear wing aero
		OnThrottle             uint8   // Differential adjustment on throttle (percentage)
		OffThrottle            uint8   // Differential adjustment off throttle (percentage)
		FrontCamber            float32 // Front camber angle (suspension geometry)
		RearCamber             float32 // Rear camber angle (suspension geometry)
		FrontToe               float32 // Front toe angle (suspension geometry)
		RearToe                float32 // Rear toe angle (suspension geometry)
		FrontSuspension        uint8   // Front suspension
		RearSuspension         uint8   // Rear suspension
		FrontAntiRollBar       uint8   // Front anti-roll bar
		RearAntiRollBar        uint8   // Front anti-roll bar
		FrontSuspensionHeight  uint8   // Front ride height
		RearSuspensionHeight   uint8   // Rear ride height
		BrakePressure          uint8   // Brake pressure (percentage)
		BrakeBias              uint8   // Brake bias (percentage)
		RearLeftTyrePressure   float32 // Rear left tyre pressure (PSI)
		RearRightTyrePressure  float32 // Rear right tyre pressure (PSI)
		FrontLeftTyrePressure  float32 // Front left tyre pressure (PSI)
		FrontRightTyrePressure float32 // Front right tyre pressure (PSI)
		Ballast                uint8   // Ballast
		FuelLoad               float32 // Fuel load
	}
)

// ParsePacketCarSetupData will parse the given data into a packet
func ParsePacketCarSetupData(decoder *encoding.Decoder) (PacketCarSetupData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketCarSetupData{}, err
	}

	return ParsePacketCarSetupDataWithHeader(decoder, header)
}

// ParsePacketCarSetupDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketCarSetupDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketCarSetupData, error) {
	if decoder.LeftToRead() < PacketCarSetupDataSize {
		return PacketCarSetupData{}, encoding.ErrBufferNotLargeEnough
	}

	return PacketCarSetupData{
		Header:    header,
		CarSetups: parseCarSetupData(decoder),
	}, nil
}

func parseCarSetupData(decoder *encoding.Decoder) [22]CarSetupData {
	items := [22]CarSetupData{}

	for i := range items {
		items[i] = CarSetupData{
			FrontWing:              decoder.Uint8(),
			RearWing:               decoder.Uint8(),
			OnThrottle:             decoder.Uint8(),
			OffThrottle:            decoder.Uint8(),
			FrontCamber:            decoder.Float32(),
			RearCamber:             decoder.Float32(),
			FrontToe:               decoder.Float32(),
			RearToe:                decoder.Float32(),
			FrontSuspension:        decoder.Uint8(),
			RearSuspension:         decoder.Uint8(),
			FrontAntiRollBar:       decoder.Uint8(),
			RearAntiRollBar:        decoder.Uint8(),
			FrontSuspensionHeight:  decoder.Uint8(),
			RearSuspensionHeight:   decoder.Uint8(),
			BrakePressure:          decoder.Uint8(),
			BrakeBias:              decoder.Uint8(),
			RearLeftTyrePressure:   decoder.Float32(),
			RearRightTyrePressure:  decoder.Float32(),
			FrontLeftTyrePressure:  decoder.Float32(),
			FrontRightTyrePressure: decoder.Float32(),
			Ballast:                decoder.Uint8(),
			FuelLoad:               decoder.Float32(),
		}
	}

	return items
}