package f1_2023

import (
	"github.com/DaanV2/go-f1-library/encoding"
)

const (
	PacketSessionHistoryDataFrequency = 20
	PacketSessionHistoryDataSize      = 1460
	PacketSessionHistoryDataVersion   = 1
)

type (
	// This packet contains lap times and tyre usage for the session. This packet works slightly differently to other packets.
	// To reduce CPU and bandwidth, each packet relates to a specific vehicle and is sent every 1/20 s, and the vehicle being sent is cycled through.
	// Therefore in a 20 car race you should receive an update for each vehicle at least once per second.
	//
	// Note that at the end of the race, after the final classification packet has been sent,
	// a final bulk update of all the session histories for the vehicles in that session will be sent.
	PacketSessionHistoryData struct {
		Header PacketHeader // Header

		CarIdx                uint8                   // Index of the car this lap data relates to
		NumLaps               uint8                   // Num laps in the data (including current partial lap)
		NumTyreStints         uint8                   // Number of tyre stints in the data
		BestLapTimeLapNum     uint8                   // Lap the best lap time was achieved on
		BestSector1LapNum     uint8                   // Lap the best Sector 1 time was achieved on
		BestSector2LapNum     uint8                   // Lap the best Sector 2 time was achieved on
		BestSector3LapNum     uint8                   // Lap the best Sector 3 time was achieved on
		LapHistoryData        [100]LapHistoryData     // 100 laps of data max
		TyreStintsHistoryData [8]TyreStintHistoryData // 8 stints of data max
	}

	LapHistoryData struct {
		LapTimeInMS        uint32 // Lap time in milliseconds
		Sector1TimeInMS    uint16 // Sector 1 time in milliseconds
		Sector1TimeMinutes uint8  // Sector 1 whole minute part
		Sector2TimeInMS    uint16 // Sector 2 time in milliseconds
		Sector2TimeMinutes uint8  // Sector 2 whole minute part
		Sector3TimeInMS    uint16 // Sector 3 time in milliseconds
		Sector3TimeMinutes uint8  // Sector 3 whole minute part
		LapValidBitFlags   uint8  // 0x01 bit set-lap valid, 0x02 bit set-sector 1 valid, 0x04 bit set-sector 2 valid, 0x08 bit set-sector 3 valid
	}

	TyreStintHistoryData struct {
		EndLap             uint8 // Lap the tyre usage ends on (255 of current tyre)
		TyreActualCompound uint8 // Actual tyres used by this driver
		TyreVisualCompound uint8 // Visual tyres used by this driver
	}
)

func (l *LapHistoryData) ValidLap() bool {
	return l.LapValidBitFlags&0x01 == 0x01
}

func (l *LapHistoryData) ValidSector1() bool {
	return l.LapValidBitFlags&0x02 == 0x02
}

func (l *LapHistoryData) ValidSector2() bool {
	return l.LapValidBitFlags&0x04 == 0x04
}

func (l *LapHistoryData) ValidSector3() bool {
	return l.LapValidBitFlags&0x08 == 0x08
}

// ParsePacketSessionHistoryData will parse the given data into a packet
func ParsePacketSessionHistoryData(decoder *encoding.Decoder) (PacketSessionHistoryData, error) {
	header, err := ParsePacketHeader(decoder)
	if err != nil {
		return PacketSessionHistoryData{}, err
	}

	return ParsePacketSessionHistoryDataWithHeader(decoder, header)
}

// ParsePacketSessionHistoryDataWithHeader will parse the given data into a packet, expected the decoder is past the header
func ParsePacketSessionHistoryDataWithHeader(decoder *encoding.Decoder, header PacketHeader) (PacketSessionHistoryData, error) {
	if decoder.LeftToRead() < PacketSessionHistoryDataSize {
		return PacketSessionHistoryData{}, encoding.ErrBufferNotLargeEnough
	}

	return PacketSessionHistoryData{
		Header:                header,
		CarIdx:                decoder.Uint8(),
		NumLaps:               decoder.Uint8(),
		NumTyreStints:         decoder.Uint8(),
		BestLapTimeLapNum:     decoder.Uint8(),
		BestSector1LapNum:     decoder.Uint8(),
		BestSector2LapNum:     decoder.Uint8(),
		BestSector3LapNum:     decoder.Uint8(),
		LapHistoryData:        parseLapHistoryData(decoder),
		TyreStintsHistoryData: parseTyreStintHistoryData(decoder),
	}, nil
}

func parseLapHistoryData(decoder *encoding.Decoder) [100]LapHistoryData {
	items := [100]LapHistoryData{}

	for i := range items {
		items[i] = LapHistoryData{
			LapTimeInMS:        decoder.Uint32(),
			Sector1TimeInMS:    decoder.Uint16(),
			Sector1TimeMinutes: decoder.Uint8(),
			Sector2TimeInMS:    decoder.Uint16(),
			Sector2TimeMinutes: decoder.Uint8(),
			Sector3TimeInMS:    decoder.Uint16(),
			Sector3TimeMinutes: decoder.Uint8(),
			LapValidBitFlags:   decoder.Uint8(),
		}
	}

	return items
}

func parseTyreStintHistoryData(decoder *encoding.Decoder) [8]TyreStintHistoryData {
	items := [8]TyreStintHistoryData{}

	for i := range items {
		items[i] = TyreStintHistoryData{
			EndLap:             decoder.Uint8(),
			TyreActualCompound: decoder.Uint8(),
			TyreVisualCompound: decoder.Uint8(),
		}
	}

	return items
}
