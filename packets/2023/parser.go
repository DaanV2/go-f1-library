package f1_2023

import (
	"errors"

	"github.com/DaanV2/go-f1-library/encoding"
	"github.com/DaanV2/go-f1-library/enums"
)

type (
	PacketParser struct {
	}

	Packet interface {
		// GetHeader returns the header of the packet
		GetHeader() PacketHeader
		// Size returns the size of the packet
		Size() int
	}
)

func NewPacketParser() *PacketParser {
	return &PacketParser{}
}

// PacketHeader parses the packet and returns the data as a PacketHeader struct
func (p *PacketParser) PacketHeader(decoder *encoding.Decoder) (PacketHeader, error) {
	return ParsePacketHeader(decoder)
}

// PacketCarDamageData parses the packet and returns the data as a PacketCarDamageData struct
func (p *PacketParser) PacketCarDamageData(decoder *encoding.Decoder, header PacketHeader) (PacketCarDamageData, error) {
	return ParsePacketCarDamageDataWithHeader(decoder, header)
}

// PacketCarSetupData parses the packet and returns the data as a PacketCarSetupData struct
func (p *PacketParser) PacketCarSetupData(decoder *encoding.Decoder, header PacketHeader) (PacketCarSetupsData, error) {
	return ParsePacketCarSetupDataWithHeader(decoder, header)
}

// PacketCarStatusData parses the packet and returns the data as a PacketCarStatusData struct
func (p *PacketParser) PacketCarStatusData(decoder *encoding.Decoder, header PacketHeader) (PacketCarStatusData, error) {
	return ParsePacketCarStatusDataWithHeader(decoder, header)
}

// PacketCarTelemetryData parses the packet and returns the data as a PacketCarTelemetryData struct
func (p *PacketParser) PacketCarTelemetryData(decoder *encoding.Decoder, header PacketHeader) (PacketCarTelemetryData, error) {
	return ParsePacketCarTelemetryDataWithHeader(decoder, header)
}

// PacketEventData parses the packet and returns the data as a PacketEventData struct
func (p *PacketParser) PacketEventData(decoder *encoding.Decoder, header PacketHeader) (PacketEventData, error) {
	return ParsePacketEventDataWithHeader(decoder, header)
}

// PacketFinalClassificationData parses the packet and returns the data as a PacketFinalClassificationData struct
func (p *PacketParser) PacketFinalClassificationData(decoder *encoding.Decoder, header PacketHeader) (PacketFinalClassificationData, error) {
	return ParsePacketFinalClassificationDataWithHeader(decoder, header)
}

// PacketLapData parses the packet and returns the data as a PacketLapData struct
func (p *PacketParser) PacketLapData(decoder *encoding.Decoder, header PacketHeader) (PacketLapData, error) {
	return ParsePacketLapDataWithHeader(decoder, header)
}

// PacketLobbyInfoData parses the packet and returns the data as a PacketLobbyInfoData struct
func (p *PacketParser) PacketLobbyInfoData(decoder *encoding.Decoder, header PacketHeader) (PacketLobbyInfoData, error) {
	return ParsePacketLobbyInfoDataWithHeader(decoder, header)
}

// PacketMotionData parses the packet and returns the data as a PacketMotionData struct
func (p *PacketParser) PacketMotionData(decoder *encoding.Decoder, header PacketHeader) (PacketMotionData, error) {
	return ParsePacketMotionDataWithHeader(decoder, header)
}

// PacketParticipantsData parses the packet and returns the data as a PacketParticipantsData struct
func (p *PacketParser) PacketParticipantsData(decoder *encoding.Decoder, header PacketHeader) (PacketParticipantsData, error) {
	return ParsePacketParticipantsDataWithHeader(decoder, header)
}

// PacketSessionData parses the packet and returns the data as a PacketSessionData struct
func (p *PacketParser) PacketSessionData(decoder *encoding.Decoder, header PacketHeader) (PacketSessionData, error) {
	return ParsePacketSessionDataWithHeader(decoder, header)
}

// PacketSessionHistoryData parses the packet and returns the data as a PacketSessionHistoryData struct
func (p *PacketParser) PacketSessionHistoryData(decoder *encoding.Decoder, header PacketHeader) (PacketSessionHistoryData, error) {
	return ParsePacketSessionHistoryDataWithHeader(decoder, header)
}

// PacketSetupData parses the packet and returns the data as a PacketSetupData struct
func (p *PacketParser) PacketTyreSetsData(decoder *encoding.Decoder, header PacketHeader) (PacketTyreSetsData, error) {
	return ParsePacketTyreSetsDataWithHeader(decoder, header)
}

// PacketMotionExData parses the packet and returns the data as a PacketMotionExData struct
func (p *PacketParser) PacketMotionExData(decoder *encoding.Decoder, header PacketHeader) (PacketMotionExData, error) {
	return ParsePacketMotionExDataWithHeader(decoder, header)
}

// ParsePacket will parse the packet and return the data as a Packet struct
func (p *PacketParser) ParsePacket(decoder *encoding.Decoder) (Packet, error) {
	header, err := p.PacketHeader(decoder)
	if err != nil {
		return nil, err
	}

	switch header.PacketId {
	case enums.PID_Motion:
		return p.PacketMotionData(decoder, header)
	case enums.PID_Session:
		return p.PacketSessionData(decoder, header)
	case enums.PID_LapData:
		return p.PacketLapData(decoder, header)
	case enums.PID_Event:
		return p.PacketEventData(decoder, header)
	case enums.PID_Participants:
		return p.PacketParticipantsData(decoder, header)
	case enums.PID_CarSetups:
		return p.PacketCarSetupData(decoder, header)
	case enums.PID_CarTelemetry:
		return p.PacketCarTelemetryData(decoder, header)
	case enums.PID_CarStatus:
		return p.PacketCarStatusData(decoder, header)
	case enums.PID_FinalClassification:
		return p.PacketFinalClassificationData(decoder, header)
	case enums.PID_LobbyInfo:
		return p.PacketLobbyInfoData(decoder, header)
	case enums.PID_CarDamage:
		return p.PacketCarDamageData(decoder, header)
	case enums.PID_SessionHistory:
		return p.PacketSessionHistoryData(decoder, header)
	case enums.PID_TyreSets:
		return p.PacketTyreSetsData(decoder, header)
	case enums.PID_MotionEx:
		return p.PacketMotionExData(decoder, header)
	}

	return nil, errors.New("unknown packet id")
}

// Returns the expected size of the packet, if nothing is found will return max packet size
func (p *PacketParser) Size(packet any) int {
	if p, ok := packet.(Packet); ok {
		return p.Size()
	}

	return MAX_PACKET_SIZE
}