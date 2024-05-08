package f1_2023

import (
	"errors"

	"github.com/DaanV2/go-f1-library/encoding"
	"github.com/DaanV2/go-f1-library/enums"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

// PacketCarDamageData parses the packet and returns the data as a PacketCarDamageData struct
func (p *Parser) PacketCarDamageData(decoder *encoding.Decoder, header PacketHeader) (PacketCarDamageData, error) {
	return ParsePacketCarDamageDataWithHeader(decoder, header)
}

// PacketCarSetupData parses the packet and returns the data as a PacketCarSetupData struct
func (p *Parser) PacketCarSetupData(decoder *encoding.Decoder, header PacketHeader) (PacketCarSetupData, error) {
	return ParsePacketCarSetupDataWithHeader(decoder, header)
}

// PacketCarStatusData parses the packet and returns the data as a PacketCarStatusData struct
func (p *Parser) PacketCarStatusData(decoder *encoding.Decoder, header PacketHeader) (PacketCarStatusData, error) {
	return ParsePacketCarStatusDataWithHeader(decoder, header)
}

// PacketCarTelemetryData parses the packet and returns the data as a PacketCarTelemetryData struct
func (p *Parser) PacketCarTelemetryData(decoder *encoding.Decoder, header PacketHeader) (PacketCarTelemetryData, error) {
	return ParsePacketCarTelemetryDataWithHeader(decoder, header)
}

// PacketEventData parses the packet and returns the data as a PacketEventData struct
func (p *Parser) PacketEventData(decoder *encoding.Decoder, header PacketHeader) (PacketEventData, error) {
	return ParsePacketEventDataWithHeader(decoder, header)
}

// PacketFinalClassificationData parses the packet and returns the data as a PacketFinalClassificationData struct
func (p *Parser) PacketFinalClassificationData(decoder *encoding.Decoder, header PacketHeader) (PacketFinalClassificationData, error) {
	return ParsePacketFinalClassificationDataWithHeader(decoder, header)
}

// PacketHeader parses the packet and returns the data as a PacketHeader struct
func (p *Parser) PacketHeader(decoder *encoding.Decoder) (PacketHeader, error) {
	return ParsePacketHeader(decoder)
}

// PacketLapData parses the packet and returns the data as a PacketLapData struct
func (p *Parser) PacketLapData(decoder *encoding.Decoder, header PacketHeader) (PacketLapData, error) {
	return ParsePacketLapDataWithHeader(decoder, header)
}

// PacketLobbyInfoData parses the packet and returns the data as a PacketLobbyInfoData struct
func (p *Parser) PacketLobbyInfoData(decoder *encoding.Decoder, header PacketHeader) (PacketLobbyInfoData, error) {
	return ParsePacketLobbyInfoDataWithHeader(decoder, header)
}

// PacketMotionData parses the packet and returns the data as a PacketMotionData struct
func (p *Parser) PacketMotionData(decoder *encoding.Decoder, header PacketHeader) (PacketMotionData, error) {
	return ParsePacketMotionDataWithHeader(decoder, header)
}

// PacketParticipantsData parses the packet and returns the data as a PacketParticipantsData struct
func (p *Parser) PacketParticipantsData(decoder *encoding.Decoder, header PacketHeader) (PacketParticipantsData, error) {
	return ParsePacketParticipantsDataWithHeader(decoder, header)
}

// PacketSessionData parses the packet and returns the data as a PacketSessionData struct
func (p *Parser) PacketSessionData(decoder *encoding.Decoder, header PacketHeader) (PacketSessionData, error) {
	return ParsePacketSessionDataWithHeader(decoder, header)
}

// PacketSessionHistoryData parses the packet and returns the data as a PacketSessionHistoryData struct
func (p *Parser) PacketSessionHistoryData(decoder *encoding.Decoder, header PacketHeader) (PacketSessionHistoryData, error) {
	return ParsePacketSessionHistoryDataWithHeader(decoder, header)
}

// PacketSetupData parses the packet and returns the data as a PacketSetupData struct
func (p *Parser) PacketTyreSetsData(decoder *encoding.Decoder, header PacketHeader) (PacketTyreSetsData, error) {
	return ParsePacketTyreSetsDataWithHeader(decoder, header)
}

// PacketMotionExData parses the packet and returns the data as a PacketMotionExData struct
func (p *Parser) PacketMotionExData(decoder *encoding.Decoder, header PacketHeader) (PacketMotionExData, error) {
	return ParsePacketMotionExDataWithHeader(decoder, header)
}

func (p *Parser) ParsePacket(decoder *encoding.Decoder) (any, error) {
	header, err := p.PacketHeader(decoder)
	if err != nil {
		return nil, err
	}

	switch header.PacketId {
	case enums.Motion:
		return p.PacketMotionData(decoder, header)
	case enums.Session:
		return p.PacketSessionData(decoder, header)
	case enums.LapData:
		return p.PacketLapData(decoder, header)
	case enums.Event:
		return p.PacketEventData(decoder, header)
	case enums.Participants:
		return p.PacketParticipantsData(decoder, header)
	case enums.CarSetups:
		return p.PacketCarSetupData(decoder, header)
	case enums.CarTelemetry:
		return p.PacketCarTelemetryData(decoder, header)
	case enums.CarStatus:
		return p.PacketCarStatusData(decoder, header)
	case enums.FinalClassification:
		return p.PacketFinalClassificationData(decoder, header)
	case enums.LobbyInfo:
		return p.PacketLobbyInfoData(decoder, header)
	case enums.CarDamage:
		return p.PacketCarDamageData(decoder, header)
	case enums.SessionHistory:
		return p.PacketSessionHistoryData(decoder, header)
	case enums.TyreSets:
		return p.PacketTyreSetsData(decoder, header)
	case enums.MotionEx:
		return p.PacketMotionExData(decoder, header)
	}

	return nil, errors.New("unknown packet id")
}
