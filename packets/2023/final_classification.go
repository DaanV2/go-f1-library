package f1_2023

const (
	PacketFinalClassificationDataSize    = 1020
	PacketFinalClassificationDataVersion = 1
)

type (
	// This packet details the final classification at the end of the race, and the data will match with the post race results screen.
	// This is especially useful for multiplayer games where it is not always possible to send lap times on the final frame because of network delay.
	PacketFinalClassificationData struct {
		Header PacketHeader // Header

		NumCars            uint8 // Number of cars in the final classification
		ClassificationData [22]FinalClassificationData
	}

	FinalClassificationData struct {
		Position          uint8        // Finishing position
		NumLaps           uint8        // Number of laps completed
		GridPosition      uint8        // Grid position of the car
		Points            uint8        // Number of points scored
		NumPitStops       uint8        // Number of pit stops made
		ResultStatus      ResultStatus // uint8, Result status
		BestLapTimeInMS   uint32       // Best lap time of the session in milliseconds
		TotalRaceTime     float64      // Total race time in seconds without penalties
		PenaltiesTime     uint8        // Total penalties accumulated in seconds
		NumPenalties      uint8        // Number of penalties applied to this driver
		NumTyreStints     uint8        // Number of tyres stints up to maximum
		TyreStintsActual  [8]uint8     // Actual tyres used by this driver
		TyreStintsVisual  [8]uint8     // Visual tyres used by this driver
		TyreStintsEndLaps [8]uint8     // The lap number stints end on
	}
)
