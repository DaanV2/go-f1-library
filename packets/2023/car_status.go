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
		Header        PacketHeader      `json:"header"`
		CarStatusData [22]CarStatusData `json:"car_status_data"`
	}

	CarStatusData struct {
		TractionControl         uint8   `json:"traction_control"`
		AntiLockBrakes          uint8   `json:"anti_lock_brakes"`
		FuelMix                 uint8   `json:"fuel_mix"`
		FrontBrakeBias          uint8   `json:"front_brake_bias"`
		PitLimiterStatus        uint8   `json:"pit_limiter_status"`
		FuelInTank              float32 `json:"fuel_in_tank"`
		FuelCapacity            float32 `json:"fuel_capacity"`
		FuelRemainingLaps       float32 `json:"fuel_remaining_laps"`
		MaxRPM                  uint16  `json:"max_rpm"`
		IdleRPM                 uint16  `json:"idle_rpm"`
		MaxGears                uint8   `json:"max_gears"`
		DrsAllowed              uint8   `json:"drs_allowed"`
		DrsActivationDistance   uint16  `json:"drs_activation_distance"`
		ActualTyreCompound      uint8   `json:"actual_tyre_compound"`
		VisualTyreCompound      uint8   `json:"visual_tyre_compound"`
		TyresAgeLaps            uint8   `json:"tyres_age_laps"`
		VehicleFiaFlags         int8    `json:"vehicle_fia_flags"`
		EnginePowerICE          float32 `json:"engine_power_ice"`
		EnginePowerMGUK         float32 `json:"engine_power_mguk"`
		ErsStoreEnergy          float32 `json:"ers_store_energy"`
		ErsDeployMode           uint8   `json:"ers_deploy_mode"`
		ErsHarvestedThisLapMGUK float32 `json:"ers_harvested_this_lap_mguk"`
		ErsHarvestedThisLapMGUH float32 `json:"ers_harvested_this_lap_mguh"`
		ErsDeployedThisLap      float32 `json:"ers_deployed_this_lap"`
		NetworkPaused           uint8   `json:"network_paused"`
	}
)

// GetHeader returns the header of the packet
func (p PacketCarStatusData) GetHeader() PacketHeader {
	return p.Header
}

// Size returns the size of the packet
func (p PacketCarStatusData) Size() int {
	return PacketCarStatusDataSize
}

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
	if decoder.LeftToRead() < (PacketCarStatusDataSize - header.Size()) {
		return PacketCarStatusData{}, encoding.ErrBufferNotLargeEnough
	}

	return PacketCarStatusData{
		Header:        header,
		CarStatusData: parseCarStatusData(decoder),
	}, nil
}

func parseCarStatusData(decoder *encoding.Decoder) [22]CarStatusData {
	items := [22]CarStatusData{}

	for i := range items {
		items[i] = CarStatusData{
			TractionControl:         decoder.Uint8(),
			AntiLockBrakes:          decoder.Uint8(),
			FuelMix:                 decoder.Uint8(),
			FrontBrakeBias:          decoder.Uint8(),
			PitLimiterStatus:        decoder.Uint8(),
			FuelInTank:              decoder.Float32(),
			FuelCapacity:            decoder.Float32(),
			FuelRemainingLaps:       decoder.Float32(),
			MaxRPM:                  decoder.Uint16(),
			IdleRPM:                 decoder.Uint16(),
			MaxGears:                decoder.Uint8(),
			DrsAllowed:              decoder.Uint8(),
			DrsActivationDistance:   decoder.Uint16(),
			ActualTyreCompound:      decoder.Uint8(),
			VisualTyreCompound:      decoder.Uint8(),
			TyresAgeLaps:            decoder.Uint8(),
			VehicleFiaFlags:         decoder.Int8(),
			EnginePowerICE:          decoder.Float32(),
			EnginePowerMGUK:         decoder.Float32(),
			ErsStoreEnergy:          decoder.Float32(),
			ErsDeployMode:           decoder.Uint8(),
			ErsHarvestedThisLapMGUK: decoder.Float32(),
			ErsHarvestedThisLapMGUH: decoder.Float32(),
			ErsDeployedThisLap:      decoder.Float32(),
			NetworkPaused:           decoder.Uint8(),
		}
	}

	return items
}
