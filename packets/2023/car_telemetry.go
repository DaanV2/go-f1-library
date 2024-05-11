package f1_2023

import "github.com/DaanV2/go-f1-library/encoding"

const (
	PacketCarTelemetryDataSize    = 1352
	PacketCarTelemetryDataVersion = 1
)

type (
	// This packet details telemetry for all the cars in the race. It details various values that would be recorded on the car such as speed, throttle application, DRS etc. Note that the rev light configurations are presented separately as well and will mimic real life driver preferences.
	//
	// Frequency: Rate as specified in menus
	PacketCarTelemetryData struct {
		Header                       PacketHeader         `json:"header"`                           // Header
		CarTelemetryData             [22]CarTelemetryData `json:"car_telemetry_data"`               // Telemetry data for all cars
		MfdPanelIndex                uint8                `json:"mfd_panel_index"`                  // Index of MFD panel open - 255 = MFD closed, Single player, race – 0 = Car setup, 1 = Pits, 2 = Damage, 3 =  Engine, 4 = Temperatures, May vary depending on game mode
		MfdPanelIndexSecondaryPlayer uint8                `json:"mfd_panel_index_secondary_player"` // See above
		SuggestedGear                int8                 `json:"suggested_gear"`                   // Suggested gear for the player (1-8) 0 if no gear suggested
	}

	CarTelemetryData struct {
		Speed                   uint16     `json:"speed"`                     // Speed of car in kilometres per hour
		Throttle                float32    `json:"throttle"`                  // Amount of throttle applied (0.0 to 1.0)
		Steer                   float32    `json:"steer"`                     // Steering (-1.0 (full lock left) to 1.0 (full lock right))
		Brake                   float32    `json:"brake"`                     // Amount of brake applied (0.0 to 1.0)
		Clutch                  uint8      `json:"clutch"`                    // Amount of clutch applied (0 to 100)
		Gear                    int8       `json:"gear"`                      // Gear selected (1-8, N=0, R=-1)
		EngineRPM               uint16     `json:"engine_rpm"`                // Engine RPM
		Drs                     uint8      `json:"drs"`                       // 0 = off, 1 = on
		RevLightsPercent        uint8      `json:"rev_lights_percent"`        // Rev lights indicator (percentage)
		RevLightsBitValue       uint16     `json:"rev_lights_bit_value"`      // Rev lights (bit 0 = leftmost LED, bit 14 = rightmost LED)
		BrakesTemperature       [4]uint16  `json:"brakes_temperature"`        // Brakes temperature (celsius)
		TyresSurfaceTemperature [4]uint8   `json:"tyres_surface_temperature"` // Tyres surface temperature (celsius)
		TyresInnerTemperature   [4]uint8   `json:"tyres_inner_temperature"`   // Tyres inner temperature (celsius)
		EngineTemperature       uint16     `json:"engine_temperature"`        // Engine temperature (celsius)
		TyresPressure           [4]float32 `json:"tyres_pressure"`            // Tyres pressure (PSI)
		SurfaceType             [4]uint8   `json:"surface_type"`              // Driving surface, see appendices
	}
)

// GetHeader returns the header of the packet
func (p PacketCarTelemetryData) GetHeader() PacketHeader {
	return p.Header
}

// Size returns the size of the packet
func (p PacketCarTelemetryData) Size() int {
	return PacketCarTelemetryDataSize
}

// ParsePacketCarTelemetryData will parse the given data into a packet
func ParsePacketCarTelemetryData(decoder *encoding.Decoder) (PacketCarTelemetryData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketCarTelemetryData{}, err
	}

	return ParsePacketCarTelemetryDataWithHeader(decoder, header)
}

// ParsePacketCarTelemetryDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketCarTelemetryDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketCarTelemetryData, error) {
	if decoder.LeftToRead() < (PacketCarTelemetryDataSize - header.Size()) {
		return PacketCarTelemetryData{}, encoding.ErrBufferNotLargeEnough
	}

	packet := PacketCarTelemetryData{
		Header:                       header,
		CarTelemetryData:             parseCarTelemetryData(decoder),
		MfdPanelIndex:                decoder.Uint8(),
		MfdPanelIndexSecondaryPlayer: decoder.Uint8(),
		SuggestedGear:                decoder.Int8(),
	}

	return packet, nil
}

func parseCarTelemetryData(decoder *encoding.Decoder) [22]CarTelemetryData {
	items := [22]CarTelemetryData{}

	for i := range items {
		items[i] = CarTelemetryData{
			Speed:                   decoder.Uint16(),
			Throttle:                decoder.Float32(),
			Steer:                   decoder.Float32(),
			Brake:                   decoder.Float32(),
			Clutch:                  decoder.Uint8(),
			Gear:                    decoder.Int8(),
			EngineRPM:               decoder.Uint16(),
			Drs:                     decoder.Uint8(),
			RevLightsPercent:        decoder.Uint8(),
			RevLightsBitValue:       decoder.Uint16(),
			BrakesTemperature:       encoding.Read4Times(decoder.Uint16),
			TyresSurfaceTemperature: encoding.Read4Times(decoder.Uint8),
			TyresInnerTemperature:   encoding.Read4Times(decoder.Uint8),
			EngineTemperature:       decoder.Uint16(),
			TyresPressure:           encoding.Read4Times(decoder.Float32),
			SurfaceType:             encoding.Read4Times(decoder.Uint8),
		}
	}

	return items
}
