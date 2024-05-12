package f1_2023

import (
	"github.com/DaanV2/go-f1-library/encoding"
)

const (
	PacketCarDamageDataFrequency = 10
	PacketCarDamageDataSize      = 953
	PacketCarDamageDataVersion   = 1
)

type (
	// This packet details car damage parameters for all the cars in the race.
	PacketCarDamageData struct {
		Header        PacketHeader      `json:"header"` // Header
		CarDamageData [22]CarDamageData `json:"car_damage_data"`
	}

	CarDamageData struct {
		TyresWear            [4]float32 `json:"tyres_wear"`              // Tyre wear (percentage)
		TyresDamage          [4]uint8   `json:"tyres_damage"`            // Tyre damage (percentage)
		BrakesDamage         [4]uint8   `json:"brakes_damage"`           // Brakes damage (percentage)
		FrontLeftWingDamage  uint8      `json:"front_left_wing_damage"`  // Front left wing damage (percentage)
		FrontRightWingDamage uint8      `json:"front_right_wing_damage"` // Front right wing damage (percentage)
		RearWingDamage       uint8      `json:"rear_wing_damage"`        // Rear wing damage (percentage)
		FloorDamage          uint8      `json:"floor_damage"`            // Floor damage (percentage)
		DiffuserDamage       uint8      `json:"diffuser_damage"`         // Diffuser damage (percentage)
		SidepodDamage        uint8      `json:"sidepod_damage"`          // Sidepod damage (percentage)
		DrsFault             uint8      `json:"drs_fault"`               // Indicator for DRS fault, 0 = OK, 1 = fault
		ErsFault             uint8      `json:"ers_fault"`               // Indicator for ERS fault, 0 = OK, 1 = fault
		GearBoxDamage        uint8      `json:"gear_box_damage"`         // Gear box damage (percentage)
		EngineDamage         uint8      `json:"engine_damage"`           // Engine damage (percentage)
		EngineMGUHWear       uint8      `json:"engine_mguh_wear"`        // Engine wear MGU-H (percentage)
		EngineESWear         uint8      `json:"engine_es_wear"`          // Engine wear ES (percentage)
		EngineCEWear         uint8      `json:"engine_ce_wear"`          // Engine wear CE (percentage)
		EngineICEWear        uint8      `json:"engine_ice_wear"`         // Engine wear ICE (percentage)
		EngineMGUKWear       uint8      `json:"engine_mguk_wear"`        // Engine wear MGU-K (percentage)
		EngineTCWear         uint8      `json:"engine_tc_wear"`          // Engine wear TC (percentage)
		EngineBlown          uint8      `json:"engine_blown"`            // Engine blown, 0 = OK, 1 = fault
		EngineSeized         uint8      `json:"engine_seized"`           // Engine seized, 0 = OK, 1 = fault
	}
)

// GetHeader returns the header of the packet
func (p PacketCarDamageData) GetHeader() PacketHeader {
	return p.Header
}

// Size returns the size of the packet
func (p PacketCarDamageData) Size() int {
	return PacketCarDamageDataSize
}

// ParsePacketCarDamageData will parse the given data into a packet
func ParsePacketCarDamageData(decoder *encoding.Decoder) (PacketCarDamageData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketCarDamageData{}, err
	}

	return ParsePacketCarDamageDataWithHeader(decoder, header)
}

// ParsePacketCarDamageDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketCarDamageDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketCarDamageData, error) {
	if decoder.LeftToRead() < (PacketCarDamageDataSize - header.Size()) {
		return PacketCarDamageData{}, encoding.NewBufferNotLargeEnoughError(PacketCarDamageDataSize, decoder.LeftToRead())
	}

	return PacketCarDamageData{
		Header:        header,
		CarDamageData: parseCarDamageData(decoder),
	}, nil
}

func parseCarDamageData(decoder *encoding.Decoder) [22]CarDamageData {
	items := [22]CarDamageData{}

	for i := range items {
		items[i] = CarDamageData{
			TyresWear:            encoding.Read4Times(decoder.Float32),
			TyresDamage:          encoding.Read4Times(decoder.Uint8),
			BrakesDamage:         encoding.Read4Times(decoder.Uint8),
			FrontLeftWingDamage:  decoder.Uint8(),
			FrontRightWingDamage: decoder.Uint8(),
			RearWingDamage:       decoder.Uint8(),
			FloorDamage:          decoder.Uint8(),
			DiffuserDamage:       decoder.Uint8(),
			SidepodDamage:        decoder.Uint8(),
			DrsFault:             decoder.Uint8(),
			ErsFault:             decoder.Uint8(),
			GearBoxDamage:        decoder.Uint8(),
			EngineDamage:         decoder.Uint8(),
			EngineMGUHWear:       decoder.Uint8(),
			EngineESWear:         decoder.Uint8(),
			EngineCEWear:         decoder.Uint8(),
			EngineICEWear:        decoder.Uint8(),
			EngineMGUKWear:       decoder.Uint8(),
			EngineTCWear:         decoder.Uint8(),
			EngineBlown:          decoder.Uint8(),
			EngineSeized:         decoder.Uint8(),
		}
	}

	return items
}
