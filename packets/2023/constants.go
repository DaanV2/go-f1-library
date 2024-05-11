package f1_2023

// MAX_PACKET_SIZE is the maximum size of a packet
const MAX_PACKET_SIZE = max(
	PacketCarDamageDataSize,
	PacketCarSetupsDataSize,
	PacketCarStatusDataSize,
	PacketCarTelemetryDataSize,
	PacketEventDataSize,
	PacketFinalClassificationDataSize,
	PacketHeaderSize,
	PacketLapDataSize,
	PacketLobbyInfoDataSize,
	PacketMotionDataSize,
	PacketMotionExDataSize,
	PacketParticipantsDataSize,
	PacketSessionDataSize,
	PacketSessionHistoryDataSize,
	PacketTyreSetsDataSize,
)
