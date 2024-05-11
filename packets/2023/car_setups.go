package f1_2023

import "github.com/DaanV2/go-f1-library/encoding"

const (
	PacketCarSetupsDataFrequency = 2
	PacketCarSetupsDataSize      = 1107
	PacketCarSetupsDataVersion   = 1
)

type (
	// This packet details the car setups for each vehicle in the session. Note that in multiplayer games,
	// other player cars will appear as blank, you will only be able to see your own car setup,
	// regardless of the "Your Telemetry" setting. Spectators will also not be able to see any car setups.
	PacketCarSetupsData struct {
		Header    PacketHeader     `json:"header"` // Header
		CarSetups [22]CarSetupData `json:"car_setups"`
	}

	CarSetupData struct {
		FrontWing              uint8   `json:"front_wing"`                // Front wing aero
		RearWing               uint8   `json:"rear_wing"`                 // Rear wing aero
		OnThrottle             uint8   `json:"on_throttle"`               // Differential adjustment on throttle (percentage)
		OffThrottle            uint8   `json:"off_throttle"`              // Differential adjustment off throttle (percentage)
		FrontCamber            float32 `json:"front_camber"`              // Front camber angle (suspension geometry)
		RearCamber             float32 `json:"rear_camber"`               // Rear camber angle (suspension geometry)
		FrontToe               float32 `json:"front_toe"`                 // Front toe angle (suspension geometry)
		RearToe                float32 `json:"rear_toe"`                  // Rear toe angle (suspension geometry)
		FrontSuspension        uint8   `json:"front_suspension"`          // Front suspension
		RearSuspension         uint8   `json:"rear_suspension"`           // Rear suspension
		FrontAntiRollBar       uint8   `json:"front_anti_roll_bar"`       // Front anti-roll bar
		RearAntiRollBar        uint8   `json:"rear_anti_roll_bar"`        // Front anti-roll bar
		FrontSuspensionHeight  uint8   `json:"front_suspension_height"`   // Front ride height
		RearSuspensionHeight   uint8   `json:"rear_suspension_height"`    // Rear ride height
		BrakePressure          uint8   `json:"brake_pressure"`            // Brake pressure (percentage)
		BrakeBias              uint8   `json:"brake_bias"`                // Brake bias (percentage)
		RearLeftTyrePressure   float32 `json:"rear_left_tyre_pressure"`   // Rear left tyre pressure (PSI)
		RearRightTyrePressure  float32 `json:"rear_right_tyre_pressure"`  // Rear right tyre pressure (PSI)
		FrontLeftTyrePressure  float32 `json:"front_left_tyre_pressure"`  // Front left tyre pressure (PSI)
		FrontRightTyrePressure float32 `json:"front_right_tyre_pressure"` // Front right tyre pressure (PSI)
		Ballast                uint8   `json:"ballast"`                   // Ballast
		FuelLoad               float32 `json:"fuel_load"`                 // Fuel load
	}
)

// GetHeader returns the header of the packet
func (p PacketCarSetupsData) GetHeader() PacketHeader {
	return p.Header
}

// Size returns the size of the packet
func (p PacketCarSetupsData) Size() int {
	return PacketCarSetupsDataSize
}

// ParsePacketCarSetupData will parse the given data into a packet
func ParsePacketCarSetupData(decoder *encoding.Decoder) (PacketCarSetupsData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketCarSetupsData{}, err
	}

	return ParsePacketCarSetupDataWithHeader(decoder, header)
}

// ParsePacketCarSetupDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketCarSetupDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketCarSetupsData, error) {
	if decoder.LeftToRead() < (PacketCarSetupsDataSize - header.Size()) {
		return PacketCarSetupsData{}, encoding.ErrBufferNotLargeEnough
	}

	return PacketCarSetupsData{
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
