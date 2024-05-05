package f1_2023

const (
	PacketCarDamageDataFrequency = 10
	PacketCarDamageDataSize      = 953
	PacketCarDamageDataVersion   = 1
)

type (
	// This packet details car damage parameters for all the cars in the race.
	PacketCarDamageData struct {
		Header PacketHeader // Header

		CarDamageData [22]CarDamageData
	}

	CarDamageData struct {
		TyresWear            [4]float32 // Tyre wear (percentage)
		TyresDamage          [4]uint8   // Tyre damage (percentage)
		BrakesDamage         [4]uint8   // Brakes damage (percentage)
		FrontLeftWingDamage  uint8      // Front left wing damage (percentage)
		FrontRightWingDamage uint8      // Front right wing damage (percentage)
		RearWingDamage       uint8      // Rear wing damage (percentage)
		FloorDamage          uint8      // Floor damage (percentage)
		DiffuserDamage       uint8      // Diffuser damage (percentage)
		SidepodDamage        uint8      // Sidepod damage (percentage)
		DrsFault             uint8      // Indicator for DRS fault, 0 = OK, 1 = fault
		ErsFault             uint8      // Indicator for ERS fault, 0 = OK, 1 = fault
		GearBoxDamage        uint8      // Gear box damage (percentage)
		EngineDamage         uint8      // Engine damage (percentage)
		EngineMGUHWear       uint8      // Engine wear MGU-H (percentage)
		EngineESWear         uint8      // Engine wear ES (percentage)
		EngineCEWear         uint8      // Engine wear CE (percentage)
		EngineICEWear        uint8      // Engine wear ICE (percentage)
		EngineMGUKWear       uint8      // Engine wear MGU-K (percentage)
		EngineTCWear         uint8      // Engine wear TC (percentage)
		EngineBlown          uint8      // Engine blown, 0 = OK, 1 = fault
		EngineSeized         uint8      // Engine seized, 0 = OK, 1 = fault
	}
)
