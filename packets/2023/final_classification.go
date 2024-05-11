package f1_2023

import "github.com/DaanV2/go-f1-library/encoding"

const (
	PacketFinalClassificationDataSize    = 1020
	PacketFinalClassificationDataVersion = 1
)

type (
	// This packet details the final classification at the end of the race, and the data will match with the post race results screen.
	// This is especially useful for multiplayer games where it is not always possible to send lap times on the final frame because of network delay.
	PacketFinalClassificationData struct {
		Header             PacketHeader                `json:"header"`   // Header
		NumCars            uint8                       `json:"num_cars"` // Number of cars in the final classification
		ClassificationData [22]FinalClassificationData `json:"classification_data"`
	}

	FinalClassificationData struct {
		Position          uint8    `json:"position"`             // Finishing position
		NumLaps           uint8    `json:"num_laps"`             // Number of laps completed
		GridPosition      uint8    `json:"grid_position"`        // Grid position of the car
		Points            uint8    `json:"points"`               // Number of points scored
		NumPitStops       uint8    `json:"num_pit_stops"`        // Number of pit stops made
		ResultStatus      uint8    `json:"result_status"`        // uint8, Result status
		BestLapTimeInMS   uint32   `json:"best_lap_time_in_ms"`  // Best lap time of the session in milliseconds
		TotalRaceTime     float64  `json:"total_race_time"`      // Total race time in seconds without penalties
		PenaltiesTime     uint8    `json:"penalties_time"`       // Total penalties accumulated in seconds
		NumPenalties      uint8    `json:"num_penalties"`        // Number of penalties applied to this driver
		NumTyreStints     uint8    `json:"num_tyre_stints"`      // Number of tyres stints up to maximum
		TyreStintsActual  [8]uint8 `json:"tyre_stints_actual"`   // Actual tyres used by this driver
		TyreStintsVisual  [8]uint8 `json:"tyre_stints_visual"`   // Visual tyres used by this driver
		TyreStintsEndLaps [8]uint8 `json:"tyre_stints_end_laps"` // The lap number stints end on
	}
)

// GetHeader returns the header of the packet
func (p PacketFinalClassificationData) GetHeader() PacketHeader {
	return p.Header
}

// Size returns the size of the packet
func (p PacketFinalClassificationData) Size() int {
	return PacketFinalClassificationDataSize
}


// ParsePacketFinalClassificationData will parse the given data into a packet
func ParsePacketFinalClassificationData(decoder *encoding.Decoder) (PacketFinalClassificationData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketFinalClassificationData{}, err
	}

	return ParsePacketFinalClassificationDataWithHeader(decoder, header)
}

// ParsePacketFinalClassificationDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketFinalClassificationDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketFinalClassificationData, error) {
	if decoder.LeftToRead() < (PacketFinalClassificationDataSize - header.Size()) {
		return PacketFinalClassificationData{}, encoding.ErrBufferNotLargeEnough
	}

	return PacketFinalClassificationData{
		Header:             header,
		NumCars:            decoder.Uint8(),
		ClassificationData: parseFinalClassificationData(decoder),
	}, nil
}

func parseFinalClassificationData(decoder *encoding.Decoder) [22]FinalClassificationData {
	items := [22]FinalClassificationData{}

	for i := range items {
		items[i] = FinalClassificationData{
			Position:          decoder.Uint8(),
			NumLaps:           decoder.Uint8(),
			GridPosition:      decoder.Uint8(),
			Points:            decoder.Uint8(),
			NumPitStops:       decoder.Uint8(),
			ResultStatus:      decoder.Uint8(),
			BestLapTimeInMS:   decoder.Uint32(),
			TotalRaceTime:     decoder.Float64(),
			PenaltiesTime:     decoder.Uint8(),
			NumPenalties:      decoder.Uint8(),
			NumTyreStints:     decoder.Uint8(),
			TyreStintsActual:  encoding.Read8Times(decoder.Uint8),
			TyreStintsVisual:  encoding.Read8Times(decoder.Uint8),
			TyreStintsEndLaps: encoding.Read8Times(decoder.Uint8),
		}
	}

	return items
}
