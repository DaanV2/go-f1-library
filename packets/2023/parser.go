package f1_2023

import "github.com/DaanV2/go-f1-library/encoding"

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) PacketCarDamageData(decoder *encoding.Decoder) (PacketCarDamageData, error) {
	return ParsePacketCarDamageData(decoder)
}
func (p *Parser) PacketCarDamageDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketCarDamageData, error) {
	return ParsePacketCarDamageDataWithHeader(decoder, header)
}
func (p *Parser) PacketCarSetupData(decoder *encoding.Decoder) (PacketCarSetupData, error) {
	return ParsePacketCarSetupData(decoder)
}
func (p *Parser) PacketCarSetupDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketCarSetupData, error) {
	return ParsePacketCarSetupDataWithHeader(decoder, header)
}
func (p *Parser) PacketCarStatusData(decoder *encoding.Decoder) (PacketCarStatusData, error) {
	return ParsePacketCarStatusData(decoder)
}
func (p *Parser) PacketCarStatusDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketCarStatusData, error) {
	return ParsePacketCarStatusDataWithHeader(decoder, header)
}
func (p *Parser) PacketCarTelemetryData(decoder *encoding.Decoder) (PacketCarTelemetryData, error) {
	return ParsePacketCarTelemetryData(decoder)
}
func (p *Parser) PacketCarTelemetryDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketCarTelemetryData, error) {
	return ParsePacketCarTelemetryDataWithHeader(decoder, header)
}
func (p *Parser) PacketEventData(decoder *encoding.Decoder) (PacketEventData, error) {
	return ParsePacketEventData(decoder)
}
func (p *Parser) PacketEventDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketEventData, error) {
	return ParsePacketEventDataWithHeader(decoder, header)
}
func (p *Parser) PacketFinalClassificationData(decoder *encoding.Decoder) (PacketFinalClassificationData, error) {
	return ParsePacketFinalClassificationData(decoder)
}
func (p *Parser) PacketFinalClassificationDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketFinalClassificationData, error) {
	return ParsePacketFinalClassificationDataWithHeader(decoder, header)
}
func (p *Parser) PacketHeader(decoder *encoding.Decoder) (PacketHeader, error) {
	return ParsePacketHeader(decoder)
}
func (p *Parser) PacketLapData(decoder *encoding.Decoder) (PacketLapData, error) {
	return ParsePacketLapData(decoder)
}
func (p *Parser) PacketLapDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketLapData, error) {
	return ParsePacketLapDataWithHeader(decoder, header)
}
func (p *Parser) PacketLobbyInfoData(decoder *encoding.Decoder) (PacketLobbyInfoData, error) {
	return ParsePacketLobbyInfoData(decoder)
}
func (p *Parser) PacketLobbyInfoDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketLobbyInfoData, error) {
	return ParsePacketLobbyInfoDataWithHeader(decoder, header)
}
func (p *Parser) PacketMotionData(decoder *encoding.Decoder) (PacketMotionData, error) {
	return ParsePacketMotionData(decoder)
}
func (p *Parser) PacketMotionDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketMotionData, error) {
	return ParsePacketMotionDataWithHeader(decoder, header)
}
func (p *Parser) PacketParticipantsData(decoder *encoding.Decoder) (PacketParticipantsData, error) {
	return ParsePacketParticipantsData(decoder)
}
func (p *Parser) PacketParticipantsDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketParticipantsData, error) {
	return ParsePacketParticipantsDataWithHeader(decoder, header)
}
func (p *Parser) PacketSessionData(decoder *encoding.Decoder) (PacketSessionData, error) {
	return ParsePacketSessionData(decoder)
}
func (p *Parser) PacketSessionDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketSessionData, error) {
	return ParsePacketSessionDataWithHeader(decoder, header)
}
func (p *Parser) PacketSessionHistoryData(decoder *encoding.Decoder) (PacketSessionHistoryData, error) {
	return ParsePacketSessionHistoryData(decoder)
}
func (p *Parser) PacketSessionHistoryDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketSessionHistoryData, error) {
	return ParsePacketSessionHistoryDataWithHeader(decoder, header)
}
func (p *Parser) PacketTyreSetsData(decoder *encoding.Decoder) (PacketTyreSetsData, error) {
	return ParsePacketTyreSetsData(decoder)
}
func (p *Parser) PacketTyreSetsDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketTyreSetsData, error) {
	return ParsePacketTyreSetsDataWithHeader(decoder, header)
}
