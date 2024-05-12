package f1_2023

import "github.com/DaanV2/go-f1-library/encoding"

const (
	PacketTyreSetsDataFrequency = 20 // Frequency: 20 per second but cycling through cars
	PacketTyreSetsDataSize      = 231
	PacketTyreSetsDataVersion   = 1
)

type (
	// This packets gives a more in-depth details about tyre sets assigned to a vehicle during the session.
	//
	// Frequency: 20 per second but cycling through cars
	PacketTyreSetsData struct {
		Header      PacketHeader    `json:"header"`        // Header
		CarIdx      uint8           `json:"car_idx"`       // Index of the car this data relates to
		TyreSetData [20]TyreSetData `json:"tyre_set_data"` // 13 (dry) + 7 (wet)
		FittedIdx   uint8           `json:"fitted_idx"`    // Index into array of fitted tyre
	}

	TyreSetData struct {
		ActualTyreCompound uint8 `json:"actual_tyre_compound"` // Actual tyre compound used
		VisualTyreCompound uint8 `json:"visual_tyre_compound"` // Visual tyre compound used
		Wear               uint8 `json:"wear"`                 // Tyre wear (percentage)
		Available          uint8 `json:"available"`            // Whether this set is currently available
		RecommendedSession uint8 `json:"recommended_session"`  // Recommended session for tyre set
		LifeSpan           uint8 `json:"life_span"`            // Laps left in this tyre set
		UsableLife         uint8 `json:"usable_life"`          // Max number of laps recommended for this compound
		LapDeltaTime       int16 `json:"lap_delta_time"`       // Lap delta time in milliseconds compared to fitted set
		Fitted             uint8 `json:"fitted"`               // Whether the set is fitted or not
	}
)

// GetHeader returns the header of the packet
func (p PacketTyreSetsData) GetHeader() PacketHeader {
	return p.Header
}

// Size returns the size of the packet
func (p PacketTyreSetsData) Size() int {
	return PacketTyreSetsDataSize
}

// ParsePacketTyreSetsData will parse the given data into a packet
func ParsePacketTyreSetsData(decoder *encoding.Decoder) (PacketTyreSetsData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketTyreSetsData{}, err
	}

	return ParsePacketTyreSetsDataWithHeader(decoder, header)
}

// ParsePacketTyreSetsDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketTyreSetsDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketTyreSetsData, error) {
	if decoder.LeftToRead() < (PacketTyreSetsDataSize - header.Size()) {
		return PacketTyreSetsData{}, encoding.NewBufferNotLargeEnoughError(PacketCarDamageDataSize, decoder.LeftToRead())
	}

	return PacketTyreSetsData{
		Header:      header,
		CarIdx:      decoder.Uint8(),
		TyreSetData: parseTyreSetData(decoder),
		FittedIdx:   decoder.Uint8(),
	}, nil
}

func parseTyreSetData(decoder *encoding.Decoder) [20]TyreSetData {
	items := [20]TyreSetData{}

	for i := range items {
		items[i] = TyreSetData{
			ActualTyreCompound: decoder.Uint8(),
			VisualTyreCompound: decoder.Uint8(),
			Wear:               decoder.Uint8(),
			Available:          decoder.Uint8(),
			RecommendedSession: decoder.Uint8(),
			LifeSpan:           decoder.Uint8(),
			UsableLife:         decoder.Uint8(),
			LapDeltaTime:       decoder.Int16(),
			Fitted:             decoder.Uint8(),
		}
	}

	return items
}
