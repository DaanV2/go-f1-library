package f1_2023

import "github.com/DaanV2/go-f1-library/encoding"

const (
	PacketCarStatusDataSize    = 1239
	PacketCarStatusDataVersion = 1
)

type (
	// This packet details car statuses for all the cars in the race.
	// Frequency: Rate as specified in menus
	PacketCarStatusData struct {
		Header PacketHeader // Header

		CarStatusData [22]CarStatusData
	}

	CarStatusData struct {
		TractionControl         uint8   // Traction control - 0 = off, 1 = medium, 2 = full
		AntiLockBrakes          uint8   // 0 (off) - 1 (on)
		FuelMix                 uint8   // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
		FrontBrakeBias          uint8   // Front brake bias (percentage)
		PitLimiterStatus        uint8   // Pit limiter status - 0 = off, 1 = on
		FuelInTank              float32 // Current fuel mass
		FuelCapacity            float32 // Fuel capacity
		FuelRemainingLaps       float32 // Fuel remaining in terms of laps (value on MFD)
		MaxRPM                  uint16  // Cars max RPM, point of rev limiter
		IdleRPM                 uint16  // Cars idle RPM
		MaxGears                uint8   // Maximum number of gears
		DrsAllowed              uint8   // 0 = not allowed, 1 = allowed
		DrsActivationDistance   uint16  // 0 = DRS not available, non-zero - DRS will be available in [X] metres
		ActualTyreCompound      uint8   // F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1, 21 = C0, 7 = inter, 8 = wet, F1 Classic - 9 = dry, 10 = wet, F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard, 15 = wet
		VisualTyreCompound      uint8   // F1 visual (can be different from actual compound), 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet, F1 Classic – same as above, F2 ‘19, 15 = wet, 19 – super soft, 20 = soft, 21 = medium , 22 = hard
		TyresAgeLaps            uint8   // Age in laps of the current set of tyres
		VehicleFiaFlags         int8    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow
		EnginePowerICE          float32 // Engine power output of ICE (W)
		EnginePowerMGUK         float32 // Engine power output of MGU-K (W)
		ErsStoreEnergy          float32 // ERS energy store in Joules
		ErsDeployMode           uint8   // ERS deployment mode, 0 = none, 1 = medium, 2 = hotlap, 3 = overtake
		ErsHarvestedThisLapMGUK float32 // ERS energy harvested this lap by MGU-K
		ErsHarvestedThisLapMGUH float32 // ERS energy harvested this lap by MGU-H
		ErsDeployedThisLap      float32 // ERS energy deployed this lap
		NetworkPaused           uint8   // Whether the car is paused in a network game
	}
)

// ParsePacketCarStatusData will parse the given data into a packet
func ParsePacketCarStatusData(decoder *encoding.Decoder) (PacketCarStatusData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketCarStatusData{}, err
	}

	return ParsePacketCarStatusDataWithHeader(decoder, header)
}

// ParsePacketCarStatusDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketCarStatusDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketCarStatusData, error) {
	if decoder.LeftToRead() < PacketCarStatusDataSize {
		return PacketCarStatusData{}, encoding.ErrBufferNotLargeEnough
	}

	return PacketCarStatusData{
		Header:    header,
		CarStatusData: parseCarStatusData(decoder),
	}, nil
}

func parseCarStatusData(decoder *encoding.Decoder) [22]CarStatusData {
	items := [22]CarStatusData{}

	for i := range items {
		items[i] = CarStatusData{
			TractionControl:          decoder.Uint8(),
			AntiLockBrakes:           decoder.Uint8(),
			FuelMix:                  decoder.Uint8(),
			FrontBrakeBias:           decoder.Uint8(),
			PitLimiterStatus:         decoder.Uint8(),
			FuelInTank:               decoder.Float32(),
			FuelCapacity:             decoder.Float32(),
			FuelRemainingLaps:        decoder.Float32(),
			MaxRPM:                   decoder.Uint16(),
			IdleRPM:                  decoder.Uint16(),
			MaxGears:                 decoder.Uint8(),
			DrsAllowed:               decoder.Uint8(),
			DrsActivationDistance:    decoder.Uint16(),
			ActualTyreCompound:       decoder.Uint8(),
			VisualTyreCompound:       decoder.Uint8(),
			TyresAgeLaps:             decoder.Uint8(),
			VehicleFiaFlags:          decoder.Int8(),
			EnginePowerICE:           decoder.Float32(),
			EnginePowerMGUK:          decoder.Float32(),
			ErsStoreEnergy:           decoder.Float32(),
			ErsDeployMode:            decoder.Uint8(),
			ErsHarvestedThisLapMGUK:  decoder.Float32(),
			ErsHarvestedThisLapMGUH:  decoder.Float32(),
			ErsDeployedThisLap:       decoder.Float32(),
			NetworkPaused:            decoder.Uint8(),
		}
	}

	return items
}
